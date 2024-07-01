// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/joshsoftware/peerly-backend/internal/pkg/dto"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateAppreciation provides a mock function with given fields: ctx, apprecication
func (_m *Service) CreateAppreciation(ctx context.Context, apprecication dto.Appreciation) (dto.Appreciation, error) {
	ret := _m.Called(ctx, apprecication)

	if len(ret) == 0 {
		panic("no return value specified for CreateAppreciation")
	}

	var r0 dto.Appreciation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.Appreciation) (dto.Appreciation, error)); ok {
		return rf(ctx, apprecication)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.Appreciation) dto.Appreciation); ok {
		r0 = rf(ctx, apprecication)
	} else {
		r0 = ret.Get(0).(dto.Appreciation)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.Appreciation) error); ok {
		r1 = rf(ctx, apprecication)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAppreciation provides a mock function with given fields: ctx, filter
func (_m *Service) GetAppreciation(ctx context.Context, filter dto.AppreciationFilter) ([]dto.ResponseAppreciation, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetAppreciation")
	}

	var r0 []dto.ResponseAppreciation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.AppreciationFilter) ([]dto.ResponseAppreciation, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.AppreciationFilter) []dto.ResponseAppreciation); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.ResponseAppreciation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.AppreciationFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAppreciationById provides a mock function with given fields: ctx, appreciationId
func (_m *Service) GetAppreciationById(ctx context.Context, appreciationId int) (dto.ResponseAppreciation, error) {
	ret := _m.Called(ctx, appreciationId)

	if len(ret) == 0 {
		panic("no return value specified for GetAppreciationById")
	}

	var r0 dto.ResponseAppreciation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (dto.ResponseAppreciation, error)); ok {
		return rf(ctx, appreciationId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) dto.ResponseAppreciation); ok {
		r0 = rf(ctx, appreciationId)
	} else {
		r0 = ret.Get(0).(dto.ResponseAppreciation)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, appreciationId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateAppreciation provides a mock function with given fields: ctx, isValid, apprId
func (_m *Service) ValidateAppreciation(ctx context.Context, isValid bool, apprId int) (bool, error) {
	ret := _m.Called(ctx, isValid, apprId)

	if len(ret) == 0 {
		panic("no return value specified for ValidateAppreciation")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bool, int) (bool, error)); ok {
		return rf(ctx, isValid, apprId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bool, int) bool); ok {
		r0 = rf(ctx, isValid, apprId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, bool, int) error); ok {
		r1 = rf(ctx, isValid, apprId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
