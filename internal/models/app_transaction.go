// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// AppTransaction represents information that represents the customer's purchase of the app, cryptographically signed by the App Store
// https://developer.apple.com/documentation/storekit/apptransaction
type AppTransaction struct {
	// ReceiptType is the server environment that signs the app transaction
	// https://developer.apple.com/documentation/storekit/apptransaction/3963901-environment
	ReceiptType *Environment `json:"receiptType,omitempty"`

	// AppAppleId is the unique identifier the App Store uses to identify the app
	// https://developer.apple.com/documentation/storekit/apptransaction/3954436-appid
	AppAppleId *int64 `json:"appAppleId,omitempty"`

	// BundleId is the bundle identifier that the app transaction applies to
	// https://developer.apple.com/documentation/storekit/apptransaction/3954439-bundleid
	BundleId *string `json:"bundleId,omitempty"`

	// ApplicationVersion is the app version that the app transaction applies to
	// https://developer.apple.com/documentation/storekit/apptransaction/3954437-appversion
	ApplicationVersion *string `json:"applicationVersion,omitempty"`

	// VersionExternalIdentifier is the version external identifier of the app
	// https://developer.apple.com/documentation/storekit/apptransaction/3954438-appversionid
	VersionExternalIdentifier *int64 `json:"versionExternalIdentifier,omitempty"`

	// ReceiptCreationDate is the date that the App Store signed the JWS app transaction
	// https://developer.apple.com/documentation/storekit/apptransaction/3954449-signeddate
	ReceiptCreationDate *int64 `json:"receiptCreationDate,omitempty"`

	// OriginalPurchaseDate is the date the user originally purchased the app from the App Store
	// https://developer.apple.com/documentation/storekit/apptransaction/3954448-originalpurchasedate
	OriginalPurchaseDate *int64 `json:"originalPurchaseDate,omitempty"`

	// OriginalApplicationVersion is the app version that the user originally purchased from the App Store
	// https://developer.apple.com/documentation/storekit/apptransaction/3954447-originalappversion
	OriginalApplicationVersion *string `json:"originalApplicationVersion,omitempty"`

	// DeviceVerification is the Base64 device verification value to use to verify whether the app transaction belongs to the device
	// https://developer.apple.com/documentation/storekit/apptransaction/3954441-deviceverification
	DeviceVerification *string `json:"deviceVerification,omitempty"`

	// DeviceVerificationNonce is the UUID used to compute the device verification value
	// https://developer.apple.com/documentation/storekit/apptransaction/3954442-deviceverificationnonce
	DeviceVerificationNonce *string `json:"deviceVerificationNonce,omitempty"`

	// PreorderDate is the date the customer placed an order for the app before it's available in the App Store
	// https://developer.apple.com/documentation/storekit/apptransaction/4013175-preorderdate
	PreorderDate *int64 `json:"preorderDate,omitempty"`
}
