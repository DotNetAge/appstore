// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// ConsumptionRequestReason represents the reason the customer requested the refund
// https://developer.apple.com/documentation/appstoreservernotifications/consumptionrequestreason
type ConsumptionRequestReason string

const (
	// ConsumptionRequestReasonAccidentalPurchase indicates an accidental purchase
	ConsumptionRequestReasonAccidentalPurchase ConsumptionRequestReason = "ACCIDENTAL_PURCHASE"
	// ConsumptionRequestReasonOther indicates other reason
	ConsumptionRequestReasonOther ConsumptionRequestReason = "OTHER"
)
