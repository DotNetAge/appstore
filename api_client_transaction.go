package appstore

import (
	"context"
	"fmt"
	"net/url"

	"github.com/DotNetAge/appstore/models"
)

// GetTransactionHistory gets a customer's in-app purchase transaction history for your app.
// https://developer.apple.com/documentation/appstoreserverapi/get_transaction_history
func (c *AppStoreServerAPIClient) GetTransactionHistory(ctx context.Context, transactionID string, revision string, request *models.TransactionHistoryRequest, version GetTransactionHistoryVersion) (*models.HistoryResponse, error) {
	var response models.HistoryResponse
	path := "/inApps/" + string(version) + "/history/" + transactionID

	queryParams := url.Values{}
	if revision != "" {
		queryParams.Add("revision", revision)
	}

	if request != nil {
		if request.StartDate != nil {
			queryParams.Add("startDate", fmt.Sprintf("%d", *request.StartDate))
		}
		if request.EndDate != nil {
			queryParams.Add("endDate", fmt.Sprintf("%d", *request.EndDate))
		}
		if request.ProductIds != nil {
			for _, id := range request.ProductIds {
				queryParams.Add("productId", id)
			}
		}
		if request.ProductTypes != nil {
			for _, t := range request.ProductTypes {
				queryParams.Add("productType", string(t))
			}
		}
		if request.Sort != nil {
			queryParams.Add("sort", string(*request.Sort))
		}
		if request.SubscriptionGroupIdentifiers != nil {
			for _, id := range request.SubscriptionGroupIdentifiers {
				queryParams.Add("subscriptionGroupIdentifier", id)
			}
		}
		if request.InAppOwnershipType != nil {
			queryParams.Add("inAppOwnershipType", string(*request.InAppOwnershipType))
		}
		if request.Revoked != nil {
			queryParams.Add("revoked", fmt.Sprintf("%t", *request.Revoked))
		}
	}

	if err := c.makeRequest(ctx, path, "GET", queryParams, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetTransactionInfo gets information about a single transaction for your app.
// https://developer.apple.com/documentation/appstoreserverapi/get_transaction_info
func (c *AppStoreServerAPIClient) GetTransactionInfo(ctx context.Context, transactionID string) (*models.TransactionInfoResponse, error) {
	var response models.TransactionInfoResponse
	path := "/inApps/v1/transactions/" + transactionID
	if err := c.makeRequest(ctx, path, "GET", url.Values{}, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// LookUpOrderID gets a customer's in-app purchases from a receipt using the order ID.
// https://developer.apple.com/documentation/appstoreserverapi/look_up_order_id
func (c *AppStoreServerAPIClient) LookUpOrderID(ctx context.Context, orderID string) (*models.OrderLookupResponse, error) {
	var response models.OrderLookupResponse
	path := "/inApps/v1/lookup/" + orderID
	if err := c.makeRequest(ctx, path, "GET", url.Values{}, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// SendConsumptionData sends consumption information about a consumable in-app purchase to the App Store.
// https://developer.apple.com/documentation/appstoreserverapi/send_consumption_information
func (c *AppStoreServerAPIClient) SendConsumptionData(ctx context.Context, transactionID string, request *models.ConsumptionRequest) error {
	path := "/inApps/v1/transactions/consumption/" + transactionID
	return c.makeRequest(ctx, path, "PUT", url.Values{}, request, nil)
}
