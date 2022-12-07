// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	class "api-alta-dashboard/features/class"

	mock "github.com/stretchr/testify/mock"
)

// ClassRepository is an autogenerated mock type for the RepositoryInterface type
type ClassRepository struct {
	mock.Mock
}

// CreateClass provides a mock function with given fields: input
func (_m *ClassRepository) CreateClass(input class.Core) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(class.Core) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteClass provides a mock function with given fields: id
func (_m *ClassRepository) DeleteClass(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllClass provides a mock function with given fields:
func (_m *ClassRepository) GetAllClass() ([]class.Core, error) {
	ret := _m.Called()

	var r0 []class.Core
	if rf, ok := ret.Get(0).(func() []class.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]class.Core)
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

// GetAllWithSearchClass provides a mock function with given fields: query
func (_m *ClassRepository) GetAllWithSearchClass(query string) ([]class.Core, error) {
	ret := _m.Called(query)

	var r0 []class.Core
	if rf, ok := ret.Get(0).(func(string) []class.Core); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]class.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIdClass provides a mock function with given fields: id
func (_m *ClassRepository) GetByIdClass(id int) (class.Core, error) {
	ret := _m.Called(id)

	var r0 class.Core
	if rf, ok := ret.Get(0).(func(int) class.Core); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(class.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateClass provides a mock function with given fields: input, id
func (_m *ClassRepository) UpdateClass(input class.Core, id int) error {
	ret := _m.Called(input, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(class.Core, int) error); ok {
		r0 = rf(input, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewClassRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewClassRepository creates a new instance of ClassRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClassRepository(t mockConstructorTestingTNewClassRepository) *ClassRepository {
	mock := &ClassRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
