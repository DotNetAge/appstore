package pkg

import (
	"context"

	"github.com/DotNetAge/appstore/pkg/internal/models"
)

// ExtendRenewalDateForAllActiveSubscribers uses a subscription's product identifier to extend the renewal date for all of its eligible active subscribers.
// https://developer.apple.com/documentation/appstoreserverapi/extend_subscription_renewal_dates_for_all_active_subscribers
func (c *AppStoreServerClient) ExtendRenewalDateForAllActiveSubscribers(ctx context.Context, request *models.MassExtendRenewalDateRequest) (*string, error) {
	resp, err := c.client.ExtendRenewalDateForAllActiveSubscribers(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.RequestIdentifier, nil
}

// ExtendSubscriptionRenewalDate extends the renewal date of a customer's active subscription using the original transaction identifier.
// https://developer.apple.com/documentation/appstoreserverapi/extend_a_subscription_renewal_date
func (c *AppStoreServerClient) ExtendSubscriptionRenewalDate(ctx context.Context, originalTransactionID string, request *models.ExtendRenewalDateRequest) (*models.ExtendRenewalDateResponse, error) {
	return c.client.ExtendSubscriptionRenewalDate(ctx, originalTransactionID, request)
}

// GetAllSubscriptionStatuses gets the statuses for all of a customer's auto-renewable subscriptions in your app.
// https://developer.apple.com/documentation/appstoreserverapi/get_all_subscription_statuses
func (c *AppStoreServerClient) GetAllSubscriptionStatuses(ctx context.Context, transactionID string, status []models.Status) (*models.StatusResponse, error) {
	return c.client.GetAllSubscriptionStatuses(ctx, transactionID, status)
}

// GetRefundHistory gets a paginated list of all of a customer's refunded in-app purchases for your app.
// https://developer.apple.com/documentation/appstoreserverapi/get_refund_history
func (c *AppStoreServerClient) GetRefundHistory(ctx context.Context, transactionID string, revision string) (*models.RefundHistoryResponse, error) {
	return c.client.GetRefundHistory(ctx, transactionID, revision)
}

// GetStatusOfSubscriptionRenewalDateExtensions checks whether a renewal date extension request completed, and provides the final count of successful or failed extensions.
// https://developer.apple.com/documentation/appstoreserverapi/get_status_of_subscription_renewal_date_extensions
func (c *AppStoreServerClient) GetStatusOfSubscriptionRenewalDateExtensions(ctx context.Context, requestIdentifier string, productID string) (*models.MassExtendRenewalDateStatusResponse, error) {
	return c.client.GetStatusOfSubscriptionRenewalDateExtensions(ctx, requestIdentifier, productID)
}
