// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// ExtendRenewalDateRequest represents the request body that contains subscription-renewal-extension data for an individual subscription
// https://developer.apple.com/documentation/appstoreserverapi/extendrenewaldaterequest
type ExtendRenewalDateRequest struct {
	// ExtendByDays is the number of days to extend the subscription renewal date
	// maximum: 90
	// https://developer.apple.com/documentation/appstoreserverapi/extendbydays
	ExtendByDays *int `json:"extendByDays,omitempty"`

	// ExtendReasonCode is the reason code for the subscription date extension
	// https://developer.apple.com/documentation/appstoreserverapi/extendreasoncode
	ExtendReasonCode *ExtendReasonCode `json:"extendReasonCode,omitempty"`

	// RequestIdentifier is a unique identifier to track each subscription-renewal-date extension request
	// https://developer.apple.com/documentation/appstoreserverapi/requestidentifier
	RequestIdentifier *string `json:"requestIdentifier,omitempty"`
}
