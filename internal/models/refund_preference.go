// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// RefundPreference represents your preferred outcome for the refund request.
// https://developer.apple.com/documentation/appstoreserverapi/refundpreference
type RefundPreference int

const (
	// RefundPreferenceUndeclared indicates the refund preference is undeclared
	RefundPreferenceUndeclared RefundPreference = 0
	// RefundPreferencePreferGrant indicates you prefer to grant the refund request
	RefundPreferencePreferGrant RefundPreference = 1
	// RefundPreferencePreferDecline indicates you prefer to decline the refund request
	RefundPreferencePreferDecline RefundPreference = 2
	// RefundPreferenceNoPreference indicates you have no preference for the refund request
	RefundPreferenceNoPreference RefundPreference = 3
)
