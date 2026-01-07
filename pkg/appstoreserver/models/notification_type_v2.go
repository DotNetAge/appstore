// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// NotificationTypeV2 represents the type that describes the in-app purchase or external purchase event for which the App Store sends the version 2 notification
// https://developer.apple.com/documentation/appstoreservernotifications/notificationtype
type NotificationTypeV2 string

const (
	// NotificationTypeV2Subscribed indicates a notification for a subscribed event
	// https://developer.apple.com/documentation/appstoreservernotifications/subscribed
	NotificationTypeV2Subscribed NotificationTypeV2 = "SUBSCRIBED"
	
	// NotificationTypeV2DidChangeRenewalPref indicates a notification for a did-change-renewal-pref event
	// https://developer.apple.com/documentation/appstoreservernotifications/did_change_renewal_pref
	NotificationTypeV2DidChangeRenewalPref NotificationTypeV2 = "DID_CHANGE_RENEWAL_PREF"
	
	// NotificationTypeV2DidChangeRenewalStatus indicates a notification for a did-change-renewal-status event
	// https://developer.apple.com/documentation/appstoreservernotifications/did_change_renewal_status
	NotificationTypeV2DidChangeRenewalStatus NotificationTypeV2 = "DID_CHANGE_RENEWAL_STATUS"
	
	// NotificationTypeV2OfferRedeemed indicates a notification for an offer-redeemed event
	// https://developer.apple.com/documentation/appstoreservernotifications/offer_redeemed
	NotificationTypeV2OfferRedeemed NotificationTypeV2 = "OFFER_REDEEMED"
	
	// NotificationTypeV2DidRenew indicates a notification for a did-renew event
	// https://developer.apple.com/documentation/appstoreservernotifications/did_renew
	NotificationTypeV2DidRenew NotificationTypeV2 = "DID_RENEW"
	
	// NotificationTypeV2Expired indicates a notification for an expired event
	// https://developer.apple.com/documentation/appstoreservernotifications/expired
	NotificationTypeV2Expired NotificationTypeV2 = "EXPIRED"
	
	// NotificationTypeV2DidFailToRenew indicates a notification for a did-fail-to-renew event
	// https://developer.apple.com/documentation/appstoreservernotifications/did_fail_to_renew
	NotificationTypeV2DidFailToRenew NotificationTypeV2 = "DID_FAIL_TO_RENEW"
	
	// NotificationTypeV2GracePeriodExpired indicates a notification for a grace-period-expired event
	// https://developer.apple.com/documentation/appstoreservernotifications/grace_period_expired
	NotificationTypeV2GracePeriodExpired NotificationTypeV2 = "GRACE_PERIOD_EXPIRED"
	
	// NotificationTypeV2PriceIncrease indicates a notification for a price-increase event
	// https://developer.apple.com/documentation/appstoreservernotifications/price_increase
	NotificationTypeV2PriceIncrease NotificationTypeV2 = "PRICE_INCREASE"
	
	// NotificationTypeV2Refund indicates a notification for a refund event
	// https://developer.apple.com/documentation/appstoreservernotifications/refund
	NotificationTypeV2Refund NotificationTypeV2 = "REFUND"
	
	// NotificationTypeV2RefundDeclined indicates a notification for a refund-declined event
	// https://developer.apple.com/documentation/appstoreservernotifications/refund_declined
	NotificationTypeV2RefundDeclined NotificationTypeV2 = "REFUND_DECLINED"
	
	// NotificationTypeV2ConsumptionRequest indicates a notification for a consumption-request event
	// https://developer.apple.com/documentation/appstoreservernotifications/consumption_request
	NotificationTypeV2ConsumptionRequest NotificationTypeV2 = "CONSUMPTION_REQUEST"
	
	// NotificationTypeV2RenewalExtended indicates a notification for a renewal-extended event
	// https://developer.apple.com/documentation/appstoreservernotifications/renewal_extended
	NotificationTypeV2RenewalExtended NotificationTypeV2 = "RENEWAL_EXTENDED"
	
	// NotificationTypeV2Revoke indicates a notification for a revoke event
	// https://developer.apple.com/documentation/appstoreservernotifications/revoke
	NotificationTypeV2Revoke NotificationTypeV2 = "REVOKE"
	
	// NotificationTypeV2Test indicates a notification for a test event
	// https://developer.apple.com/documentation/appstoreservernotifications/test
	NotificationTypeV2Test NotificationTypeV2 = "TEST"
	
	// NotificationTypeV2RenewalExtension indicates a notification for a renewal-extension event
	// https://developer.apple.com/documentation/appstoreservernotifications/renewal_extension
	NotificationTypeV2RenewalExtension NotificationTypeV2 = "RENEWAL_EXTENSION"
	
	// NotificationTypeV2RefundReversed indicates a notification for a refund-reversed event
	// https://developer.apple.com/documentation/appstoreservernotifications/refund_reversed
	NotificationTypeV2RefundReversed NotificationTypeV2 = "REFUND_REVERSED"
	
	// NotificationTypeV2ExternalPurchaseToken indicates a notification for an external-purchase-token event
	// https://developer.apple.com/documentation/appstoreservernotifications/external_purchase_token
	NotificationTypeV2ExternalPurchaseToken NotificationTypeV2 = "EXTERNAL_PURCHASE_TOKEN"
	
	// NotificationTypeV2OneTimeCharge indicates a notification for a one-time-charge event
	// https://developer.apple.com/documentation/appstoreservernotifications/one_time_charge
	NotificationTypeV2OneTimeCharge NotificationTypeV2 = "ONE_TIME_CHARGE"
)
