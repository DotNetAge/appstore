// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// LastTransactionsItem represents the most recent App Store-signed transaction information and App Store-signed renewal information for an auto-renewable subscription
// https://developer.apple.com/documentation/appstoreserverapi/lasttransactionsitem
type LastTransactionsItem struct {
	// Status is the status of the auto-renewable subscription
	// https://developer.apple.com/documentation/appstoreserverapi/status
	Status *Status `json:"status,omitempty"`

	// OriginalTransactionId is the original transaction identifier of a purchase
	// https://developer.apple.com/documentation/appstoreserverapi/originaltransactionid
	OriginalTransactionId *string `json:"originalTransactionId,omitempty"`

	// SignedTransactionInfo is transaction information signed by the App Store, in JSON Web Signature (JWS) format
	// https://developer.apple.com/documentation/appstoreserverapi/jwstransaction
	SignedTransactionInfo *string `json:"signedTransactionInfo,omitempty"`

	// SignedRenewalInfo is subscription renewal information, signed by the App Store, in JSON Web Signature (JWS) format
	// https://developer.apple.com/documentation/appstoreserverapi/jwsrenewalinfo
	SignedRenewalInfo *string `json:"signedRenewalInfo,omitempty"`
}
