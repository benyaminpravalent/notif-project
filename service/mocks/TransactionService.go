// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/richardsahvic/jamtangan/domain/model"
	mock "github.com/stretchr/testify/mock"
)

// TransactionService is an autogenerated mock type for the TransactionService type
type TransactionService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, request
func (_m *TransactionService) Create(ctx context.Context, request model.CreateTransactionRequest) (int, *model.BaseResponse) {
	ret := _m.Called(ctx, request)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateTransactionRequest) int); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 *model.BaseResponse
	if rf, ok := ret.Get(1).(func(context.Context, model.CreateTransactionRequest) *model.BaseResponse); ok {
		r1 = rf(ctx, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.BaseResponse)
		}
	}

	return r0, r1
}

// GetDetail provides a mock function with given fields: ctx, orderID
func (_m *TransactionService) GetDetail(ctx context.Context, orderID string) (int, *model.BaseResponse) {
	ret := _m.Called(ctx, orderID)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, orderID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 *model.BaseResponse
	if rf, ok := ret.Get(1).(func(context.Context, string) *model.BaseResponse); ok {
		r1 = rf(ctx, orderID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.BaseResponse)
		}
	}

	return r0, r1
}
