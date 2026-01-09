// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// MassExtendRenewalDateResponse represents a response that indicates the server successfully received the subscription-renewal-date extension request
// https://developer.apple.com/documentation/appstoreserverapi/massextendrenewaldateresponse
type MassExtendRenewalDateResponse struct {
	// RequestIdentifier is a unique identifier to track each subscription-renewal-date extension request
	// https://developer.apple.com/documentation/appstoreserverapi/requestidentifier
	RequestIdentifier *string `json:"requestIdentifier,omitempty"`
}
