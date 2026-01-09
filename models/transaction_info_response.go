// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// TransactionInfoResponse represents a response that contains signed transaction information for a single transaction
// https://developer.apple.com/documentation/appstoreserverapi/transactioninforesponse
type TransactionInfoResponse struct {
	// SignedTransactionInfo is a customer's in-app purchase transaction, signed by Apple, in JSON Web Signature (JWS) format
	// https://developer.apple.com/documentation/appstoreserverapi/jwstransaction
	SignedTransactionInfo *string `json:"signedTransactionInfo,omitempty"`
}
