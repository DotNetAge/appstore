// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// PlayTime represents the amount of time that the customer used the app.
// https://developer.apple.com/documentation/appstoreserverapi/playtime
type PlayTime int

const (
	// PlayTimeUndeclared indicates the play time is undeclared
	PlayTimeUndeclared PlayTime = 0
	// PlayTimeZeroToFiveMinutes indicates the play time is zero to five minutes
	PlayTimeZeroToFiveMinutes PlayTime = 1
	// PlayTimeFiveToSixtyMinutes indicates the play time is five to sixty minutes
	PlayTimeFiveToSixtyMinutes PlayTime = 2
	// PlayTimeOneToSixHours indicates the play time is one to six hours
	PlayTimeOneToSixHours PlayTime = 3
	// PlayTimeSixHoursToTwentyFourHours indicates the play time is six hours to twenty four hours
	PlayTimeSixHoursToTwentyFourHours PlayTime = 4
	// PlayTimeOneDayToFourDays indicates the play time is one day to four days
	PlayTimeOneDayToFourDays PlayTime = 5
	// PlayTimeFourDaysToSixteenDays indicates the play time is four days to sixteen days
	PlayTimeFourDaysToSixteenDays PlayTime = 6
	// PlayTimeOverSixteenDays indicates the play time is over sixteen days
	PlayTimeOverSixteenDays PlayTime = 7
)
