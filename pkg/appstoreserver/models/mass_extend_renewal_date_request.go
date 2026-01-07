// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// MassExtendRenewalDateRequest represents the request body that contains subscription-renewal-extension data to apply for all eligible active subscribers
// https://developer.apple.com/documentation/appstoreserverapi/massextendrenewaldaterequest
type MassExtendRenewalDateRequest struct {
	// ExtendByDays is the number of days to extend the subscription renewal date
	// maximum: 90
	// https://developer.apple.com/documentation/appstoreserverapi/extendbydays
	ExtendByDays *int `json:"extendByDays,omitempty"`

	// ExtendReasonCode is the reason code for the subscription-renewal-date extension
	// https://developer.apple.com/documentation/appstoreserverapi/extendreasoncode
	ExtendReasonCode *ExtendReasonCode `json:"extendReasonCode,omitempty"`

	// RequestIdentifier is a unique identifier to track each subscription-renewal-date extension request
	// https://developer.apple.com/documentation/appstoreserverapi/requestidentifier
	RequestIdentifier *string `json:"requestIdentifier,omitempty"`

	// StorefrontCountryCodes is a list of storefront country codes to limit the storefronts for a subscription-renewal-date extension
	// https://developer.apple.com/documentation/appstoreserverapi/storefrontcountrycodes
	StorefrontCountryCodes []string `json:"storefrontCountryCodes,omitempty"`

	// ProductId is the unique identifier for the product, that you create in App Store Connect
	// https://developer.apple.com/documentation/appstoreserverapi/productid
	ProductId *string `json:"productId,omitempty"`
}
