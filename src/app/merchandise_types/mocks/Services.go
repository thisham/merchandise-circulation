// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	merchandise_types "merchandise-circulation-api/src/app/merchandise_types"

	mock "github.com/stretchr/testify/mock"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	mock.Mock
}

// GetAllMerchandiseTypes provides a mock function with given fields:
func (_m *Services) GetAllMerchandiseTypes() ([]merchandise_types.Domain, error) {
	ret := _m.Called()

	var r0 []merchandise_types.Domain
	if rf, ok := ret.Get(0).(func() []merchandise_types.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]merchandise_types.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
