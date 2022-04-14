// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	merchandises "merchandise-circulation-api/src/app/merchandises"

	mock "github.com/stretchr/testify/mock"
)

// Data is an autogenerated mock type for the Data type
type Data struct {
	mock.Mock
}

// InsertData provides a mock function with given fields: data
func (_m *Data) InsertData(data merchandises.Domain) (merchandises.Domain, error) {
	ret := _m.Called(data)

	var r0 merchandises.Domain
	if rf, ok := ret.Get(0).(func(merchandises.Domain) merchandises.Domain); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(merchandises.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(merchandises.Domain) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectData provides a mock function with given fields:
func (_m *Data) SelectData() []merchandises.Domain {
	ret := _m.Called()

	var r0 []merchandises.Domain
	if rf, ok := ret.Get(0).(func() []merchandises.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]merchandises.Domain)
		}
	}

	return r0
}
