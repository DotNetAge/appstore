// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// ExpirationIntent represents the reason an auto-renewable subscription expired
// https://developer.apple.com/documentation/appstoreserverapi/expirationintent
type ExpirationIntent int

const (
	// ExpirationIntentCustomerCancelled indicates the customer canceled the subscription
	ExpirationIntentCustomerCancelled ExpirationIntent = 1
	// ExpirationIntentBillingError indicates the subscription expired due to a billing error
	ExpirationIntentBillingError ExpirationIntent = 2
	// ExpirationIntentCustomerDidNotConsentToPriceIncrease indicates the customer did not consent to a price increase
	ExpirationIntentCustomerDidNotConsentToPriceIncrease ExpirationIntent = 3
	// ExpirationIntentProductNotAvailable indicates the product is no longer available
	ExpirationIntentProductNotAvailable ExpirationIntent = 4
	// ExpirationIntentOther indicates the subscription expired for another reason
	ExpirationIntentOther ExpirationIntent = 5
)
