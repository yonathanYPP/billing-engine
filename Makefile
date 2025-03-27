GO := go
GOFLAGS := -mod=vendor
GOTEST := $(GO) test -mod=vendor
GOTESTS := gotests -all -w

MOCKGEN := mockgen
MOCK_DIR := app/tests
USECASE_DIR := app/usecases
SERVICE_DIR := app/service
TEST_DIR := app/tests

.PHONY: test
test:
	$(GOTEST) ./$(MOCK_DIR)/...

.PHONY: mock
mock:
	mkdir -p $(MOCK_DIR)
	$(MOCKGEN) -source=$(USECASE_DIR)/loan_usecase.go -destination=$(MOCK_DIR)/mock_loan_usecase.go -package=mocks
	$(MOCKGEN) -source=$(SERVICE_DIR)/loan_service.go -destination=$(MOCK_DIR)/mock_loan_service.go -package=mocks

.PHONY: generate-tests
generate-tests: mock
	$(GOTESTS) $(MOCK_DIR)/mock_loan_usecase.go
	$(GOTESTS) $(MOCK_DIR)/mock_loan_service.go

.PHONY: clean
clean:
	$(GO) clean
	rm -rf $(MOCK_DIR)/*_test.go

.PHONY: vendor
vendor:
	$(GO) mod vendor
