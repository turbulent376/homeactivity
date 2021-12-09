// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "git.jetbrains.space/orbi/fcsd/auth/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// SessionStorage is an autogenerated mock type for the SessionStorage type
type SessionStorage struct {
	mock.Mock
}

// CreateSession provides a mock function with given fields: ctx, session
func (_m *SessionStorage) CreateSession(ctx context.Context, session *domain.Session) error {
	ret := _m.Called(ctx, session)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Session) error); ok {
		r0 = rf(ctx, session)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSession provides a mock function with given fields: ctx, id
func (_m *SessionStorage) DeleteSession(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSession provides a mock function with given fields: ctx, id
func (_m *SessionStorage) GetSession(ctx context.Context, id string) (*domain.Session, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Session
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Session); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserSessions provides a mock function with given fields: ctx, id
func (_m *SessionStorage) GetUserSessions(ctx context.Context, id string) ([]*domain.Session, error) {
	ret := _m.Called(ctx, id)

	var r0 []*domain.Session
	if rf, ok := ret.Get(0).(func(context.Context, string) []*domain.Session); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSession provides a mock function with given fields: ctx, session
func (_m *SessionStorage) UpdateSession(ctx context.Context, session *domain.Session) error {
	ret := _m.Called(ctx, session)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Session) error); ok {
		r0 = rf(ctx, session)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}