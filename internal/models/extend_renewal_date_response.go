// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// ExtendRenewalDateResponse represents a response that indicates whether an individual renewal-date extension succeeded, and related details
// https://developer.apple.com/documentation/appstoreserverapi/extendrenewaldateresponse
type ExtendRenewalDateResponse struct {
	// OriginalTransactionId is the original transaction identifier of a purchase
	// https://developer.apple.com/documentation/appstoreserverapi/originaltransactionid
	OriginalTransactionId *string `json:"originalTransactionId,omitempty"`

	// WebOrderLineItemId is the unique identifier of subscription-purchase events across devices, including renewals
	// https://developer.apple.com/documentation/appstoreserverapi/weborderlineitemid
	WebOrderLineItemId *string `json:"webOrderLineItemId,omitempty"`

	// Success is a Boolean value that indicates whether the subscription-renewal-date extension succeeded
	// https://developer.apple.com/documentation/appstoreserverapi/success
	Success *bool `json:"success,omitempty"`

	// EffectiveDate is the new subscription expiration date for a subscription-renewal extension
	// https://developer.apple.com/documentation/appstoreserverapi/effectivedate
	EffectiveDate *int64 `json:"effectiveDate,omitempty"`
}
