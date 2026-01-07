package internal

import (
	"context"
	"net/url"

	"github.com/DotNetAge/appstore/internal/models"
)

// GetTestNotificationStatus checks the status of the test App Store server notification sent to your server.
// https://developer.apple.com/documentation/appstoreserverapi/get_test_notification_status
func (c *AppStoreServerAPIClient) GetTestNotificationStatus(ctx context.Context, testNotificationToken string) (*models.CheckTestNotificationResponse, error) {
	var response models.CheckTestNotificationResponse
	path := "/inApps/v1/notifications/test/" + testNotificationToken
	if err := c.makeRequest(ctx, path, "GET", url.Values{}, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetNotificationHistory gets a list of notifications that the App Store server attempted to send to your server.
// https://developer.apple.com/documentation/appstoreserverapi/get_notification_history
func (c *AppStoreServerAPIClient) GetNotificationHistory(ctx context.Context, paginationToken string, request *models.NotificationHistoryRequest) (*models.NotificationHistoryResponse, error) {
	var response models.NotificationHistoryResponse
	path := "/inApps/v1/notifications/history"

	queryParams := url.Values{}
	if paginationToken != "" {
		queryParams.Add("paginationToken", paginationToken)
	}

	if err := c.makeRequest(ctx, path, "POST", queryParams, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// RequestTestNotification asks App Store Server Notifications to send a test notification to your server.
// https://developer.apple.com/documentation/appstoreserverapi/request_a_test_notification
func (c *AppStoreServerAPIClient) RequestTestNotification(ctx context.Context) (*models.SendTestNotificationResponse, error) {
	var response models.SendTestNotificationResponse
	path := "/inApps/v1/notifications/test"
	if err := c.makeRequest(ctx, path, "POST", url.Values{}, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
