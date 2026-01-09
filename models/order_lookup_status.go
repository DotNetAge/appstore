// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// OrderLookupStatus represents a value that indicates whether the order ID in the request is valid for your app
// https://developer.apple.com/documentation/appstoreserverapi/orderlookupstatus
type OrderLookupStatus int

const (
	// OrderLookupStatusValid indicates the order ID is valid
	OrderLookupStatusValid OrderLookupStatus = 0
	// OrderLookupStatusInvalid indicates the order ID is invalid
	OrderLookupStatusInvalid OrderLookupStatus = 1
)
