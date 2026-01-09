// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// UserStatus represents the status of a customer's account within your app.
// https://developer.apple.com/documentation/appstoreserverapi/userstatus
type UserStatus int

const (
	// UserStatusUndeclared indicates the user status is undeclared
	UserStatusUndeclared UserStatus = 0
	// UserStatusActive indicates the user account is active
	UserStatusActive UserStatus = 1
	// UserStatusSuspended indicates the user account is suspended
	UserStatusSuspended UserStatus = 2
	// UserStatusTerminated indicates the user account is terminated
	UserStatusTerminated UserStatus = 3
	// UserStatusLimitedAccess indicates the user account has limited access
	UserStatusLimitedAccess UserStatus = 4
)
