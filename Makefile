GO := go
GOFLAGS := -mod=vendor
GOTEST := $(GO) test -mod=vendor
GOTESTS := gotests -all -w

MOCKGEN := mockgen
MOCK_DIR := app/tests
USECASE_DIR := app/usecases
SERVICE_DIR := app/service
TEST_DIR := app/tests

USECASE_FILES=$(wildcard $(USECASE_DIR)/*.go)
SERVICE_FILES=$(wildcard $(SERVICE_DIR)/*.go)

.PHONY: install-tools
install-tools:
	$(GO) install github.com/golang/mock/mockgen@latest
	$(GO) install github.com/cweill/gotests/gotests@latest

.PHONY: test
test:
	$(GOTEST) ./$(MOCK_DIR)/...

.PHONY: mock
mock: install-tools
	@for file in $(USECASE_FILES) $(SERVICE_FILES); do \
		basefile=$$(basename $$file .go); \
		mkdir -p $(MOCK_DIR)/$$basefile; \
		$(MOCKGEN) -source=$$file -destination=$(MOCK_DIR)/$$basefile/mock_$$basefile.go -package=mocks; \
	done

.PHONY: generate-tests
generate-tests: mock
	@for file in $(USECASE_FILES) $(SERVICE_FILES); do \
		basefile=$$(basename $$file .go); \
		mkdir -p $(TEST_DIR)/$$basefile; \
		$(GOTESTS) $(MOCK_DIR)/$$basefile/mock_$$basefile.go -outputdir=$(TEST_DIR)/$$basefile; \
	done

.PHONY: clean
clean:
	$(GO) clean
	rm -rf $(MOCK_DIR)/*_test.go

.PHONY: vendor
vendor:
	$(GO) mod vendor
