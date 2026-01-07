// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// ExtendReasonCode represents the code that represents the reason for the subscription-renewal-date extension
// https://developer.apple.com/documentation/appstoreserverapi/extendreasoncode
type ExtendReasonCode int

const (
	// ExtendReasonCodeUndeclared indicates an undeclared reason for the extension
	ExtendReasonCodeUndeclared ExtendReasonCode = 0
	// ExtendReasonCodeCustomerSatisfaction indicates the extension is for customer satisfaction
	ExtendReasonCodeCustomerSatisfaction ExtendReasonCode = 1
	// ExtendReasonCodeOther indicates the extension is for another reason
	ExtendReasonCodeOther ExtendReasonCode = 2
	// ExtendReasonCodeServiceIssueOrOutage indicates the extension is due to a service issue or outage
	ExtendReasonCodeServiceIssueOrOutage ExtendReasonCode = 3
)
