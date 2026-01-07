// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// LifetimeDollarsRefunded represents the dollar amount of refunds the customer has received in your app, since purchasing the app, across all platforms.
// https://developer.apple.com/documentation/appstoreserverapi/lifetimedollarsrefunded
type LifetimeDollarsRefunded int

const (
	// LifetimeDollarsRefundedUndeclared indicates the lifetime dollars refunded amount is undeclared
	LifetimeDollarsRefundedUndeclared LifetimeDollarsRefunded = 0
	// LifetimeDollarsRefundedZeroDollars indicates the lifetime dollars refunded amount is zero dollars
	LifetimeDollarsRefundedZeroDollars LifetimeDollarsRefunded = 1
	// LifetimeDollarsRefundedOneCentToFortyNineDollarsAndNinetyNineCents indicates the lifetime dollars refunded amount is one cent to forty nine dollars and ninety nine cents
	LifetimeDollarsRefundedOneCentToFortyNineDollarsAndNinetyNineCents LifetimeDollarsRefunded = 2
	// LifetimeDollarsRefundedFiftyDollarsToNinetyNineDollarsAndNinetyNineCents indicates the lifetime dollars refunded amount is fifty dollars to ninety nine dollars and ninety nine cents
	LifetimeDollarsRefundedFiftyDollarsToNinetyNineDollarsAndNinetyNineCents LifetimeDollarsRefunded = 3
	// LifetimeDollarsRefundedOneHundredDollarsToFourHundredNinetyNineDollarsAndNinetyNineCents indicates the lifetime dollars refunded amount is one hundred dollars to four hundred ninety nine dollars and ninety nine cents
	LifetimeDollarsRefundedOneHundredDollarsToFourHundredNinetyNineDollarsAndNinetyNineCents LifetimeDollarsRefunded = 4
	// LifetimeDollarsRefundedFiveHundredDollarsToNineHundredNinetyNineDollarsAndNinetyNineCents indicates the lifetime dollars refunded amount is five hundred dollars to nine hundred ninety nine dollars and ninety nine cents
	LifetimeDollarsRefundedFiveHundredDollarsToNineHundredNinetyNineDollarsAndNinetyNineCents LifetimeDollarsRefunded = 5
	// LifetimeDollarsRefundedOneThousandDollarsToOneThousandNineHundredNinetyNineDollarsAndNinetyNineCents indicates the lifetime dollars refunded amount is one thousand dollars to one thousand nine hundred ninety nine dollars and ninety nine cents
	LifetimeDollarsRefundedOneThousandDollarsToOneThousandNineHundredNinetyNineDollarsAndNinetyNineCents LifetimeDollarsRefunded = 6
	// LifetimeDollarsRefundedTwoThousandDollarsOrGreater indicates the lifetime dollars refunded amount is two thousand dollars or greater
	LifetimeDollarsRefundedTwoThousandDollarsOrGreater LifetimeDollarsRefunded = 7
)
