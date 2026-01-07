// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// AccountTenure represents the age of the customer's account.
// https://developer.apple.com/documentation/appstoreserverapi/accounttenure
type AccountTenure int

const (
	// AccountTenureUndeclared indicates the account tenure is undeclared
	AccountTenureUndeclared AccountTenure = 0
	// AccountTenureZeroToThreeDays indicates the account tenure is zero to three days
	AccountTenureZeroToThreeDays AccountTenure = 1
	// AccountTenureThreeDaysToTenDays indicates the account tenure is three days to ten days
	AccountTenureThreeDaysToTenDays AccountTenure = 2
	// AccountTenureTenDaysToThirtyDays indicates the account tenure is ten days to thirty days
	AccountTenureTenDaysToThirtyDays AccountTenure = 3
	// AccountTenureThirtyDaysToNinetyDays indicates the account tenure is thirty days to ninety days
	AccountTenureThirtyDaysToNinetyDays AccountTenure = 4
	// AccountTenureNinetyDaysToOneHundredEightyDays indicates the account tenure is ninety days to one hundred eighty days
	AccountTenureNinetyDaysToOneHundredEightyDays AccountTenure = 5
	// AccountTenureOneHundredEightyDaysToThreeHundredSixtyFiveDays indicates the account tenure is one hundred eighty days to three hundred sixty five days
	AccountTenureOneHundredEightyDaysToThreeHundredSixtyFiveDays AccountTenure = 6
	// AccountTenureGreaterThanThreeHundredSixtyFiveDays indicates the account tenure is greater than three hundred sixty five days
	AccountTenureGreaterThanThreeHundredSixtyFiveDays AccountTenure = 7
)
