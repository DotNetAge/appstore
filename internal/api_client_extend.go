package internal

import (
	"context"
	"fmt"
	"net/url"

	"github.com/DotNetAge/appstore/pkg/internal/models"
)

// ExtendRenewalDateForAllActiveSubscribers uses a subscription's product identifier to extend the renewal date for all of its eligible active subscribers.
// https://developer.apple.com/documentation/appstoreserverapi/extend_subscription_renewal_dates_for_all_active_subscribers
func (c *AppStoreServerAPIClient) ExtendRenewalDateForAllActiveSubscribers(ctx context.Context, request *models.MassExtendRenewalDateRequest) (*models.MassExtendRenewalDateResponse, error) {
	var response models.MassExtendRenewalDateResponse
	if err := c.makeRequest(ctx, "/inApps/v1/subscriptions/extend/mass", "POST", url.Values{}, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ExtendSubscriptionRenewalDate extends the renewal date of a customer's active subscription using the original transaction identifier.
// https://developer.apple.com/documentation/appstoreserverapi/extend_a_subscription_renewal_date
func (c *AppStoreServerAPIClient) ExtendSubscriptionRenewalDate(ctx context.Context, originalTransactionID string, request *models.ExtendRenewalDateRequest) (*models.ExtendRenewalDateResponse, error) {
	var response models.ExtendRenewalDateResponse
	path := "/inApps/v1/subscriptions/extend/" + originalTransactionID
	if err := c.makeRequest(ctx, path, "PUT", url.Values{}, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetAllSubscriptionStatuses gets the statuses for all of a customer's auto-renewable subscriptions in your app.
// https://developer.apple.com/documentation/appstoreserverapi/get_all_subscription_statuses
func (c *AppStoreServerAPIClient) GetAllSubscriptionStatuses(ctx context.Context, transactionID string, status []models.Status) (*models.StatusResponse, error) {
	var response models.StatusResponse
	path := "/inApps/v1/subscriptions/" + transactionID

	queryParams := url.Values{}
	if len(status) > 0 {
		for _, s := range status {
			queryParams.Add("status", fmt.Sprintf("%d", s))
		}
	}

	if err := c.makeRequest(ctx, path, "GET", queryParams, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetRefundHistory gets a paginated list of all of a customer's refunded in-app purchases for your app.
// https://developer.apple.com/documentation/appstoreserverapi/get_refund_history
func (c *AppStoreServerAPIClient) GetRefundHistory(ctx context.Context, transactionID string, revision string) (*models.RefundHistoryResponse, error) {
	var response models.RefundHistoryResponse
	path := "/inApps/v2/refund/lookup/" + transactionID

	queryParams := url.Values{}
	if revision != "" {
		queryParams.Add("revision", revision)
	}

	if err := c.makeRequest(ctx, path, "GET", queryParams, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetStatusOfSubscriptionRenewalDateExtensions checks whether a renewal date extension request completed, and provides the final count of successful or failed extensions.
// https://developer.apple.com/documentation/appstoreserverapi/get_status_of_subscription_renewal_date_extensions
func (c *AppStoreServerAPIClient) GetStatusOfSubscriptionRenewalDateExtensions(ctx context.Context, requestIdentifier string, productID string) (*models.MassExtendRenewalDateStatusResponse, error) {
	var response models.MassExtendRenewalDateStatusResponse
	path := "/inApps/v1/subscriptions/extend/mass/" + productID + "/" + requestIdentifier

	if err := c.makeRequest(ctx, path, "GET", url.Values{}, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
