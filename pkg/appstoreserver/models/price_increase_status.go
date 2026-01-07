// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// PriceIncreaseStatus represents the status that indicates whether an auto-renewable subscription is subject to a price increase
// https://developer.apple.com/documentation/appstoreserverapi/priceincreasestatus
type PriceIncreaseStatus int

const (
	// PriceIncreaseStatusCustomerHasNotResponded indicates the customer has not responded to the price increase
	PriceIncreaseStatusCustomerHasNotResponded PriceIncreaseStatus = 0
	// PriceIncreaseStatusCustomerConsentedOrWasNotifiedWithoutNeedingConsent indicates the customer consented to the price increase or was notified without needing consent
	PriceIncreaseStatusCustomerConsentedOrWasNotifiedWithoutNeedingConsent PriceIncreaseStatus = 1
)
