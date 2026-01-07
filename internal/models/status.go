// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// Status represents the status of an auto-renewable subscription.
// https://developer.apple.com/documentation/appstoreserverapi/status
type Status int

const (
	// StatusActive indicates the subscription is active
	StatusActive Status = 1
	// StatusExpired indicates the subscription has expired
	StatusExpired Status = 2
	// StatusBillingRetry indicates the subscription is in billing retry status
	StatusBillingRetry Status = 3
	// StatusBillingGracePeriod indicates the subscription is in billing grace period
	StatusBillingGracePeriod Status = 4
	// StatusRevoked indicates the subscription has been revoked
	StatusRevoked Status = 5
)
