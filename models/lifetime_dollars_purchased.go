// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// LifetimeDollarsPurchased represents the total amount, in USD, of in-app purchases the customer has made in your app, across all platforms.
// https://developer.apple.com/documentation/appstoreserverapi/lifetimedollarspurchased
type LifetimeDollarsPurchased int

const (
	// LifetimeDollarsPurchasedUndeclared indicates the lifetime dollars purchased amount is undeclared
	LifetimeDollarsPurchasedUndeclared LifetimeDollarsPurchased = 0
	// LifetimeDollarsPurchasedZeroDollars indicates the lifetime dollars purchased amount is zero dollars
	LifetimeDollarsPurchasedZeroDollars LifetimeDollarsPurchased = 1
	// LifetimeDollarsPurchasedOneCentToFortyNineDollarsAndNinetyNineCents indicates the lifetime dollars purchased amount is one cent to forty nine dollars and ninety nine cents
	LifetimeDollarsPurchasedOneCentToFortyNineDollarsAndNinetyNineCents LifetimeDollarsPurchased = 2
	// LifetimeDollarsPurchasedFiftyDollarsToNinetyNineDollarsAndNinetyNineCents indicates the lifetime dollars purchased amount is fifty dollars to ninety nine dollars and ninety nine cents
	LifetimeDollarsPurchasedFiftyDollarsToNinetyNineDollarsAndNinetyNineCents LifetimeDollarsPurchased = 3
	// LifetimeDollarsPurchasedOneHundredDollarsToFourHundredNinetyNineDollarsAndNinetyNineCents indicates the lifetime dollars purchased amount is one hundred dollars to four hundred ninety nine dollars and ninety nine cents
	LifetimeDollarsPurchasedOneHundredDollarsToFourHundredNinetyNineDollarsAndNinetyNineCents LifetimeDollarsPurchased = 4
	// LifetimeDollarsPurchasedFiveHundredDollarsToNineHundredNinetyNineDollarsAndNinetyNineCents indicates the lifetime dollars purchased amount is five hundred dollars to nine hundred ninety nine dollars and ninety nine cents
	LifetimeDollarsPurchasedFiveHundredDollarsToNineHundredNinetyNineDollarsAndNinetyNineCents LifetimeDollarsPurchased = 5
	// LifetimeDollarsPurchasedOneThousandDollarsToOneThousandNineHundredNinetyNineDollarsAndNinetyNineCents indicates the lifetime dollars purchased amount is one thousand dollars to one thousand nine hundred ninety nine dollars and ninety nine cents
	LifetimeDollarsPurchasedOneThousandDollarsToOneThousandNineHundredNinetyNineDollarsAndNinetyNineCents LifetimeDollarsPurchased = 6
	// LifetimeDollarsPurchasedTwoThousandDollarsOrGreater indicates the lifetime dollars purchased amount is two thousand dollars or greater
	LifetimeDollarsPurchasedTwoThousandDollarsOrGreater LifetimeDollarsPurchased = 7
)
