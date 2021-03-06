// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/turbulent376/homeactivity/activity/internal/domain"
)

// ActivityStorage is an autogenerated mock type for the ActivityStorage type
type ActivityStorage struct {
	mock.Mock
}

// CreateActivity provides a mock function with given fields: ctx, activity
func (_m *ActivityStorage) CreateActivity(ctx context.Context, activity *domain.Activity) (*domain.Activity, error) {
	ret := _m.Called(ctx, activity)

	var r0 *domain.Activity
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Activity) *domain.Activity); ok {
		r0 = rf(ctx, activity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Activity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Activity) error); ok {
		r1 = rf(ctx, activity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteActivity provides a mock function with given fields: ctx, id
func (_m *ActivityStorage) DeleteActivity(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetActivity provides a mock function with given fields: ctx, id
func (_m *ActivityStorage) GetActivity(ctx context.Context, id string) (bool, *domain.Activity, error) {
	ret := _m.Called(ctx, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *domain.Activity
	if rf, ok := ret.Get(1).(func(context.Context, string) *domain.Activity); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.Activity)
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

// ListActivities provides a mock function with given fields: ctx, userId
func (_m *ActivityStorage) ListActivities(ctx context.Context, userId string) (bool, []*domain.Activity, error) {
	ret := _m.Called(ctx, userId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 []*domain.Activity
	if rf, ok := ret.Get(1).(func(context.Context, string) []*domain.Activity); ok {
		r1 = rf(ctx, userId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*domain.Activity)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, userId)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListActivitiesByFamily provides a mock function with given fields: ctx, familyId
func (_m *ActivityStorage) ListActivitiesByFamily(ctx context.Context, familyId string) (bool, []*domain.Activity, error) {
	ret := _m.Called(ctx, familyId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, familyId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 []*domain.Activity
	if rf, ok := ret.Get(1).(func(context.Context, string) []*domain.Activity); ok {
		r1 = rf(ctx, familyId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*domain.Activity)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, familyId)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateActivity provides a mock function with given fields: ctx, activity
func (_m *ActivityStorage) UpdateActivity(ctx context.Context, activity *domain.Activity) (*domain.Activity, error) {
	ret := _m.Called(ctx, activity)

	var r0 *domain.Activity
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Activity) *domain.Activity); ok {
		r0 = rf(ctx, activity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Activity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Activity) error); ok {
		r1 = rf(ctx, activity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
