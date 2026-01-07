// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// SendAttemptItem represents the success or error information and the date the App Store server records when it attempts to send a server notification to your server.
// https://developer.apple.com/documentation/appstoreserverapi/sendattemptitem
type SendAttemptItem struct {
	// AttemptDate is the date the App Store server attempts to send a notification
	AttemptDate *int64 `json:"attemptDate,omitempty"`
	// SendAttemptResult is the success or error information the App Store server records when it attempts to send an App Store server notification to your server
	SendAttemptResult *SendAttemptResult `json:"sendAttemptResult,omitempty"`
}
