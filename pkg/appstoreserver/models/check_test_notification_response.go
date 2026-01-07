// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// CheckTestNotificationResponse represents a response that contains the contents of the test notification sent by the App Store server and the result from your server.
// https://developer.apple.com/documentation/appstoreserverapi/checktestnotificationresponse
type CheckTestNotificationResponse struct {
	// SignedPayload is a cryptographically signed payload, in JSON Web Signature (JWS) format, containing the response body for a version 2 notification
	SignedPayload *string `json:"signedPayload,omitempty"`
	// SendAttempts is an array of information the App Store server records for its attempts to send the TEST notification to your server
	SendAttempts []SendAttemptItem `json:"sendAttempts,omitempty"`
}
