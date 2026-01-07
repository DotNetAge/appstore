// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// AutoRenewStatus represents the renewal status for an auto-renewable subscription
// https://developer.apple.com/documentation/appstoreserverapi/autorenewstatus
type AutoRenewStatus int

const (
	// AutoRenewStatusOff indicates auto-renew is off
	AutoRenewStatusOff AutoRenewStatus = 0
	// AutoRenewStatusOn indicates auto-renew is on
	AutoRenewStatusOn AutoRenewStatus = 1
)
