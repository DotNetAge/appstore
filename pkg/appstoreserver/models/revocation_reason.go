// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// RevocationReason represents the reason for a refunded transaction
// https://developer.apple.com/documentation/appstoreserverapi/revocationreason
type RevocationReason int

const (
	// RevocationReasonRefundedForOtherReason indicates the transaction was refunded for another reason
	RevocationReasonRefundedForOtherReason RevocationReason = 0
	// RevocationReasonRefundedDueToIssue indicates the transaction was refunded due to an issue
	RevocationReasonRefundedDueToIssue RevocationReason = 1
)
