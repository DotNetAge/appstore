// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// OfferDiscountType represents the payment mode you configure for an introductory offer, promotional offer, or offer code on an auto-renewable subscription
// https://developer.apple.com/documentation/appstoreserverapi/offerdiscounttype
type OfferDiscountType string

const (
	// OfferDiscountTypeFreeTrial indicates a free trial offer
	OfferDiscountTypeFreeTrial OfferDiscountType = "FREE_TRIAL"
	// OfferDiscountTypePayAsYouGo indicates a pay-as-you-go offer
	OfferDiscountTypePayAsYouGo OfferDiscountType = "PAY_AS_YOU_GO"
	// OfferDiscountTypePayUpFront indicates a pay-up-front offer
	OfferDiscountTypePayUpFront OfferDiscountType = "PAY_UP_FRONT"
)
