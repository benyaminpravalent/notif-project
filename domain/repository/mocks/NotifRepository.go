// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
)

// NotifRepository is an autogenerated mock type for the NotifRepository type
type NotifRepository struct {
	mock.Mock
}

// GenerateKey provides a mock function with given fields: merchantID, key
func (_m *NotifRepository) GenerateKey(merchantID int64, key string) error {
	ret := _m.Called(merchantID, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, string) error); ok {
		r0 = rf(merchantID, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
