// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/joshsoftware/peerly-backend/internal/pkg/dto"
	mock "github.com/stretchr/testify/mock"

	repository "github.com/joshsoftware/peerly-backend/internal/repository"

	sqlx "github.com/jmoiron/sqlx"
)

// AppreciationStorer is an autogenerated mock type for the AppreciationStorer type
type AppreciationStorer struct {
	mock.Mock
}

// BeginTx provides a mock function with given fields: ctx
func (_m *AppreciationStorer) BeginTx(ctx context.Context) (repository.Transaction, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for BeginTx")
	}

	var r0 repository.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (repository.Transaction, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) repository.Transaction); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAppreciation provides a mock function with given fields: ctx, tx, appreciation
func (_m *AppreciationStorer) CreateAppreciation(ctx context.Context, tx repository.Transaction, appreciation dto.Appreciation) (repository.Appreciation, error) {
	ret := _m.Called(ctx, tx, appreciation)

	if len(ret) == 0 {
		panic("no return value specified for CreateAppreciation")
	}

	var r0 repository.Appreciation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.Transaction, dto.Appreciation) (repository.Appreciation, error)); ok {
		return rf(ctx, tx, appreciation)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repository.Transaction, dto.Appreciation) repository.Appreciation); ok {
		r0 = rf(ctx, tx, appreciation)
	} else {
		r0 = ret.Get(0).(repository.Appreciation)
	}

	if rf, ok := ret.Get(1).(func(context.Context, repository.Transaction, dto.Appreciation) error); ok {
		r1 = rf(ctx, tx, appreciation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAppreciation provides a mock function with given fields: ctx, tx, apprId
func (_m *AppreciationStorer) DeleteAppreciation(ctx context.Context, tx repository.Transaction, apprId int32) error {
	ret := _m.Called(ctx, tx, apprId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAppreciation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.Transaction, int32) error); ok {
		r0 = rf(ctx, tx, apprId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAppreciationById provides a mock function with given fields: ctx, tx, appreciationId
func (_m *AppreciationStorer) GetAppreciationById(ctx context.Context, tx repository.Transaction, appreciationId int32) (repository.AppreciationInfo, error) {
	ret := _m.Called(ctx, tx, appreciationId)

	if len(ret) == 0 {
		panic("no return value specified for GetAppreciationById")
	}

	var r0 repository.AppreciationInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.Transaction, int32) (repository.AppreciationInfo, error)); ok {
		return rf(ctx, tx, appreciationId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repository.Transaction, int32) repository.AppreciationInfo); ok {
		r0 = rf(ctx, tx, appreciationId)
	} else {
		r0 = ret.Get(0).(repository.AppreciationInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, repository.Transaction, int32) error); ok {
		r1 = rf(ctx, tx, appreciationId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleTransaction provides a mock function with given fields: ctx, tx, isSuccess
func (_m *AppreciationStorer) HandleTransaction(ctx context.Context, tx repository.Transaction, isSuccess bool) error {
	ret := _m.Called(ctx, tx, isSuccess)

	if len(ret) == 0 {
		panic("no return value specified for HandleTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.Transaction, bool) error); ok {
		r0 = rf(ctx, tx, isSuccess)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitiateQueryExecutor provides a mock function with given fields: tx
func (_m *AppreciationStorer) InitiateQueryExecutor(tx repository.Transaction) sqlx.Ext {
	ret := _m.Called(tx)

	if len(ret) == 0 {
		panic("no return value specified for InitiateQueryExecutor")
	}

	var r0 sqlx.Ext
	if rf, ok := ret.Get(0).(func(repository.Transaction) sqlx.Ext); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sqlx.Ext)
		}
	}

	return r0
}

// IsUserPresent provides a mock function with given fields: ctx, tx, userID
func (_m *AppreciationStorer) IsUserPresent(ctx context.Context, tx repository.Transaction, userID int64) (bool, error) {
	ret := _m.Called(ctx, tx, userID)

	if len(ret) == 0 {
		panic("no return value specified for IsUserPresent")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.Transaction, int64) (bool, error)); ok {
		return rf(ctx, tx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repository.Transaction, int64) bool); ok {
		r0 = rf(ctx, tx, userID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, repository.Transaction, int64) error); ok {
		r1 = rf(ctx, tx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAppreciations provides a mock function with given fields: ctx, tx, filter
func (_m *AppreciationStorer) ListAppreciations(ctx context.Context, tx repository.Transaction, filter dto.AppreciationFilter) ([]repository.AppreciationInfo, repository.Pagination, error) {
	ret := _m.Called(ctx, tx, filter)

	if len(ret) == 0 {
		panic("no return value specified for ListAppreciations")
	}

	var r0 []repository.AppreciationInfo
	var r1 repository.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.Transaction, dto.AppreciationFilter) ([]repository.AppreciationInfo, repository.Pagination, error)); ok {
		return rf(ctx, tx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repository.Transaction, dto.AppreciationFilter) []repository.AppreciationInfo); ok {
		r0 = rf(ctx, tx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.AppreciationInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, repository.Transaction, dto.AppreciationFilter) repository.Pagination); ok {
		r1 = rf(ctx, tx, filter)
	} else {
		r1 = ret.Get(1).(repository.Pagination)
	}

	if rf, ok := ret.Get(2).(func(context.Context, repository.Transaction, dto.AppreciationFilter) error); ok {
		r2 = rf(ctx, tx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewAppreciationStorer creates a new instance of AppreciationStorer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAppreciationStorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *AppreciationStorer {
	mock := &AppreciationStorer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
