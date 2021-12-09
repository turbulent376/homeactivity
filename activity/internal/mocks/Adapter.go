// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"

	config "git.jetbrains.space/orbi/fcsd/timesheet/internal/config"

	domain "git.jetbrains.space/orbi/fcsd/timesheet/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// Adapter is an autogenerated mock type for the Adapter type
type Adapter struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Adapter) Close() {
	_m.Called()
}

// CreateEvent provides a mock function with given fields: ctx, Event
func (_m *Adapter) CreateEvent(ctx context.Context, Event *domain.Event) (*domain.Event, error) {
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

// CreateTimetable provides a mock function with given fields: ctx, Timesheet
func (_m *Adapter) CreateTimetable(ctx context.Context, Timesheet *domain.Timesheet) (*domain.Timesheet, error) {
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

// DeleteEvent provides a mock function with given fields: ctx, id
func (_m *Adapter) DeleteEvent(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTimetable provides a mock function with given fields: ctx, id
func (_m *Adapter) DeleteTimetable(ctx context.Context, id string) error {
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
func (_m *Adapter) GetEvent(ctx context.Context, id string) (bool, *domain.Event, error) {
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

// GetTimetable provides a mock function with given fields: ctx, id
func (_m *Adapter) GetTimetable(ctx context.Context, id string) (bool, *domain.Timesheet, error) {
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

// Init provides a mock function with given fields: cfg
func (_m *Adapter) Init(cfg *config.Storages) error {
	ret := _m.Called(cfg)

	var r0 error
	if rf, ok := ret.Get(0).(func(*config.Storages) error); ok {
		r0 = rf(cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchTimetable provides a mock function with given fields: ctx, rq
func (_m *Adapter) SearchTimetable(ctx context.Context, rq *domain.SearchTimesheetRequest) (*domain.Timesheet, error) {
	ret := _m.Called(ctx, rq)

	var r0 *domain.Timesheet
	if rf, ok := ret.Get(0).(func(context.Context, *domain.SearchTimesheetRequest) *domain.Timesheet); ok {
		r0 = rf(ctx, rq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Timesheet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.SearchTimesheetRequest) error); ok {
		r1 = rf(ctx, rq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEvent provides a mock function with given fields: ctx, Event
func (_m *Adapter) UpdateEvent(ctx context.Context, Event *domain.Event) (*domain.Event, error) {
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

// UpdateTimetable provides a mock function with given fields: ctx, Timesheet
func (_m *Adapter) UpdateTimetable(ctx context.Context, Timesheet *domain.Timesheet) (*domain.Timesheet, error) {
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
