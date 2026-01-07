// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// StatusResponse represents a response that contains status information for all of a customer's auto-renewable subscriptions in your app
// https://developer.apple.com/documentation/appstoreserverapi/statusresponse
type StatusResponse struct {
	// Environment is the server environment, sandbox or production, in which the App Store generated the response
	// https://developer.apple.com/documentation/appstoreserverapi/environment
	Environment *Environment `json:"environment,omitempty"`

	// BundleId is the bundle identifier of an app
	// https://developer.apple.com/documentation/appstoreserverapi/bundleid
	BundleId *string `json:"bundleId,omitempty"`

	// AppAppleId is the unique identifier of an app in the App Store
	// https://developer.apple.com/documentation/appstoreservernotifications/appappleid
	AppAppleId *int64 `json:"appAppleId,omitempty"`

	// Data is an array of information for auto-renewable subscriptions, including App Store-signed transaction information and App Store-signed renewal information
	// https://developer.apple.com/documentation/appstoreserverapi/subscriptiongroupidentifieritem
	Data []SubscriptionGroupIdentifierItem `json:"data,omitempty"`
}
