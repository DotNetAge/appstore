// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// ConsumptionStatus represents the extent to which the customer consumed the in-app purchase.
// https://developer.apple.com/documentation/appstoreserverapi/consumptionstatus
type ConsumptionStatus int

const (
	// ConsumptionStatusUndeclared indicates the consumption status is undeclared
	ConsumptionStatusUndeclared ConsumptionStatus = 0
	// ConsumptionStatusNotConsumed indicates the in-app purchase has not been consumed
	ConsumptionStatusNotConsumed ConsumptionStatus = 1
	// ConsumptionStatusPartiallyConsumed indicates the in-app purchase has been partially consumed
	ConsumptionStatusPartiallyConsumed ConsumptionStatus = 2
	// ConsumptionStatusFullyConsumed indicates the in-app purchase has been fully consumed
	ConsumptionStatusFullyConsumed ConsumptionStatus = 3
)
