// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// ProductType represents the type of product for transaction history request
// https://developer.apple.com/documentation/appstoreserverapi/producttype
type ProductType string

const (
	// ProductTypeAutoRenewable indicates an auto-renewable subscription
	ProductTypeAutoRenewable ProductType = "AUTO_RENEWABLE"
	// ProductTypeNonRenewable indicates a non-renewable subscription
	ProductTypeNonRenewable ProductType = "NON_RENEWABLE"
	// ProductTypeConsumable indicates a consumable in-app purchase
	ProductTypeConsumable ProductType = "CONSUMABLE"
	// ProductTypeNonConsumable indicates a non-consumable in-app purchase
	ProductTypeNonConsumable ProductType = "NON_CONSUMABLE"
)

// Order represents the sort order for transaction history request
// https://developer.apple.com/documentation/appstoreserverapi/order
type Order string

const (
	// OrderAscending indicates ascending order
	OrderAscending Order = "ASCENDING"
	// OrderDescending indicates descending order
	OrderDescending Order = "DESCENDING"
)

// TransactionHistoryRequest represents a request for transaction history
// https://developer.apple.com/documentation/appstoreserverapi/transactionhistoryrequest
type TransactionHistoryRequest struct {
	// StartDate is an optional start date of the timespan for the transaction history records
	// The startDate must precede the endDate if you specify both dates
	// To be included in results, the transaction's purchaseDate must be equal to or greater than the startDate
	// https://developer.apple.com/documentation/appstoreserverapi/startdate
	StartDate *int64 `json:"startDate,omitempty"`

	// EndDate is an optional end date of the timespan for the transaction history records
	// Choose an endDate that's later than the startDate if you specify both dates
	// Using an endDate in the future is valid
	// To be included in results, the transaction's purchaseDate must be less than the endDate
	// https://developer.apple.com/documentation/appstoreserverapi/enddate
	EndDate *int64 `json:"endDate,omitempty"`

	// ProductIds is an optional filter that indicates the product identifiers to include in the transaction history
	// https://developer.apple.com/documentation/appstoreserverapi/productid
	ProductIds []string `json:"productIds,omitempty"`

	// ProductTypes is an optional filter that indicates the product types to include in the transaction history
	ProductTypes []ProductType `json:"productTypes,omitempty"`

	// Sort is an optional sort order for the transaction history records
	// The response sorts the transaction records by their recently modified date
	// The default value is ASCENDING, so you receive the oldest records first
	Sort *Order `json:"sort,omitempty"`

	// SubscriptionGroupIdentifiers is an optional filter that indicates the subscription group identifiers to include in the transaction history
	// https://developer.apple.com/documentation/appstoreserverapi/subscriptiongroupidentifier
	SubscriptionGroupIdentifiers []string `json:"subscriptionGroupIdentifiers,omitempty"`

	// InAppOwnershipType is an optional filter that limits the transaction history by the in-app ownership type
	// https://developer.apple.com/documentation/appstoreserverapi/inappownershiptype
	InAppOwnershipType *InAppOwnershipType `json:"inAppOwnershipType,omitempty"`

	// Revoked is an optional Boolean value that indicates whether the response includes only revoked transactions when true, or contains only nonrevoked transactions when false
	// By default, the request doesn't include this parameter
	Revoked *bool `json:"revoked,omitempty"`
}
