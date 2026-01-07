// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// Data represents the app metadata and the signed renewal and transaction information
// https://developer.apple.com/documentation/appstoreservernotifications/data
type Data struct {
	// Environment is the server environment that the notification applies to, either sandbox or production
	// https://developer.apple.com/documentation/appstoreservernotifications/environment
	Environment *Environment `json:"environment,omitempty"`

	// AppAppleId is the unique identifier of an app in the App Store
	// https://developer.apple.com/documentation/appstoreservernotifications/appappleid
	AppAppleId *int64 `json:"appAppleId,omitempty"`

	// BundleId is the bundle identifier of an app
	// https://developer.apple.com/documentation/appstoreserverapi/bundleid
	BundleId *string `json:"bundleId,omitempty"`

	// BundleVersion is the version of the build that identifies an iteration of the bundle
	// https://developer.apple.com/documentation/appstoreservernotifications/bundleversion
	BundleVersion *string `json:"bundleVersion,omitempty"`

	// SignedTransactionInfo is transaction information signed by the App Store, in JSON Web Signature (JWS) format
	// https://developer.apple.com/documentation/appstoreserverapi/jwstransaction
	SignedTransactionInfo *string `json:"signedTransactionInfo,omitempty"`

	// SignedRenewalInfo is subscription renewal information, signed by the App Store, in JSON Web Signature (JWS) format
	// https://developer.apple.com/documentation/appstoreserverapi/jwsrenewalinfo
	SignedRenewalInfo *string `json:"signedRenewalInfo,omitempty"`

	// Status is the status of an auto-renewable subscription as of the signedDate in the responseBodyV2DecodedPayload
	// https://developer.apple.com/documentation/appstoreservernotifications/status
	Status *Status `json:"status,omitempty"`

	// ConsumptionRequestReason is the reason the customer requested the refund
	// https://developer.apple.com/documentation/appstoreservernotifications/consumptionrequestreason
	ConsumptionRequestReason *ConsumptionRequestReason `json:"consumptionRequestReason,omitempty"`
}
