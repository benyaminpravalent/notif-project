// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/richardsahvic/jamtangan/domain/model"
	mock "github.com/stretchr/testify/mock"
)

// ProductService is an autogenerated mock type for the ProductService type
type ProductService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, request
func (_m *ProductService) Create(ctx context.Context, request model.CreateProductRequest) (int, *model.BaseResponse) {
	ret := _m.Called(ctx, request)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateProductRequest) int); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 *model.BaseResponse
	if rf, ok := ret.Get(1).(func(context.Context, model.CreateProductRequest) *model.BaseResponse); ok {
		r1 = rf(ctx, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.BaseResponse)
		}
	}

	return r0, r1
}

// GetByBrandID provides a mock function with given fields: ctx, brandID
func (_m *ProductService) GetByBrandID(ctx context.Context, brandID string) (int, *model.BaseResponse) {
	ret := _m.Called(ctx, brandID)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, brandID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 *model.BaseResponse
	if rf, ok := ret.Get(1).(func(context.Context, string) *model.BaseResponse); ok {
		r1 = rf(ctx, brandID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.BaseResponse)
		}
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, productID
func (_m *ProductService) GetByID(ctx context.Context, productID string) (int, *model.BaseResponse) {
	ret := _m.Called(ctx, productID)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, productID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 *model.BaseResponse
	if rf, ok := ret.Get(1).(func(context.Context, string) *model.BaseResponse); ok {
		r1 = rf(ctx, productID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.BaseResponse)
		}
	}

	return r0, r1
}
