// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	model "github.com/project/notif-project/domain/model"
	mock "github.com/stretchr/testify/mock"
)

// NotifRepository is an autogenerated mock type for the NotifRepository type
type NotifRepository struct {
	mock.Mock
}

// CheckOnProsessNotif provides a mock function with given fields: form
func (_m *NotifRepository) CheckOnProsessNotif(form model.CheckOnProsessNotif) (int64, error) {
	ret := _m.Called(form)

	var r0 int64
	if rf, ok := ret.Get(0).(func(model.CheckOnProsessNotif) int64); ok {
		r0 = rf(form)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.CheckOnProsessNotif) error); ok {
		r1 = rf(form)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckUrlExistence provides a mock function with given fields: form
func (_m *NotifRepository) CheckUrlExistence(form model.Url) (int64, error) {
	ret := _m.Called(form)

	var r0 int64
	if rf, ok := ret.Get(0).(func(model.Url) int64); ok {
		r0 = rf(form)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Url) error); ok {
		r1 = rf(form)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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

// GetMerchantUrlDetail provides a mock function with given fields: merchantID, notificationType
func (_m *NotifRepository) GetMerchantUrlDetail(merchantID int64, notificationType string) (model.GetMerchantUrlDetail, error) {
	ret := _m.Called(merchantID, notificationType)

	var r0 model.GetMerchantUrlDetail
	if rf, ok := ret.Get(0).(func(int64, string) model.GetMerchantUrlDetail); ok {
		r0 = rf(merchantID, notificationType)
	} else {
		r0 = ret.Get(0).(model.GetMerchantUrlDetail)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string) error); ok {
		r1 = rf(merchantID, notificationType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUrlDetail provides a mock function with given fields: urlID
func (_m *NotifRepository) GetUrlDetail(urlID int64) (model.GetUrlDetailRes, error) {
	ret := _m.Called(urlID)

	var r0 model.GetUrlDetailRes
	if rf, ok := ret.Get(0).(func(int64) model.GetUrlDetailRes); ok {
		r0 = rf(urlID)
	} else {
		r0 = ret.Get(0).(model.GetUrlDetailRes)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(urlID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertNotifExecution provides a mock function with given fields: form
func (_m *NotifRepository) InsertNotifExecution(form model.InsertNotifExecution) error {
	ret := _m.Called(form)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.InsertNotifExecution) error); ok {
		r0 = rf(form)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertUrl provides a mock function with given fields: form
func (_m *NotifRepository) InsertUrl(form model.Url) error {
	ret := _m.Called(form)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Url) error); ok {
		r0 = rf(form)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateNotifStatus provides a mock function with given fields: form
func (_m *NotifRepository) UpdateNotifStatus(form model.UpdateNotifStatus) error {
	ret := _m.Called(form)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.UpdateNotifStatus) error); ok {
		r0 = rf(form)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UrlToggleStatus provides a mock function with given fields: urlID
func (_m *NotifRepository) UrlToggleStatus(urlID int64) error {
	ret := _m.Called(urlID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(urlID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}