// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// SendTestNotificationResponse represents a response that contains the test notification token
// https://developer.apple.com/documentation/appstoreserverapi/sendtestnotificationresponse
type SendTestNotificationResponse struct {
	// TestNotificationToken is a unique identifier for a notification test that the App Store server sends to your server
	// https://developer.apple.com/documentation/appstoreserverapi/testnotificationtoken
	TestNotificationToken *string `json:"testNotificationToken,omitempty"`
}
