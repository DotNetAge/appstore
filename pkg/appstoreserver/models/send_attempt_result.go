// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// SendAttemptResult represents the success or error information the App Store server records when it attempts to send an App Store server notification to your server.
// https://developer.apple.com/documentation/appstoreserverapi/sendattemptresult
type SendAttemptResult string

const (
	// SendAttemptResultSuccess indicates the notification was successfully sent
	SendAttemptResultSuccess SendAttemptResult = "SUCCESS"
	// SendAttemptResultTimedOut indicates the notification attempt timed out
	SendAttemptResultTimedOut SendAttemptResult = "TIMED_OUT"
	// SendAttemptResultTLSIssue indicates there was a TLS issue with the notification attempt
	SendAttemptResultTLSIssue SendAttemptResult = "TLS_ISSUE"
	// SendAttemptResultCircularRedirect indicates there was a circular redirect with the notification attempt
	SendAttemptResultCircularRedirect SendAttemptResult = "CIRCULAR_REDIRECT"
	// SendAttemptResultNoResponse indicates there was no response to the notification attempt
	SendAttemptResultNoResponse SendAttemptResult = "NO_RESPONSE"
	// SendAttemptResultSocketIssue indicates there was a socket issue with the notification attempt
	SendAttemptResultSocketIssue SendAttemptResult = "SOCKET_ISSUE"
	// SendAttemptResultUnsupportedCharset indicates there was an unsupported charset with the notification attempt
	SendAttemptResultUnsupportedCharset SendAttemptResult = "UNSUPPORTED_CHARSET"
	// SendAttemptResultInvalidResponse indicates there was an invalid response to the notification attempt
	SendAttemptResultInvalidResponse SendAttemptResult = "INVALID_RESPONSE"
	// SendAttemptResultPrematureClose indicates the connection was prematurely closed during the notification attempt
	SendAttemptResultPrematureClose SendAttemptResult = "PREMATURE_CLOSE"
	// SendAttemptResultUnsuccessfulHTTPResponseCode indicates there was an unsuccessful HTTP response code to the notification attempt
	SendAttemptResultUnsuccessfulHTTPResponseCode SendAttemptResult = "UNSUCCESSFUL_HTTP_RESPONSE_CODE"
	// SendAttemptResultOther indicates there was another issue with the notification attempt
	SendAttemptResultOther SendAttemptResult = "OTHER"
)
