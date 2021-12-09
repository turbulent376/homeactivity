// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "git.jetbrains.space/orbi/fcsd/timesheet/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// TimesheetService is an autogenerated mock type for the TimesheetService type
type ActivityService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, Timesheet
func (_m *ActivityService) Create(ctx context.Context, activity *domain.Activity) (*domain.Activity, error) {
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

// CreateEvent provides a mock function with given fields: ctx, Event
func (_m *ActivityService) CreateEvent(ctx context.Context, at *domain.ActivityType) (*domain.ActivityType, error) {
	ret := _m.Called(ctx, at)

	var r0 *domain.ActivityType
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ActivityType) *domain.ActivityType); ok {
		r0 = rf(ctx, at)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ActivityType)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.ActivityType) error); ok {
		r1 = rf(ctx, at)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ActivityService) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteEvent provides a mock function with given fields: ctx, id
func (_m *TimesheetService) DeleteEvent(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *TimesheetService) Get(ctx context.Context, id string) (bool, *domain.Timesheet, error) {
	ret := _m.Called(ctx, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *domain.Timesheet
	if rf, ok := ret.Get(1).(func(context.Context, string) *domain.Timesheet); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.Timesheet)
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

// GetEvent provides a mock function with given fields: ctx, id
func (_m *TimesheetService) GetEvent(ctx context.Context, id string) (bool, *domain.Event, error) {
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

// Search provides a mock function with given fields: ctx, rq
func (_m *TimesheetService) Search(ctx context.Context, rq *domain.SearchTimesheetRequest) (bool, *domain.Timesheet, error) {
	ret := _m.Called(ctx, rq)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *domain.SearchTimesheetRequest) bool); ok {
		r0 = rf(ctx, rq)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *domain.Timesheet
	if rf, ok := ret.Get(1).(func(context.Context, *domain.SearchTimesheetRequest) *domain.Timesheet); ok {
		r1 = rf(ctx, rq)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.Timesheet)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *domain.SearchTimesheetRequest) error); ok {
		r2 = rf(ctx, rq)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: ctx, Timesheet
func (_m *TimesheetService) Update(ctx context.Context, Timesheet *domain.Timesheet) (*domain.Timesheet, error) {
	ret := _m.Called(ctx, Timesheet)

	var r0 *domain.Timesheet
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Timesheet) *domain.Timesheet); ok {
		r0 = rf(ctx, Timesheet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Timesheet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Timesheet) error); ok {
		r1 = rf(ctx, Timesheet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEvent provides a mock function with given fields: ctx, Event
func (_m *TimesheetService) UpdateEvent(ctx context.Context, Event *domain.Event) (*domain.Event, error) {
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