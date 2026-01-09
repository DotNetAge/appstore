// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// HistoryResponse represents a response that contains the customer's transaction history for an app
// https://developer.apple.com/documentation/appstoreserverapi/historyresponse
type HistoryResponse struct {
	// Revision is a token you use in a query to request the next set of transactions for the customer
	// https://developer.apple.com/documentation/appstoreserverapi/revision
	Revision *string `json:"revision,omitempty"`

	// HasMore is a Boolean value indicating whether the App Store has more transaction data
	// https://developer.apple.com/documentation/appstoreserverapi/hasmore
	HasMore *bool `json:"hasMore,omitempty"`

	// BundleId is the bundle identifier of an app
	// https://developer.apple.com/documentation/appstoreserverapi/bundleid
	BundleId *string `json:"bundleId,omitempty"`

	// AppAppleId is the unique identifier of an app in the App Store
	// https://developer.apple.com/documentation/appstoreservernotifications/appappleid
	AppAppleId *int64 `json:"appAppleId,omitempty"`

	// Environment is the server environment in which you're making the request, whether sandbox or production
	// https://developer.apple.com/documentation/appstoreserverapi/environment
	Environment *Environment `json:"environment,omitempty"`

	// SignedTransactions is an array of in-app purchase transactions for the customer, signed by Apple, in JSON Web Signature format
	// https://developer.apple.com/documentation/appstoreserverapi/jwstransaction
	SignedTransactions []string `json:"signedTransactions,omitempty"`
}
