// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	config "git.jetbrains.space/orbi/fcsd/auth/internal/config"

	mock "github.com/stretchr/testify/mock"

	protonotification "git.jetbrains.space/orbi/fcsd/proto/notification"
)

// NotificationAdapter is an autogenerated mock type for the NotificationAdapter type
type NotificationAdapter struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *NotificationAdapter) Close() {
	_m.Called()
}

// Init provides a mock function with given fields: cfg
func (_m *NotificationAdapter) Init(cfg *config.Adapter) error {
	ret := _m.Called(cfg)

	var r0 error
	if rf, ok := ret.Get(0).(func(*config.Adapter) error); ok {
		r0 = rf(cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendNotify provides a mock function with given fields: ctx, notifToken
func (_m *NotificationAdapter) SendNotify(ctx context.Context, notifToken string) (*protonotification.MessageResponse, error) {
	ret := _m.Called(ctx, notifToken)

	var r0 *protonotification.MessageResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) *protonotification.MessageResponse); ok {
		r0 = rf(ctx, notifToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protonotification.MessageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, notifToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
