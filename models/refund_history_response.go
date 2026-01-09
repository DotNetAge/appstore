// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// RefundHistoryResponse represents a response that contains an array of signed JSON Web Signature (JWS) refunded transactions, and paging information
// https://developer.apple.com/documentation/appstoreserverapi/refundhistoryresponse
type RefundHistoryResponse struct {
	// SignedTransactions is a list of up to 20 JWS transactions, or an empty array if the customer hasn't received any refunds in your app
	// The transactions are sorted in ascending order by revocationDate
	// https://developer.apple.com/documentation/appstoreserverapi/jwstransaction
	SignedTransactions []string `json:"signedTransactions,omitempty"`

	// Revision is a token you use in a query to request the next set of transactions for the customer
	// https://developer.apple.com/documentation/appstoreserverapi/revision
	Revision *string `json:"revision,omitempty"`

	// HasMore is a Boolean value indicating whether the App Store has more transaction data
	// https://developer.apple.com/documentation/appstoreserverapi/hasmore
	HasMore *bool `json:"hasMore,omitempty"`
}
