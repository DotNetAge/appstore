// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// ConsumptionRequest represents the request body containing consumption information.
// https://developer.apple.com/documentation/appstoreserverapi/consumptionrequest
type ConsumptionRequest struct {
	// CustomerConsented is a Boolean value that indicates whether the customer consented to provide consumption data to the App Store
	CustomerConsented *bool `json:"customerConsented,omitempty"`
	// ConsumptionStatus is a value that indicates the extent to which the customer consumed the in-app purchase
	ConsumptionStatus *ConsumptionStatus `json:"consumptionStatus,omitempty"`
	// Platform is a value that indicates the platform on which the customer consumed the in-app purchase
	Platform *Platform `json:"platform,omitempty"`
	// SampleContentProvided is a Boolean value that indicates whether you provided, prior to its purchase, a free sample or trial of the content
	SampleContentProvided *bool `json:"sampleContentProvided,omitempty"`
	// DeliveryStatus is a value that indicates whether the app successfully delivered an in-app purchase that works properly
	DeliveryStatus *DeliveryStatus `json:"deliveryStatus,omitempty"`
	// AppAccountToken is the UUID that an app optionally generates to map a customer's in-app purchase with its resulting App Store transaction
	AppAccountToken *string `json:"appAccountToken,omitempty"`
	// AccountTenure is the age of the customer's account
	AccountTenure *AccountTenure `json:"accountTenure,omitempty"`
	// PlayTime is a value that indicates the amount of time that the customer used the app
	PlayTime *PlayTime `json:"playTime,omitempty"`
	// LifetimeDollarsRefunded is a value that indicates the total amount, in USD, of refunds the customer has received, in your app, across all platforms
	LifetimeDollarsRefunded *LifetimeDollarsRefunded `json:"lifetimeDollarsRefunded,omitempty"`
	// LifetimeDollarsPurchased is a value that indicates the total amount, in USD, of in-app purchases the customer has made in your app, across all platforms
	LifetimeDollarsPurchased *LifetimeDollarsPurchased `json:"lifetimeDollarsPurchased,omitempty"`
	// UserStatus is the status of the customer's account
	UserStatus *UserStatus `json:"userStatus,omitempty"`
	// RefundPreference is a value that indicates your preference, based on your operational logic, as to whether Apple should grant the refund
	RefundPreference *RefundPreference `json:"refundPreference,omitempty"`
}
