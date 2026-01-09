// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// MassExtendRenewalDateStatusResponse represents a response that indicates the current status of a request to extend the subscription renewal date to all eligible subscribers
// https://developer.apple.com/documentation/appstoreserverapi/massextendrenewaldatestatusresponse
type MassExtendRenewalDateStatusResponse struct {
	// RequestIdentifier is a unique identifier to track each subscription-renewal-date extension request
	// https://developer.apple.com/documentation/appstoreserverapi/requestidentifier
	RequestIdentifier *string `json:"requestIdentifier,omitempty"`

	// Complete is a Boolean value that indicates whether the App Store completed the request to extend a subscription renewal date to active subscribers
	// https://developer.apple.com/documentation/appstoreserverapi/complete
	Complete *bool `json:"complete,omitempty"`

	// CompleteDate is the UNIX time, in milliseconds, that the App Store completes a request to extend a subscription renewal date for eligible subscribers
	// https://developer.apple.com/documentation/appstoreserverapi/completedate
	CompleteDate *int64 `json:"completeDate,omitempty"`

	// SucceededCount is the count of subscriptions that successfully receive a subscription-renewal-date extension
	// https://developer.apple.com/documentation/appstoreserverapi/succeededcount
	SucceededCount *int `json:"succeededCount,omitempty"`

	// FailedCount is the count of subscriptions that fail to receive a subscription-renewal-date extension
	// https://developer.apple.com/documentation/appstoreserverapi/failedcount
	FailedCount *int `json:"failedCount,omitempty"`
}
