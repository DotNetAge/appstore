// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// APIError represents the error codes returned by the App Store Server API
// https://developer.apple.com/documentation/appstoreserverapi/error_codes
type APIError int

const (
	// APIErrorGeneralBadRequest indicates an invalid request
	APIErrorGeneralBadRequest APIError = 4000000
	// APIErrorInvalidAppIdentifier indicates an invalid app identifier
	APIErrorInvalidAppIdentifier APIError = 4000002
	// APIErrorInvalidRequestRevision indicates an invalid request revision
	APIErrorInvalidRequestRevision APIError = 4000005
	// APIErrorInvalidTransactionID indicates an invalid transaction identifier
	APIErrorInvalidTransactionID APIError = 4000006
	// APIErrorInvalidOriginalTransactionID indicates an invalid original transaction identifier
	APIErrorInvalidOriginalTransactionID APIError = 4000008
	// APIErrorInvalidExtendByDays indicates an invalid extend-by-days value
	APIErrorInvalidExtendByDays APIError = 4000009
	// APIErrorInvalidExtendReasonCode indicates an invalid reason code
	APIErrorInvalidExtendReasonCode APIError = 4000010
	// APIErrorInvalidRequestIdentifier indicates an invalid request identifier
	APIErrorInvalidRequestIdentifier APIError = 4000011
	// APIErrorStartDateTooFarInPast indicates that the start date is earlier than the earliest allowed date
	APIErrorStartDateTooFarInPast APIError = 4000012
	// APIErrorStartDateAfterEndDate indicates that the end date precedes the start date, or the two dates are equal
	APIErrorStartDateAfterEndDate APIError = 4000013
	// APIErrorInvalidPaginationToken indicates the pagination token is invalid
	APIErrorInvalidPaginationToken APIError = 4000014
	// APIErrorInvalidStartDate indicates the start date is invalid
	APIErrorInvalidStartDate APIError = 4000015
	// APIErrorInvalidEndDate indicates the end date is invalid
	APIErrorInvalidEndDate APIError = 4000016
	// APIErrorPaginationTokenExpired indicates the pagination token expired
	APIErrorPaginationTokenExpired APIError = 4000017
	// APIErrorInvalidNotificationType indicates the notification type or subtype is invalid
	APIErrorInvalidNotificationType APIError = 4000018
	// APIErrorMultipleFiltersSupplied indicates the request is invalid because it has too many constraints applied
	APIErrorMultipleFiltersSupplied APIError = 4000019
	// APIErrorInvalidTestNotificationToken indicates the test notification token is invalid
	APIErrorInvalidTestNotificationToken APIError = 4000020
	// APIErrorInvalidSort indicates an invalid sort parameter
	APIErrorInvalidSort APIError = 4000021
	// APIErrorInvalidProductType indicates an invalid product type parameter
	APIErrorInvalidProductType APIError = 4000022
	// APIErrorInvalidProductID indicates the product ID parameter is invalid
	APIErrorInvalidProductID APIError = 4000023
	// APIErrorInvalidSubscriptionGroupIdentifier indicates an invalid subscription group identifier
	APIErrorInvalidSubscriptionGroupIdentifier APIError = 4000024
	// APIErrorInvalidExcludeRevoked indicates the query parameter exclude-revoked is invalid
	APIErrorInvalidExcludeRevoked APIError = 4000025
	// APIErrorInvalidInAppOwnershipType indicates an invalid in-app ownership type parameter
	APIErrorInvalidInAppOwnershipType APIError = 4000026
	// APIErrorInvalidEmptyStorefrontCountryCodeList indicates a required storefront country code is empty
	APIErrorInvalidEmptyStorefrontCountryCodeList APIError = 4000027
	// APIErrorInvalidStorefrontCountryCode indicates a storefront code is invalid
	APIErrorInvalidStorefrontCountryCode APIError = 4000028
	// APIErrorInvalidRevoked indicates the revoked parameter contains an invalid value
	APIErrorInvalidRevoked APIError = 4000030
	// APIErrorInvalidStatus indicates the status parameter is invalid
	APIErrorInvalidStatus APIError = 4000031
	// APIErrorInvalidAccountTenure indicates the value of the account tenure field is invalid
	APIErrorInvalidAccountTenure APIError = 4000032
	// APIErrorInvalidAppAccountToken indicates the value of the app account token field is invalid
	APIErrorInvalidAppAccountToken APIError = 4000033
	// APIErrorInvalidConsumptionStatus indicates the value of the consumption status field is invalid
	APIErrorInvalidConsumptionStatus APIError = 4000034
	// APIErrorInvalidCustomerConsented indicates the customer consented field is invalid or doesn't indicate that the customer consented
	APIErrorInvalidCustomerConsented APIError = 4000035
	// APIErrorInvalidDeliveryStatus indicates the value in the delivery status field is invalid
	APIErrorInvalidDeliveryStatus APIError = 4000036
	// APIErrorInvalidLifetimeDollarsPurchased indicates the value in the lifetime dollars purchased field is invalid
	APIErrorInvalidLifetimeDollarsPurchased APIError = 4000037
	// APIErrorInvalidLifetimeDollarsRefunded indicates the value in the lifetime dollars refunded field is invalid
	APIErrorInvalidLifetimeDollarsRefunded APIError = 4000038
	// APIErrorInvalidPlatform indicates the value in the platform field is invalid
	APIErrorInvalidPlatform APIError = 4000039
	// APIErrorInvalidPlayTime indicates the value in the playtime field is invalid
	APIErrorInvalidPlayTime APIError = 4000040
	// APIErrorInvalidSampleContentProvided indicates the value in the sample content provided field is invalid
	APIErrorInvalidSampleContentProvided APIError = 4000041
	// APIErrorInvalidUserStatus indicates the value in the user status field is invalid
	APIErrorInvalidUserStatus APIError = 4000042
	// APIErrorInvalidTransactionNotConsumable indicates the transaction identifier doesn't represent a consumable in-app purchase
	APIErrorInvalidTransactionNotConsumable APIError = 4000043
	// APIErrorInvalidTransactionTypeNotSupported indicates the transaction identifier represents an unsupported in-app purchase type
	APIErrorInvalidTransactionTypeNotSupported APIError = 4000047
	// APIErrorSubscriptionExtensionIneligible indicates the subscription doesn't qualify for a renewal-date extension due to its subscription state
	APIErrorSubscriptionExtensionIneligible APIError = 4030004
	// APIErrorSubscriptionMaxExtension indicates the subscription doesn't qualify for a renewal-date extension because it has already received the maximum extensions
	APIErrorSubscriptionMaxExtension APIError = 4030005
	// APIErrorFamilySharedSubscriptionExtensionIneligible indicates a subscription isn't directly eligible for a renewal date extension because the user obtained it through Family Sharing
	APIErrorFamilySharedSubscriptionExtensionIneligible APIError = 4030007
	// APIErrorAccountNotFound indicates the App Store account wasn't found
	APIErrorAccountNotFound APIError = 4040001
	// APIErrorAccountNotFoundRetryable indicates the App Store account wasn't found, but you can try again
	APIErrorAccountNotFoundRetryable APIError = 4040002
	// APIErrorAppNotFound indicates the app wasn't found
	APIErrorAppNotFound APIError = 4040003
	// APIErrorAppNotFoundRetryable indicates the app wasn't found, but you can try again
	APIErrorAppNotFoundRetryable APIError = 4040004
	// APIErrorOriginalTransactionIDNotFound indicates an original transaction identifier wasn't found
	APIErrorOriginalTransactionIDNotFound APIError = 4040005
	// APIErrorOriginalTransactionIDNotFoundRetryable indicates the original transaction identifier wasn't found, but you can try again
	APIErrorOriginalTransactionIDNotFoundRetryable APIError = 4040006
	// APIErrorServerNotificationURLNotFound indicates that the App Store server couldn't find a notifications URL for your app in this environment
	APIErrorServerNotificationURLNotFound APIError = 4040007
	// APIErrorTestNotificationNotFound indicates that the test notification token is expired or the test notification status isn't available
	APIErrorTestNotificationNotFound APIError = 4040008
	// APIErrorStatusRequestNotFound indicates the server didn't find a subscription-renewal-date extension request for the request identifier and product identifier you provided
	APIErrorStatusRequestNotFound APIError = 4040009
	// APIErrorTransactionIDNotFound indicates a transaction identifier wasn't found
	APIErrorTransactionIDNotFound APIError = 4040010
	// APIErrorRateLimitExceeded indicates that the request exceeded the rate limit
	APIErrorRateLimitExceeded APIError = 4290000
	// APIErrorGeneralInternal indicates a general internal error
	APIErrorGeneralInternal APIError = 5000000
)
