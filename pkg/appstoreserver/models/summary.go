// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// Summary represents the payload data for a subscription-renewal-date extension notification
// https://developer.apple.com/documentation/appstoreservernotifications/summary
type Summary struct {
	// Environment is the server environment that the notification applies to, either sandbox or production
	// https://developer.apple.com/documentation/appstoreservernotifications/environment
	Environment *Environment `json:"environment,omitempty"`

	// AppAppleId is the unique identifier of an app in the App Store
	// https://developer.apple.com/documentation/appstoreservernotifications/appappleid
	AppAppleId *int64 `json:"appAppleId,omitempty"`

	// BundleId is the bundle identifier of an app
	// https://developer.apple.com/documentation/appstoreserverapi/bundleid
	BundleId *string `json:"bundleId,omitempty"`

	// ProductId is the unique identifier for the product, that you create in App Store Connect
	// https://developer.apple.com/documentation/appstoreserverapi/productid
	ProductId *string `json:"productId,omitempty"`

	// RequestIdentifier is a string that contains a unique identifier you provide to track each subscription-renewal-date extension request
	// https://developer.apple.com/documentation/appstoreserverapi/requestidentifier
	RequestIdentifier *string `json:"requestIdentifier,omitempty"`

	// StorefrontCountryCodes is a list of storefront country codes you provide to limit the storefronts for a subscription-renewal-date extension
	// https://developer.apple.com/documentation/appstoreserverapi/storefrontcountrycodes
	StorefrontCountryCodes []string `json:"storefrontCountryCodes,omitempty"`

	// SucceededCount is the count of subscriptions that successfully receive a subscription-renewal-date extension
	// https://developer.apple.com/documentation/appstoreserverapi/succeededcount
	SucceededCount *int `json:"succeededCount,omitempty"`

	// FailedCount is the count of subscriptions that fail to receive a subscription-renewal-date extension
	// https://developer.apple.com/documentation/appstoreserverapi/failedcount
	FailedCount *int `json:"failedCount,omitempty"`
}
