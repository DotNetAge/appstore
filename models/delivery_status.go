// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// DeliveryStatus represents whether the app successfully delivered an in-app purchase that works properly.
// https://developer.apple.com/documentation/appstoreserverapi/deliverystatus
type DeliveryStatus int

const (
	// DeliveryStatusDeliveredAndWorkingProperly indicates the app successfully delivered an in-app purchase that works properly
	DeliveryStatusDeliveredAndWorkingProperly DeliveryStatus = 0
	// DeliveryStatusDidNotDeliverDueToQualityIssue indicates the app did not deliver the in-app purchase due to a quality issue
	DeliveryStatusDidNotDeliverDueToQualityIssue DeliveryStatus = 1
	// DeliveryStatusDeliveredWrongItem indicates the app delivered the wrong item
	DeliveryStatusDeliveredWrongItem DeliveryStatus = 2
	// DeliveryStatusDidNotDeliverDueToServerOutage indicates the app did not deliver the in-app purchase due to a server outage
	DeliveryStatusDidNotDeliverDueToServerOutage DeliveryStatus = 3
	// DeliveryStatusDidNotDeliverDueToInGameCurrencyChange indicates the app did not deliver the in-app purchase due to an in-game currency change
	DeliveryStatusDidNotDeliverDueToInGameCurrencyChange DeliveryStatus = 4
	// DeliveryStatusDidNotDeliverForOtherReason indicates the app did not deliver the in-app purchase for another reason
	DeliveryStatusDidNotDeliverForOtherReason DeliveryStatus = 5
)
