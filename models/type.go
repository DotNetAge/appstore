// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// Type represents the type of in-app purchase products you can offer in your app
// https://developer.apple.com/documentation/appstoreserverapi/type
type Type string

const (
	// TypeAutoRenewableSubscription indicates an auto-renewable subscription
	TypeAutoRenewableSubscription Type = "Auto-Renewable Subscription"
	// TypeNonConsumable indicates a non-consumable in-app purchase
	TypeNonConsumable Type = "Non-Consumable"
	// TypeConsumable indicates a consumable in-app purchase
	TypeConsumable Type = "Consumable"
	// TypeNonRenewingSubscription indicates a non-renewing subscription
	TypeNonRenewingSubscription Type = "Non-Renewing Subscription"
)
