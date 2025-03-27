// Code generated by MockGen. DO NOT EDIT.
// Source: app/service/payment_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "billing-engine/app/models"
	reflect "reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestNewMockPaymentService(t *testing.T) {
	type args struct {
		ctrl *gomock.Controller
	}
	tests := []struct {
		name string
		args args
		want *MockPaymentService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMockPaymentService(tt.args.ctrl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMockPaymentService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPaymentService_EXPECT(t *testing.T) {
	type fields struct {
		ctrl     *gomock.Controller
		recorder *MockPaymentServiceMockRecorder
	}
	tests := []struct {
		name   string
		fields fields
		want   *MockPaymentServiceMockRecorder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockPaymentService{
				ctrl:     tt.fields.ctrl,
				recorder: tt.fields.recorder,
			}
			if got := m.EXPECT(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPaymentService.EXPECT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPaymentService_GetPaymentsByLoanID(t *testing.T) {
	type fields struct {
		ctrl     *gomock.Controller
		recorder *MockPaymentServiceMockRecorder
	}
	type args struct {
		LoanID uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Payment
		want1   int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockPaymentService{
				ctrl:     tt.fields.ctrl,
				recorder: tt.fields.recorder,
			}
			got, got1, err := m.GetPaymentsByLoanID(tt.args.LoanID)
			if (err != nil) != tt.wantErr {
				t.Errorf("MockPaymentService.GetPaymentsByLoanID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPaymentService.GetPaymentsByLoanID() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MockPaymentService.GetPaymentsByLoanID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMockPaymentServiceMockRecorder_GetPaymentsByLoanID(t *testing.T) {
	type fields struct {
		mock *MockPaymentService
	}
	type args struct {
		LoanID interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *gomock.Call
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := &MockPaymentServiceMockRecorder{
				mock: tt.fields.mock,
			}
			if got := mr.GetPaymentsByLoanID(tt.args.LoanID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPaymentServiceMockRecorder.GetPaymentsByLoanID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPaymentService_PayPendingPayment(t *testing.T) {
	type fields struct {
		ctrl     *gomock.Controller
		recorder *MockPaymentServiceMockRecorder
	}
	type args struct {
		LoanID             uint
		totalAmount        float64
		startInstallment   int64
		pendingInstallment int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Payment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockPaymentService{
				ctrl:     tt.fields.ctrl,
				recorder: tt.fields.recorder,
			}
			got, err := m.PayPendingPayment(tt.args.LoanID, tt.args.totalAmount, tt.args.startInstallment, tt.args.pendingInstallment)
			if (err != nil) != tt.wantErr {
				t.Errorf("MockPaymentService.PayPendingPayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPaymentService.PayPendingPayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPaymentServiceMockRecorder_PayPendingPayment(t *testing.T) {
	type fields struct {
		mock *MockPaymentService
	}
	type args struct {
		LoanID             interface{}
		totalAmount        interface{}
		startInstallment   interface{}
		pendingInstallment interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *gomock.Call
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := &MockPaymentServiceMockRecorder{
				mock: tt.fields.mock,
			}
			if got := mr.PayPendingPayment(tt.args.LoanID, tt.args.totalAmount, tt.args.startInstallment, tt.args.pendingInstallment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPaymentServiceMockRecorder.PayPendingPayment() = %v, want %v", got, tt.want)
			}
		})
	}
}
