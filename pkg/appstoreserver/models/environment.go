// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

// Package models provides data models for the App Store Server Library
package models

// Environment represents the server environment, either sandbox or production.
// https://developer.apple.com/documentation/appstoreserverapi/environment
type Environment string

const (
	// EnvironmentSandbox represents the sandbox environment
	EnvironmentSandbox Environment = "Sandbox"
	// EnvironmentProduction represents the production environment
	EnvironmentProduction Environment = "Production"
	// EnvironmentXcode represents the Xcode environment for testing
	EnvironmentXcode Environment = "Xcode"
	// EnvironmentLocalTesting represents the local testing environment
	EnvironmentLocalTesting Environment = "LocalTesting"
)
