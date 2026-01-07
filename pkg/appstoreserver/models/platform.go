// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

// Platform represents the platform on which the customer consumed the in-app purchase.
// https://developer.apple.com/documentation/appstoreserverapi/platform
type Platform int

const (
	// PlatformUndeclared indicates the platform is undeclared
	PlatformUndeclared Platform = 0
	// PlatformApple indicates the in-app purchase was consumed on an Apple platform
	PlatformApple Platform = 1
	// PlatformNonApple indicates the in-app purchase was consumed on a non-Apple platform
	PlatformNonApple Platform = 2
)
