// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// SubscriptionGroupIdentifierItem represents information for auto-renewable subscriptions, including signed transaction information and signed renewal information, for one subscription group
// https://developer.apple.com/documentation/appstoreserverapi/subscriptiongroupidentifieritem
type SubscriptionGroupIdentifierItem struct {
	// SubscriptionGroupIdentifier is the identifier of the subscription group that the subscription belongs to
	// https://developer.apple.com/documentation/appstoreserverapi/subscriptiongroupidentifier
	SubscriptionGroupIdentifier *string `json:"subscriptionGroupIdentifier,omitempty"`

	// LastTransactions is an array of the most recent App Store-signed transaction information and App Store-signed renewal information for all auto-renewable subscriptions in the subscription group
	// https://developer.apple.com/documentation/appstoreserverapi/lasttransactionsitem
	LastTransactions []LastTransactionsItem `json:"lastTransactions,omitempty"`
}
