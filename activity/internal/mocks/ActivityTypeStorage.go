// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "git.jetbrains.space/orbi/fcsd/timesheet/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// EventStorage is an autogenerated mock type for the EventStorage type
type EventStorage struct {
	mock.Mock
}

// CreateEvent provides a mock function with given fields: ctx, Event
func (_m *EventStorage) CreateEvent(ctx context.Context, Event *domain.Event) (*domain.Event, error) {
	ret := _m.Called(ctx, Event)

	var r0 *domain.Event
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Event) *domain.Event); ok {
		r0 = rf(ctx, Event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Event) error); ok {
		r1 = rf(ctx, Event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteEvent provides a mock function with given fields: ctx, id
func (_m *EventStorage) DeleteEvent(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetEvent provides a mock function with given fields: ctx, id
func (_m *EventStorage) GetEvent(ctx context.Context, id string) (bool, *domain.Event, error) {
	ret := _m.Called(ctx, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *domain.Event
	if rf, ok := ret.Get(1).(func(context.Context, string) *domain.Event); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.Event)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateEvent provides a mock function with given fields: ctx, Event
func (_m *EventStorage) UpdateEvent(ctx context.Context, Event *domain.Event) (*domain.Event, error) {
	ret := _m.Called(ctx, Event)

	var r0 *domain.Event
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Event) *domain.Event); ok {
		r0 = rf(ctx, Event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Event) error); ok {
		r1 = rf(ctx, Event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
