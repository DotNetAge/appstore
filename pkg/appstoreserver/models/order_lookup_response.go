// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// OrderLookupResponse represents a response that includes the order lookup status and an array of signed transactions for the in-app purchases in the order
// https://developer.apple.com/documentation/appstoreserverapi/orderlookupresponse
type OrderLookupResponse struct {
	// Status is the status that indicates whether the order ID is valid
	// https://developer.apple.com/documentation/appstoreserverapi/orderlookupstatus
	Status *OrderLookupStatus `json:"status,omitempty"`

	// SignedTransactions is an array of in-app purchase transactions that are part of order, signed by Apple, in JSON Web Signature format
	// https://developer.apple.com/documentation/appstoreserverapi/jwstransaction
	SignedTransactions []string `json:"signedTransactions,omitempty"`
}
