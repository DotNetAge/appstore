package pkg

import (
	"context"
	"fmt"

	"github.com/DotNetAge/appstore/internal/models"
)

// GetTestNotificationStatus checks the status of the test App Store server notification sent to your server.
// https://developer.apple.com/documentation/appstoreserverapi/get_test_notification_status
func (c *AppStoreServerClient) GetTestNotificationStatus(ctx context.Context, testNotificationToken string) (*models.ResponseBodyV2DecodedPayload, error) {
	resp, err := c.client.GetTestNotificationStatus(ctx, testNotificationToken)
	if err != nil {
		return nil, err
	}
	if resp.SignedPayload == nil {
		return nil, fmt.Errorf("signed payload is nil")
	}

	decodedPayload, err := c.verifier.VerifyAndDecodeNotification(*resp.SignedPayload)
	if err != nil {
		return nil, err
	}
	if decodedPayload == nil {
		return nil, fmt.Errorf("decoded payload is nil")
	}
	return decodedPayload, nil
}

// NotificationHistoryOption is a function type for configuring notification history requests
type NotificationHistoryOption func(*notificationHistoryOptions)

// notificationHistoryOptions holds the options for notification history requests
type notificationHistoryOptions struct {
	startDate           *int64
	endDate             *int64
	notificationType    *models.NotificationTypeV2
	notificationSubtype *models.Subtype
	transactionId       *string
	onlyFailures        *bool
}

// WithNotificationStartDate sets the start date for the notification history request
func WithNotificationStartDate(startDate int64) NotificationHistoryOption {
	return func(opts *notificationHistoryOptions) {
		opts.startDate = &startDate
	}
}

// WithNotificationEndDate sets the end date for the notification history request
func WithNotificationEndDate(endDate int64) NotificationHistoryOption {
	return func(opts *notificationHistoryOptions) {
		opts.endDate = &endDate
	}
}

// WithNotificationType sets the notification type for the notification history request
func WithNotificationType(notificationType models.NotificationTypeV2) NotificationHistoryOption {
	return func(opts *notificationHistoryOptions) {
		opts.notificationType = &notificationType
	}
}

// WithNotificationSubtype sets the notification subtype for the notification history request
func WithNotificationSubtype(notificationSubtype models.Subtype) NotificationHistoryOption {
	return func(opts *notificationHistoryOptions) {
		opts.notificationSubtype = &notificationSubtype
	}
}

// WithTransactionId sets the transaction ID for the notification history request
func WithTransactionId(transactionId string) NotificationHistoryOption {
	return func(opts *notificationHistoryOptions) {
		opts.transactionId = &transactionId
	}
}

// WithOnlyFailures sets whether to only return failed notifications
func WithOnlyFailures(onlyFailures bool) NotificationHistoryOption {
	return func(opts *notificationHistoryOptions) {
		opts.onlyFailures = &onlyFailures
	}
}

// decodeNotifications decodes a slice of notification history items into decoded payloads
func (c *AppStoreServerClient) decodeNotifications(notifications []models.NotificationHistoryResponseItem) ([]*models.ResponseBodyV2DecodedPayload, error) {
	var decodedPayloads []*models.ResponseBodyV2DecodedPayload
	for _, notification := range notifications {
		if notification.SignedPayload == nil {
			return nil, fmt.Errorf("signed payload is nil")
		}

		decodedPayload, err := c.verifier.VerifyAndDecodeNotification(*notification.SignedPayload)
		if err != nil {
			return nil, err
		}
		if decodedPayload == nil {
			return nil, fmt.Errorf("decoded payload is nil")
		}
		decodedPayloads = append(decodedPayloads, decodedPayload)
	}
	return decodedPayloads, nil
}

// GetNotificationHistory gets a list of notifications that the App Store server attempted to send to your server.
// https://developer.apple.com/documentation/appstoreserverapi/get_notification_history
func (c *AppStoreServerClient) GetNotificationHistory(ctx context.Context, paginationToken string, options ...NotificationHistoryOption) ([]*models.ResponseBodyV2DecodedPayload, error) {
	// Apply options
	opts := &notificationHistoryOptions{}
	for _, option := range options {
		option(opts)
	}

	// Build request
	request := &models.NotificationHistoryRequest{
		StartDate:           opts.startDate,
		EndDate:             opts.endDate,
		NotificationType:    opts.notificationType,
		NotificationSubtype: opts.notificationSubtype,
		TransactionId:       opts.transactionId,
		OnlyFailures:        opts.onlyFailures,
	}

	resp, err := c.client.GetNotificationHistory(ctx, paginationToken, request)
	if err != nil {
		return nil, err
	}

	if resp.NotificationHistory == nil {
		return nil, fmt.Errorf("notification history is nil")
	}

	decodedPayloads, err := c.decodeNotifications(resp.NotificationHistory)
	if err != nil {
		return nil, err
	}

	for resp.HasMore != nil && *resp.HasMore {
		if resp.PaginationToken == nil {
			break
		}

		nextResp, err := c.client.GetNotificationHistory(ctx, *resp.PaginationToken, request)
		if err != nil {
			return nil, err
		}

		if nextResp.NotificationHistory == nil {
			break
		}

		nextDecoded, err := c.decodeNotifications(nextResp.NotificationHistory)
		if err != nil {
			return nil, err
		}
		decodedPayloads = append(decodedPayloads, nextDecoded...)

		resp = nextResp
	}

	return decodedPayloads, nil
}

func (c *AppStoreServerClient) TestNotification() (*models.ResponseBodyV2DecodedPayload, error) {
	// 1. Request Test Notification
	resp, err := c.client.RequestTestNotification(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to request test notification: %w", err)
	}

	token := resp.TestNotificationToken
	if token == nil {
		return nil, fmt.Errorf("test notification token is nil")
	}

	// 2. Request Test Notification Status with token
	notifyResp, err := c.client.GetTestNotificationStatus(context.Background(), *token)
	if err != nil {
		return nil, fmt.Errorf("failed to get test notification status: %w", err)
	}

	// 3. Decode SignedPayload
	if notifyResp.SignedPayload == nil {
		return nil, fmt.Errorf("signed payload is nil")
	}

	// 验证并解码通知
	responseBodyV2, err := c.verifier.VerifyAndDecodeNotification(*notifyResp.SignedPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to verify and decode notification: %w", err)
	}

	return responseBodyV2, nil
}
