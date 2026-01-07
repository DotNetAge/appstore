package pkg

import (
	"context"
	"fmt"
	"time"

	"github.com/DotNetAge/appstore/internal"
	"github.com/DotNetAge/appstore/internal/models"
)

// TransactionHistoryOption defines the option type for GetTransactionHistory
type TransactionHistoryOption func(*transactionHistoryOptions)

// transactionHistoryOptions holds the options for GetTransactionHistory
type transactionHistoryOptions struct {
	revision                     string
	version                      internal.GetTransactionHistoryVersion
	startDate                    int64
	endDate                      int64
	productIds                   []string
	productTypes                 []models.ProductType
	sort                         models.Order
	subscriptionGroupIdentifiers []string
	inAppOwnershipType           models.InAppOwnershipType
	revoked                      bool
}

// WithRevision sets the revision parameter
func WithRevision(revision string) TransactionHistoryOption {
	return func(opts *transactionHistoryOptions) {
		opts.revision = revision
	}
}

func WithStartDate(startDate time.Time) TransactionHistoryOption {
	return func(opts *transactionHistoryOptions) {
		opts.startDate = startDate.Unix()
	}
}

func WithSubscriptionGroupIdentifiers(subscriptionGroupIdentifiers ...string) TransactionHistoryOption {
	return func(opts *transactionHistoryOptions) {
		opts.subscriptionGroupIdentifiers = subscriptionGroupIdentifiers
	}
}

func WithInAppOwnershipType(inAppOwnershipType models.InAppOwnershipType) TransactionHistoryOption {
	return func(opts *transactionHistoryOptions) {
		opts.inAppOwnershipType = inAppOwnershipType
	}
}

func WithRevoked(revoked bool) TransactionHistoryOption {
	return func(opts *transactionHistoryOptions) {
		opts.revoked = revoked
	}
}
func WithProductIds(productIds ...string) TransactionHistoryOption {
	return func(opts *transactionHistoryOptions) {
		opts.productIds = productIds
	}
}

func WithSort(sort models.Order) TransactionHistoryOption {
	return func(opts *transactionHistoryOptions) {
		opts.sort = sort
	}
}

func WithProductTypes(productTypes ...models.ProductType) TransactionHistoryOption {
	return func(opts *transactionHistoryOptions) {
		opts.productTypes = productTypes
	}
}

func WithEndDate(endDate time.Time) TransactionHistoryOption {
	return func(opts *transactionHistoryOptions) {
		opts.endDate = endDate.Unix()
	}
}

// WithVersion sets the version parameter
func WithVersion(version internal.GetTransactionHistoryVersion) TransactionHistoryOption {
	return func(opts *transactionHistoryOptions) {
		opts.version = version
	}
}

// GetTransactionHistory gets a customer's in-app purchase transaction history for your app.
func (c *AppStoreServerClient) GetTransactionHistory(transactionID string, options ...TransactionHistoryOption) ([]*models.JWSTransactionDecodedPayload, error) {
	// Set default options
	opts := &transactionHistoryOptions{
		version: internal.GetTransactionHistoryVersionV1,
	}

	// Apply options
	for _, option := range options {
		option(opts)
	}

	request := &models.TransactionHistoryRequest{
		StartDate:                    &opts.startDate,
		EndDate:                      &opts.endDate,
		ProductIds:                   opts.productIds,
		ProductTypes:                 opts.productTypes,
		Sort:                         &opts.sort,
		SubscriptionGroupIdentifiers: opts.subscriptionGroupIdentifiers,
		InAppOwnershipType:           &opts.inAppOwnershipType,
		Revoked:                      &opts.revoked,
	}

	resp, err := c.client.GetTransactionHistory(context.Background(), transactionID, opts.revision, request, opts.version)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction history: %w", err)
	}

	// 验证并解码所有签名的交易
	var decodedTransactions []*models.JWSTransactionDecodedPayload
	for _, signedTransaction := range resp.SignedTransactions {
		playload, err := c.verifier.VerifyAndDecodeSignedTransaction(signedTransaction)
		if err != nil {
			return nil, fmt.Errorf("failed to verify and decode transaction: %w", err)
		}
		decodedTransactions = append(decodedTransactions, playload)
	}

	return decodedTransactions, nil
}

// GetTransactionInfo gets information about a single transaction for your app.
func (c *AppStoreServerClient) GetTransactionInfo(transactionID string) (*models.JWSTransactionDecodedPayload, error) {
	resp, err := c.client.GetTransactionInfo(context.Background(), transactionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction info: %w", err)
	}

	if resp.SignedTransactionInfo == nil {
		return nil, fmt.Errorf("signed transaction info is nil")
	}

	// 验证并解码签名的交易
	playload, err := c.verifier.VerifyAndDecodeSignedTransaction(*resp.SignedTransactionInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to verify and decode transaction: %w", err)
	}

	return playload, nil
}

// LookUpOrderID gets a customer's in-app purchases from a receipt using the order ID.
func (c *AppStoreServerClient) LookUpOrderID(orderID string) ([]*models.JWSTransactionDecodedPayload, error) {
	resp, err := c.client.LookUpOrderID(context.Background(), orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to look up order ID: %w", err)
	}

	var decodedTransactions []*models.JWSTransactionDecodedPayload

	// 验证并解码所有签名的交易
	for _, signedTransaction := range resp.SignedTransactions {
		playload, err := c.verifier.VerifyAndDecodeSignedTransaction(signedTransaction)
		if err != nil {
			return nil, fmt.Errorf("failed to verify and decode transaction: %w", err)
		}
		decodedTransactions = append(decodedTransactions, playload)
	}

	return decodedTransactions, nil
}

// SendConsumptionData sends consumption information about a consumable in-app purchase to the App Store.
func (c *AppStoreServerClient) SendConsumptionData(transactionID string, request *models.ConsumptionRequest) error {
	err := c.client.SendConsumptionData(context.Background(), transactionID, request)
	if err != nil {
		return fmt.Errorf("failed to send consumption data: %w", err)
	}

	// 构建统一的响应格式
	return nil
}

func (c *AppStoreServerClient) GetTransaction(response *ResponseBodyV2) (*models.JWSTransactionDecodedPayload, error) {
	playload, err := c.verifier.VerifyAndDecodeSignedTransaction(response.SignedPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to verify and decode transaction: %w", err)
	}
	return playload, nil
}
