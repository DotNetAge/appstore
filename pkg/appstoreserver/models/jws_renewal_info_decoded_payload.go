// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// JWSRenewalInfoDecodedPayload represents a decoded payload containing subscription renewal information for an auto-renewable subscription
// https://developer.apple.com/documentation/appstoreserverapi/jwsrenewalinfodecodedpayload
type JWSRenewalInfoDecodedPayload struct {
	// ExpirationIntent is the reason the subscription expired
	// https://developer.apple.com/documentation/appstoreserverapi/expirationintent
	ExpirationIntent *ExpirationIntent `json:"expirationIntent,omitempty"`

	// OriginalTransactionId is the original transaction identifier of a purchase
	// https://developer.apple.com/documentation/appstoreserverapi/originaltransactionid
	OriginalTransactionId *string `json:"originalTransactionId,omitempty"`

	// AutoRenewProductId is the product identifier of the product that will renew at the next billing period
	// https://developer.apple.com/documentation/appstoreserverapi/autorenewproductid
	AutoRenewProductId *string `json:"autoRenewProductId,omitempty"`

	// ProductId is the unique identifier for the product, that you create in App Store Connect
	// https://developer.apple.com/documentation/appstoreserverapi/productid
	ProductId *string `json:"productId,omitempty"`

	// AutoRenewStatus is the renewal status of the auto-renewable subscription
	// https://developer.apple.com/documentation/appstoreserverapi/autorenewstatus
	AutoRenewStatus *AutoRenewStatus `json:"autoRenewStatus,omitempty"`

	// IsInBillingRetryPeriod is a Boolean value that indicates whether the App Store is attempting to automatically renew an expired subscription
	// https://developer.apple.com/documentation/appstoreserverapi/isinbillingretryperiod
	IsInBillingRetryPeriod *bool `json:"isInBillingRetryPeriod,omitempty"`

	// PriceIncreaseStatus is the status that indicates whether the auto-renewable subscription is subject to a price increase
	// https://developer.apple.com/documentation/appstoreserverapi/priceincreasestatus
	PriceIncreaseStatus *PriceIncreaseStatus `json:"priceIncreaseStatus,omitempty"`

	// GracePeriodExpiresDate is the time when the billing grace period for subscription renewals expires
	// https://developer.apple.com/documentation/appstoreserverapi/graceperiodexpiresdate
	GracePeriodExpiresDate *int64 `json:"gracePeriodExpiresDate,omitempty"`

	// OfferType is the type of the subscription offer
	// https://developer.apple.com/documentation/appstoreserverapi/offertype
	OfferType *OfferType `json:"offerType,omitempty"`

	// OfferIdentifier is the identifier that contains the promo code or the promotional offer identifier
	// https://developer.apple.com/documentation/appstoreserverapi/offeridentifier
	OfferIdentifier *string `json:"offerIdentifier,omitempty"`

	// SignedDate is the UNIX time, in milliseconds, that the App Store signed the JSON Web Signature data
	// https://developer.apple.com/documentation/appstoreserverapi/signeddate
	SignedDate *int64 `json:"signedDate,omitempty"`

	// Environment is the server environment, either sandbox or production
	// https://developer.apple.com/documentation/appstoreserverapi/environment
	Environment *Environment `json:"environment,omitempty"`

	// RecentSubscriptionStartDate is the earliest start date of a subscription in a series of auto-renewable subscription purchases that ignores all lapses of paid service shorter than 60 days
	// https://developer.apple.com/documentation/appstoreserverapi/recentsubscriptionstartdate
	RecentSubscriptionStartDate *int64 `json:"recentSubscriptionStartDate,omitempty"`

	// RenewalDate is the UNIX time, in milliseconds, that the most recent auto-renewable subscription purchase expires
	// https://developer.apple.com/documentation/appstoreserverapi/renewaldate
	RenewalDate *int64 `json:"renewalDate,omitempty"`

	// Currency is the currency code for the renewalPrice of the subscription
	// https://developer.apple.com/documentation/appstoreserverapi/currency
	Currency *string `json:"currency,omitempty"`

	// RenewalPrice is the renewal price, in milliunits, of the auto-renewable subscription that renews at the next billing period
	// https://developer.apple.com/documentation/appstoreserverapi/renewalprice
	RenewalPrice *int64 `json:"renewalPrice,omitempty"`

	// OfferDiscountType is the payment mode of the discount offer
	// https://developer.apple.com/documentation/appstoreserverapi/offerdiscounttype
	OfferDiscountType *OfferDiscountType `json:"offerDiscountType,omitempty"`

	// EligibleWinBackOfferIds is an array of win-back offer identifiers that a customer is eligible to redeem, which sorts the identifiers to present the better offers first
	// https://developer.apple.com/documentation/appstoreserverapi/eligiblewinbackofferids
	EligibleWinBackOfferIds []string `json:"eligibleWinBackOfferIds,omitempty"`
}
