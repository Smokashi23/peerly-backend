// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/joshsoftware/peerly-backend/internal/pkg/dto"
	mock "github.com/stretchr/testify/mock"

	repository "github.com/joshsoftware/peerly-backend/internal/repository"
)

// CoreValueStorer is an autogenerated mock type for the CoreValueStorer type
type CoreValueStorer struct {
	mock.Mock
}

// CheckUniqueCoreVal provides a mock function with given fields: ctx, text
func (_m *CoreValueStorer) CheckUniqueCoreVal(ctx context.Context, text string) (bool, error) {
	ret := _m.Called(ctx, text)

	if len(ret) == 0 {
		panic("no return value specified for CheckUniqueCoreVal")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, text)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, text)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, text)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCoreValue provides a mock function with given fields: ctx, coreValue
func (_m *CoreValueStorer) CreateCoreValue(ctx context.Context, coreValue dto.CreateCoreValueReq) (repository.CoreValue, error) {
	ret := _m.Called(ctx, coreValue)

	if len(ret) == 0 {
		panic("no return value specified for CreateCoreValue")
	}

	var r0 repository.CoreValue
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.CreateCoreValueReq) (repository.CoreValue, error)); ok {
		return rf(ctx, coreValue)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.CreateCoreValueReq) repository.CoreValue); ok {
		r0 = rf(ctx, coreValue)
	} else {
		r0 = ret.Get(0).(repository.CoreValue)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.CreateCoreValueReq) error); ok {
		r1 = rf(ctx, coreValue)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCoreValue provides a mock function with given fields: ctx, coreValueID
func (_m *CoreValueStorer) GetCoreValue(ctx context.Context, coreValueID int64) (repository.CoreValue, error) {
	ret := _m.Called(ctx, coreValueID)

	if len(ret) == 0 {
		panic("no return value specified for GetCoreValue")
	}

	var r0 repository.CoreValue
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (repository.CoreValue, error)); ok {
		return rf(ctx, coreValueID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) repository.CoreValue); ok {
		r0 = rf(ctx, coreValueID)
	} else {
		r0 = ret.Get(0).(repository.CoreValue)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, coreValueID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCoreValues provides a mock function with given fields: ctx
func (_m *CoreValueStorer) ListCoreValues(ctx context.Context) ([]repository.CoreValue, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListCoreValues")
	}

	var r0 []repository.CoreValue
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]repository.CoreValue, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []repository.CoreValue); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.CoreValue)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCoreValue provides a mock function with given fields: ctx, coreValue
func (_m *CoreValueStorer) UpdateCoreValue(ctx context.Context, coreValue dto.UpdateQueryRequest) (repository.CoreValue, error) {
	ret := _m.Called(ctx, coreValue)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCoreValue")
	}

	var r0 repository.CoreValue
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.UpdateQueryRequest) (repository.CoreValue, error)); ok {
		return rf(ctx, coreValue)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.UpdateQueryRequest) repository.CoreValue); ok {
		r0 = rf(ctx, coreValue)
	} else {
		r0 = ret.Get(0).(repository.CoreValue)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.UpdateQueryRequest) error); ok {
		r1 = rf(ctx, coreValue)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCoreValueStorer creates a new instance of CoreValueStorer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCoreValueStorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *CoreValueStorer {
	mock := &CoreValueStorer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
