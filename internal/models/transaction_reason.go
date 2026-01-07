// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// TransactionReason represents the cause of a purchase transaction, which indicates whether it's a customer's purchase or a renewal for an auto-renewable subscription that the system initiates
// https://developer.apple.com/documentation/appstoreserverapi/transactionreason
type TransactionReason string

const (
	// TransactionReasonPurchase indicates a customer-initiated purchase
	TransactionReasonPurchase TransactionReason = "PURCHASE"
	// TransactionReasonRenewal indicates a system-initiated renewal for an auto-renewable subscription
	TransactionReasonRenewal TransactionReason = "RENEWAL"
)
