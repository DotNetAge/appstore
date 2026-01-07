// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// NotificationHistoryRequest represents the request body for notification history
// https://developer.apple.com/documentation/appstoreserverapi/notificationhistoryrequest
type NotificationHistoryRequest struct {
	// StartDate is the start date of the timespan for the requested App Store Server Notification history records
	// The startDate needs to precede the endDate
	// Choose a startDate that's within the past 180 days from the current date
	// https://developer.apple.com/documentation/appstoreserverapi/startdate
	StartDate *int64 `json:"startDate,omitempty"`

	// EndDate is the end date of the timespan for the requested App Store Server Notification history records
	// Choose an endDate that's later than the startDate
	// If you choose an endDate in the future, the endpoint automatically uses the current date as the endDate
	// https://developer.apple.com/documentation/appstoreserverapi/enddate
	EndDate *int64 `json:"endDate,omitempty"`

	// NotificationType is a notification type to limit the notification history records to those with this one notification type
	// Include either the transactionId or the notificationType in your query, but not both
	// https://developer.apple.com/documentation/appstoreserverapi/notificationtype
	NotificationType *NotificationTypeV2 `json:"notificationType,omitempty"`

	// NotificationSubtype is a notification subtype to limit the notification history records to those with this one notification subtype
	// If you specify a notificationSubtype, you need to also specify its related notificationType
	// https://developer.apple.com/documentation/appstoreserverapi/notificationsubtype
	NotificationSubtype *Subtype `json:"notificationSubtype,omitempty"`

	// TransactionId is the transaction identifier, which may be an original transaction identifier, of any transaction belonging to the customer
	// Provide this field to limit the notification history request to this one customer
	// Include either the transactionId or the notificationType in your query, but not both
	// https://developer.apple.com/documentation/appstoreserverapi/transactionid
	TransactionId *string `json:"transactionId,omitempty"`

	// OnlyFailures is a Boolean value you set to true to request only the notifications that haven't reached your server successfully
	// The response also includes notifications that the App Store server is currently retrying to send to your server
	// https://developer.apple.com/documentation/appstoreserverapi/onlyfailures
	OnlyFailures *bool `json:"onlyFailures,omitempty"`
}
