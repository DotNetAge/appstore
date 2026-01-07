// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// Subtype represents a string that provides details about select notification types in version 2
// https://developer.apple.com/documentation/appstoreservernotifications/subtype
type Subtype string

const (
	// SubtypeInitialBuy indicates the initial purchase of a subscription
	// https://developer.apple.com/documentation/appstoreservernotifications/initial_buy
	SubtypeInitialBuy Subtype = "INITIAL_BUY"
	
	// SubtypeResubscribe indicates a resubscription to a subscription
	// https://developer.apple.com/documentation/appstoreservernotifications/resubscribe
	SubtypeResubscribe Subtype = "RESUBSCRIBE"
	
	// SubtypeDowngrade indicates a downgrade in subscription plan
	// https://developer.apple.com/documentation/appstoreservernotifications/downgrade
	SubtypeDowngrade Subtype = "DOWNGRADE"
	
	// SubtypeUpgrade indicates an upgrade in subscription plan
	// https://developer.apple.com/documentation/appstoreservernotifications/upgrade
	SubtypeUpgrade Subtype = "UPGRADE"
	
	// SubtypeAutoRenewEnabled indicates that auto-renewal has been enabled
	// https://developer.apple.com/documentation/appstoreservernotifications/auto_renew_enabled
	SubtypeAutoRenewEnabled Subtype = "AUTO_RENEW_ENABLED"
	
	// SubtypeAutoRenewDisabled indicates that auto-renewal has been disabled
	// https://developer.apple.com/documentation/appstoreservernotifications/auto_renew_disabled
	SubtypeAutoRenewDisabled Subtype = "AUTO_RENEW_DISABLED"
	
	// SubtypeVoluntary indicates a voluntary subscription cancellation
	// https://developer.apple.com/documentation/appstoreservernotifications/voluntary
	SubtypeVoluntary Subtype = "VOLUNTARY"
	
	// SubtypeBillingRetry indicates a billing retry attempt
	// https://developer.apple.com/documentation/appstoreservernotifications/billing_retry
	SubtypeBillingRetry Subtype = "BILLING_RETRY"
	
	// SubtypePriceIncrease indicates a price increase notification
	// https://developer.apple.com/documentation/appstoreservernotifications/price_increase
	SubtypePriceIncrease Subtype = "PRICE_INCREASE"
	
	// SubtypeGracePeriod indicates a subscription is in grace period
	// https://developer.apple.com/documentation/appstoreservernotifications/grace_period
	SubtypeGracePeriod Subtype = "GRACE_PERIOD"
	
	// SubtypePending indicates a pending subscription
	// https://developer.apple.com/documentation/appstoreservernotifications/pending
	SubtypePending Subtype = "PENDING"
	
	// SubtypeAccepted indicates an accepted subscription
	// https://developer.apple.com/documentation/appstoreservernotifications/accepted
	SubtypeAccepted Subtype = "ACCEPTED"
	
	// SubtypeBillingRecovery indicates a billing recovery attempt
	// https://developer.apple.com/documentation/appstoreservernotifications/billing_recovery
	SubtypeBillingRecovery Subtype = "BILLING_RECOVERY"
	
	// SubtypeProductNotForSale indicates the product is no longer for sale
	// https://developer.apple.com/documentation/appstoreservernotifications/product_not_for_sale
	SubtypeProductNotForSale Subtype = "PRODUCT_NOT_FOR_SALE"
	
	// SubtypeSummary indicates a summary notification
	// https://developer.apple.com/documentation/appstoreservernotifications/summary
	SubtypeSummary Subtype = "SUMMARY"
	
	// SubtypeFailure indicates a failure notification
	// https://developer.apple.com/documentation/appstoreservernotifications/failure
	SubtypeFailure Subtype = "FAILURE"
	
	// SubtypeUnreported indicates an unreported notification
	// https://developer.apple.com/documentation/appstoreservernotifications/unreported
	SubtypeUnreported Subtype = "UNREPORTED"
)
