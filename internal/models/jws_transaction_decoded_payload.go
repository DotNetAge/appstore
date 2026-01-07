// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// JWSTransactionDecodedPayload represents a decoded payload containing transaction information
// https://developer.apple.com/documentation/appstoreserverapi/jwstransactiondecodedpayload
type JWSTransactionDecodedPayload struct {
	// OriginalTransactionId is the original transaction identifier of a purchase
	// https://developer.apple.com/documentation/appstoreserverapi/originaltransactionid
	OriginalTransactionId *string `json:"originalTransactionId,omitempty"`

	// TransactionId is the unique identifier for a transaction such as an in-app purchase, restored in-app purchase, or subscription renewal
	// https://developer.apple.com/documentation/appstoreserverapi/transactionid
	TransactionId *string `json:"transactionId,omitempty"`

	// WebOrderLineItemId is the unique identifier of subscription-purchase events across devices, including renewals
	// https://developer.apple.com/documentation/appstoreserverapi/weborderlineitemid
	WebOrderLineItemId *string `json:"webOrderLineItemId,omitempty"`

	// BundleId is the bundle identifier of an app
	// https://developer.apple.com/documentation/appstoreserverapi/bundleid
	BundleId *string `json:"bundleId,omitempty"`

	// ProductId is the unique identifier for the product, that you create in App Store Connect
	// https://developer.apple.com/documentation/appstoreserverapi/productid
	ProductId *string `json:"productId,omitempty"`

	// SubscriptionGroupIdentifier is the identifier of the subscription group that the subscription belongs to
	// https://developer.apple.com/documentation/appstoreserverapi/subscriptiongroupidentifier
	SubscriptionGroupIdentifier *string `json:"subscriptionGroupIdentifier,omitempty"`

	// PurchaseDate is the time that the App Store charged the user's account for an in-app purchase, a restored in-app purchase, a subscription, or a subscription renewal after a lapse
	// https://developer.apple.com/documentation/appstoreserverapi/purchasedate
	PurchaseDate *int64 `json:"purchaseDate,omitempty"`

	// OriginalPurchaseDate is the purchase date of the transaction associated with the original transaction identifier
	// https://developer.apple.com/documentation/appstoreserverapi/originalpurchasedate
	OriginalPurchaseDate *int64 `json:"originalPurchaseDate,omitempty"`

	// ExpiresDate is the UNIX time, in milliseconds, an auto-renewable subscription expires or renews
	// https://developer.apple.com/documentation/appstoreserverapi/expiresdate
	ExpiresDate *int64 `json:"expiresDate,omitempty"`

	// Quantity is the number of consumable products purchased
	// https://developer.apple.com/documentation/appstoreserverapi/quantity
	Quantity *int `json:"quantity,omitempty"`

	// Type is the type of the in-app purchase
	// https://developer.apple.com/documentation/appstoreserverapi/type
	Type *Type `json:"type,omitempty"`

	// AppAccountToken is the UUID that an app optionally generates to map a customer's in-app purchase with its resulting App Store transaction
	// https://developer.apple.com/documentation/appstoreserverapi/appaccounttoken
	AppAccountToken *string `json:"appAccountToken,omitempty"`

	// InAppOwnershipType is a string that describes whether the transaction was purchased by the user, or is available to them through Family Sharing
	// https://developer.apple.com/documentation/appstoreserverapi/inappownershiptype
	InAppOwnershipType *InAppOwnershipType `json:"inAppOwnershipType,omitempty"`

	// SignedDate is the UNIX time, in milliseconds, that the App Store signed the JSON Web Signature data
	// https://developer.apple.com/documentation/appstoreserverapi/signeddate
	SignedDate *int64 `json:"signedDate,omitempty"`

	// RevocationReason is the reason that the App Store refunded the transaction or revoked it from family sharing
	// https://developer.apple.com/documentation/appstoreserverapi/revocationreason
	RevocationReason *RevocationReason `json:"revocationReason,omitempty"`

	// RevocationDate is the UNIX time, in milliseconds, that Apple Support refunded a transaction
	// https://developer.apple.com/documentation/appstoreserverapi/revocationdate
	RevocationDate *int64 `json:"revocationDate,omitempty"`

	// IsUpgraded is the Boolean value that indicates whether the user upgraded to another subscription
	// https://developer.apple.com/documentation/appstoreserverapi/isupgraded
	IsUpgraded *bool `json:"isUpgraded,omitempty"`

	// OfferType is a value that represents the promotional offer type
	// https://developer.apple.com/documentation/appstoreserverapi/offertype
	OfferType *OfferType `json:"offerType,omitempty"`

	// OfferIdentifier is the identifier that contains the promo code or the promotional offer identifier
	// https://developer.apple.com/documentation/appstoreserverapi/offeridentifier
	OfferIdentifier *string `json:"offerIdentifier,omitempty"`

	// Environment is the server environment, either sandbox or production
	// https://developer.apple.com/documentation/appstoreserverapi/environment
	Environment *Environment `json:"environment,omitempty"`

	// Storefront is the three-letter code that represents the country or region associated with the App Store storefront for the purchase
	// https://developer.apple.com/documentation/appstoreserverapi/storefront
	Storefront *string `json:"storefront,omitempty"`

	// StorefrontId is an Apple-defined value that uniquely identifies the App Store storefront associated with the purchase
	// https://developer.apple.com/documentation/appstoreserverapi/storefrontid
	StorefrontId *string `json:"storefrontId,omitempty"`

	// TransactionReason is the reason for the purchase transaction, which indicates whether it's a customer's purchase or a renewal for an auto-renewable subscription that the system initiates
	// https://developer.apple.com/documentation/appstoreserverapi/transactionreason
	TransactionReason *TransactionReason `json:"transactionReason,omitempty"`

	// Currency is the three-letter ISO 4217 currency code for the price of the product
	// https://developer.apple.com/documentation/appstoreserverapi/currency
	Currency *string `json:"currency,omitempty"`

	// Price is the price, in milliunits, of the in-app purchase or subscription offer that you configured in App Store Connect
	// https://developer.apple.com/documentation/appstoreserverapi/price
	Price *int64 `json:"price,omitempty"`

	// OfferDiscountType is the payment mode you configure for an introductory offer, promotional offer, or offer code on an auto-renewable subscription
	// https://developer.apple.com/documentation/appstoreserverapi/offerdiscounttype
	OfferDiscountType *OfferDiscountType `json:"offerDiscountType,omitempty"`
}
