// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Artifact is an autogenerated mock type for the Artifact type
type Artifact struct {
	mock.Mock
}

// FileName provides a mock function with given fields:
func (_m *Artifact) FileName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetData provides a mock function with given fields:
func (_m *Artifact) GetData() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
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

// Save provides a mock function with given fields: path
func (_m *Artifact) Save(path string) (bool, error) {
	ret := _m.Called(path)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveToDir provides a mock function with given fields: dir
func (_m *Artifact) SaveToDir(dir string) (bool, error) {
	ret := _m.Called(dir)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(dir)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(dir)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
