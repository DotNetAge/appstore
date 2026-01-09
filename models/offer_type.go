// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// OfferType represents the type of subscription offer
// https://developer.apple.com/documentation/appstoreserverapi/offertype
type OfferType int

const (
	// OfferTypeIntroductoryOffer indicates an introductory offer
	OfferTypeIntroductoryOffer OfferType = 1
	// OfferTypePromotionalOffer indicates a promotional offer
	OfferTypePromotionalOffer OfferType = 2
	// OfferTypeSubscriptionOfferCode indicates a subscription offer code
	OfferTypeSubscriptionOfferCode OfferType = 3
	// OfferTypeWinBackOffer indicates a win-back offer
	OfferTypeWinBackOffer OfferType = 4
)
