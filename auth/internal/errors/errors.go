// this file is generated by servgen util based on a template at 2021-06-26 10:37:24 +0300 MSK
package errors

import (
	"context"

	"git.jetbrains.space/orbi/fcsd/kit/er"
	pb "git.jetbrains.space/orbi/fcsd/proto/auth"
)

var (
	ErrUserInvalidEmail = func(ctx context.Context) error {
		return er.WithBuilder(pb.ErrCodeAuthInvalidEmail, "invalid email").C(ctx).Err()
	}
	ErrUserInvalidPassword = func(ctx context.Context) error {
		return er.WithBuilder(pb.ErrCodeAuthInvalidPassword, "invalid password").C(ctx).Err()
	}
	ErrUserIdEmpty = func(ctx context.Context) error {
		return er.WithBuilder(pb.ErrCodeAuthUserIdEmpty, "empty user id").C(ctx).Err()
	}
	ErrEmailAlreadyExist = func(ctx context.Context) error {
		return er.WithBuilder(pb.ErrCodeAuthInvalidEmail, "this email address is already being used").C(ctx).Err()
	}
	ErrGetUserByField = func(ctx context.Context, field string) error {
		return er.WithBuilder(pb.ErrCodeAuthUserEmptyField, "failed to find user by field").F(er.FF{"field": field}).C(ctx).Err()
	}
	ErrUserNotFound = func(ctx context.Context, id string) error {
		return er.WithBuilder(pb.ErrCodeAuthObjectNotFound, "user not found").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrUserEmptySession = func(ctx context.Context, id string) error {
		return er.WithBuilder(pb.ErrCodeAuthUserEmptyField, "user has not sessions").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrSessionNotFound = func(cause error, ctx context.Context) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthObjectNotFound, "session not found").C(ctx).Err()
	}
	ErrUserSetSession = func(cause error, ctx context.Context) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthObjectNotFound, "failed to set user new session").C(ctx).Err()
	}
	ErrParseJWT = func(cause error, ctx context.Context) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthToken, "failed to parse user jwt token").C(ctx).Err()
	}
	ErrUserCreate = func(ctx context.Context, id string) error {
		return er.WithBuilder(pb.ErrCodeAuthObjectCreate, "failed to create user").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrSessionCreate = func(ctx context.Context, id string) error {
		return er.WithBuilder(pb.ErrCodeAuthObjectCreate, "failed to create session").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrUserUpdate = func(ctx context.Context, id string) error {
		return er.WithBuilder(pb.ErrCodeAuthObjectUpdate, "failed to user update").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrSessionUpdate = func(ctx context.Context, id string) error {
		return er.WithBuilder(pb.ErrCodeAuthObjectUpdate, "failed to session update").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrUserDelete = func(ctx context.Context, id string) error {
		return er.WithBuilder(pb.ErrCodeAuthObjectDelete, "already deleted").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrSessionDelete = func(ctx context.Context, id string) error {
		return er.WithBuilder(pb.ErrCodeAuthObjectDelete, "already deleted").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrFirebaseVerifyToken = func(cause error, ctx context.Context) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthToken, "failed firebase verify token").C(ctx).Err()
	}
	ErrAuthSignToken = func(cause error, ctx context.Context) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthToken, "token can't sign").C(ctx).Err()
	}
	ErrAuthEmptyToken = func(ctx context.Context) error {
		return er.WithBuilder(pb.ErrCodeAuthToken, "empty field token").C(ctx).Err()
	}

	ErrAuthRefreshToken = func(cause error, ctx context.Context) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthToken, "token didn't refresh").C(ctx).Err()
	}
	ErrAuthInvalidRefreshToken = func(ctx context.Context) error {
		return er.WithBuilder(pb.ErrCodeAuthToken, "invalid refresh token").C(ctx).Err()
	}
	ErrUserStorageCreate = func(cause error, ctx context.Context) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthStorageGetDb, "failed to save user").C(ctx).Err()
	}
	ErrSessionStorageCreate = func(cause error, ctx context.Context) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthStorageGetDb, "failed to save session").C(ctx).Err()
	}
	ErrSessionStorageDelete = func(cause error, ctx context.Context) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthStorageGetDb, "failed to delete session").C(ctx).Err()
	}
	ErrUserStorageGetDb = func(cause error, ctx context.Context, id string) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthStorageGetDb, "failed to get user from db").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrUserStorageGetCache = func(cause error, ctx context.Context, id string) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthStorageGetCache, "failed to get user from cache").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrSessionStorageGetDb = func(cause error, ctx context.Context, id string) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthStorageGetDb, "failed to get sessions from db").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrSessionStorageGetCache = func(cause error, ctx context.Context, id string) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthStorageGetCache, "failed to get sessions from cache").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrSessionStorageSetCache = func(cause error, ctx context.Context, id string) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthStorageGetCache, "failed to save sessions to cache").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrUserStorageGetDbByEmail = func(cause error, ctx context.Context, id string) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthStorageGetDb, "failed to get user by email from db").F(er.FF{"email": id}).C(ctx).Err()
	}
	ErrUserStorageGetCacheByEmail = func(cause error, ctx context.Context, id string) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthStorageGetCache, "failed to get user by email from cache").F(er.FF{"email": id}).C(ctx).Err()
	}
	ErrUserStorageGetDbByToken = func(cause error, ctx context.Context, id string) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthStorageGetDb, "failed to get user by token from db").F(er.FF{"idToken": id}).C(ctx).Err()
	}
	ErrUserStorageGetCacheByToken = func(cause error, ctx context.Context, id string) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthStorageGetCache, "failed to get user by token from cache").F(er.FF{"idToken": id}).C(ctx).Err()
	}
	ErrUserStorageSetCache = func(cause error, ctx context.Context, id string) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthStorageSetCache, "failed to save user to cache").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrUserStorageUpdate = func(cause error, ctx context.Context, id string) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthStorageUpdate, "failed to update user").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrSessionStorageUpdate = func(cause error, ctx context.Context, id string) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthStorageUpdate, "failed to update session").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrUserSendNotification = func(cause error, ctx context.Context) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeAuthSendNotification, "user notification failed").C(ctx).Err()
	}
)
