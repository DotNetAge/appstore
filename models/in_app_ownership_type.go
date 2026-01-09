// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// InAppOwnershipType represents the relationship of the user with the family-shared purchase to which they have access
// https://developer.apple.com/documentation/appstoreserverapi/inappownershiptype
type InAppOwnershipType string

const (
	// InAppOwnershipTypeFamilyShared indicates the user obtained the product through Family Sharing
	// https://developer.apple.com/documentation/appstoreserverapi/inappownershiptype/family_shared
	InAppOwnershipTypeFamilyShared InAppOwnershipType = "FAMILY_SHARED"
	
	// InAppOwnershipTypePurchased indicates the user purchased the product for their own account
	// https://developer.apple.com/documentation/appstoreserverapi/inappownershiptype/purchased
	InAppOwnershipTypePurchased InAppOwnershipType = "PURCHASED"
)
