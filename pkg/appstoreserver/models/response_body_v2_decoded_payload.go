// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// ResponseBodyV2DecodedPayload represents a decoded payload containing the version 2 notification data
// https://developer.apple.com/documentation/appstoreservernotifications/responsebodyv2decodedpayload
type ResponseBodyV2DecodedPayload struct {
	// NotificationType is the in-app purchase event for which the App Store sends this version 2 notification
	// https://developer.apple.com/documentation/appstoreservernotifications/notificationtype
	NotificationType *NotificationTypeV2 `json:"notificationType,omitempty"`

	// Subtype is additional information that identifies the notification event
	// The subtype field is present only for specific version 2 notifications
	// https://developer.apple.com/documentation/appstoreservernotifications/subtype
	Subtype *Subtype `json:"subtype,omitempty"`

	// NotificationUUID is a unique identifier for the notification
	// https://developer.apple.com/documentation/appstoreservernotifications/notificationuuid
	NotificationUUID *string `json:"notificationUUID,omitempty"`

	// Data is the object that contains the app metadata and signed renewal and transaction information
	// The data, summary, and externalPurchaseToken fields are mutually exclusive
	// The payload contains only one of these fields
	// https://developer.apple.com/documentation/appstoreservernotifications/data
	Data *Data `json:"data,omitempty"`

	// Version is a string that indicates the notification's App Store Server Notifications version number
	// https://developer.apple.com/documentation/appstoreservernotifications/version
	Version *string `json:"version,omitempty"`

	// SignedDate is the UNIX time, in milliseconds, that the App Store signed the JSON Web Signature data
	// https://developer.apple.com/documentation/appstoreserverapi/signeddate
	SignedDate *int64 `json:"signedDate,omitempty"`

	// Summary is the summary data that appears when the App Store server completes your request to extend a subscription renewal date for eligible subscribers
	// The data, summary, and externalPurchaseToken fields are mutually exclusive
	// The payload contains only one of these fields
	// https://developer.apple.com/documentation/appstoreservernotifications/summary
	Summary *Summary `json:"summary,omitempty"`

	// ExternalPurchaseToken is the field that appears when the notificationType is EXTERNAL_PURCHASE_TOKEN
	// The data, summary, and externalPurchaseToken fields are mutually exclusive
	// The payload contains only one of these fields
	// https://developer.apple.com/documentation/appstoreservernotifications/externalpurchasetoken
	ExternalPurchaseToken *ExternalPurchaseToken `json:"externalPurchaseToken,omitempty"`
}
