// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// NotificationHistoryResponse represents a response that contains the App Store Server Notifications history for your app
// https://developer.apple.com/documentation/appstoreserverapi/notificationhistoryresponse
type NotificationHistoryResponse struct {
	// PaginationToken is a pagination token that you return to the endpoint on a subsequent call to receive the next set of results
	// https://developer.apple.com/documentation/appstoreserverapi/paginationtoken
	PaginationToken *string `json:"paginationToken,omitempty"`

	// HasMore is a Boolean value indicating whether the App Store has more transaction data
	// https://developer.apple.com/documentation/appstoreserverapi/hasmore
	HasMore *bool `json:"hasMore,omitempty"`

	// NotificationHistory is an array of App Store server notification history records
	// https://developer.apple.com/documentation/appstoreserverapi/notificationhistoryresponseitem
	NotificationHistory []NotificationHistoryResponseItem `json:"notificationHistory,omitempty"`
}
