// Copyright (c) 2024 Apple Inc. Licensed under MIT License.

package models

// ExternalPurchaseToken represents the payload data that contains an external purchase token
// https://developer.apple.com/documentation/appstoreservernotifications/externalpurchasetoken
type ExternalPurchaseToken struct {
	// ExternalPurchaseId is the field of an external purchase token that uniquely identifies the token
	// https://developer.apple.com/documentation/appstoreservernotifications/externalpurchaseid
	ExternalPurchaseId *string `json:"externalPurchaseId,omitempty"`

	// TokenCreationDate is the field of an external purchase token that contains the UNIX date, in milliseconds, when the system created the token
	// https://developer.apple.com/documentation/appstoreservernotifications/tokencreationdate
	TokenCreationDate *int64 `json:"tokenCreationDate,omitempty"`

	// AppAppleId is the unique identifier of an app in the App Store
	// https://developer.apple.com/documentation/appstoreservernotifications/appappleid
	AppAppleId *int64 `json:"appAppleId,omitempty"`

	// BundleId is the bundle identifier of an app
	// https://developer.apple.com/documentation/appstoreservernotifications/bundleid
	BundleId *string `json:"bundleId,omitempty"`
}
