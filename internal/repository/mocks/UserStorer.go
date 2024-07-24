// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/joshsoftware/peerly-backend/internal/pkg/dto"
	mock "github.com/stretchr/testify/mock"

	repository "github.com/joshsoftware/peerly-backend/internal/repository"
)

// UserStorer is an autogenerated mock type for the UserStorer type
type UserStorer struct {
	mock.Mock
}

// CreateNewUser provides a mock function with given fields: ctx, user
func (_m *UserStorer) CreateNewUser(ctx context.Context, user dto.User) (repository.User, error) {
	ret := _m.Called(ctx, user)

	var r0 repository.User
	if rf, ok := ret.Get(0).(func(context.Context, dto.User) repository.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(repository.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dto.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGradeByName provides a mock function with given fields: ctx, name
func (_m *UserStorer) GetGradeByName(ctx context.Context, name string) (repository.Grade, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetGradeByName")
	}

	var r0 repository.Grade
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (repository.Grade, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) repository.Grade); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(repository.Grade)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRewardMultiplier provides a mock function with given fields: ctx
func (_m *UserStorer) GetRewardMultiplier(ctx context.Context) (int64, error) {
	ret := _m.Called(ctx)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoleByName provides a mock function with given fields: ctx, name
func (_m *UserStorer) GetRoleByName(ctx context.Context, name string) (int64, error) {
	ret := _m.Called(ctx, name)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalUserCount provides a mock function with given fields: ctx, reqData
func (_m *UserStorer) GetTotalUserCount(ctx context.Context, reqData dto.UserListReq) (int64, error) {
	ret := _m.Called(ctx, reqData)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, dto.UserListReq) int64); ok {
		r0 = rf(ctx, reqData)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dto.UserListReq) error); ok {
		r1 = rf(ctx, reqData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *UserStorer) GetUserByEmail(ctx context.Context, email string) (repository.User, error) {
	ret := _m.Called(ctx, email)

	var r0 repository.User
	if rf, ok := ret.Get(0).(func(context.Context, string) repository.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(repository.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUsers provides a mock function with given fields: ctx, reqData
func (_m *UserStorer) ListUsers(ctx context.Context, reqData dto.UserListReq) ([]repository.User, error) {
	ret := _m.Called(ctx, reqData)

	var r0 []repository.User
	if rf, ok := ret.Get(0).(func(context.Context, dto.UserListReq) []repository.User); ok {
		r0 = rf(ctx, reqData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.UserListReq) error); ok {
		r1 = rf(ctx, reqData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SyncData provides a mock function with given fields: ctx, updateData
func (_m *UserStorer) SyncData(ctx context.Context, updateData dto.User) error {
	ret := _m.Called(ctx, updateData)

	if len(ret) == 0 {
		panic("no return value specified for SyncData")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.User) error); ok {
		r0 = rf(ctx, updateData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func NewUserStorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserStorer {
	mock := &UserStorer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
