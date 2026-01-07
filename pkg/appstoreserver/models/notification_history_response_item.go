// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// NotificationHistoryResponseItem represents the App Store server notification history record, including the signed notification payload and the result of the server's first send attempt
// https://developer.apple.com/documentation/appstoreserverapi/notificationhistoryresponseitem
type NotificationHistoryResponseItem struct {
	// SignedPayload is a cryptographically signed payload, in JSON Web Signature (JWS) format, containing the response body for a version 2 notification
	// https://developer.apple.com/documentation/appstoreservernotifications/signedpayload
	SignedPayload *string `json:"signedPayload,omitempty"`

	// SendAttempts is an array of information the App Store server records for its attempts to send a notification to your server
	// The maximum number of entries in the array is six
	// https://developer.apple.com/documentation/appstoreserverapi/sendattemptitem
	SendAttempts []SendAttemptItem `json:"sendAttempts,omitempty"`
}
