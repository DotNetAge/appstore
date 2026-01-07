// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package models

import (
	"fmt"
)

// APIException represents an exception returned by the App Store Server API
type APIException struct {
	// HTTPStatusCode is the HTTP status code returned by the API
	HTTPStatusCode int `json:"-"`
	// APIError is the API error code if available
	APIError APIError `json:"-"`
	// RawAPIError is the raw error code from the API response
	RawAPIError *int `json:"errorCode,omitempty"`
	// ErrorMessage is the error message from the API response
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// Error implements the error interface for APIException
func (e *APIException) Error() string {
	if e.ErrorMessage != nil {
		return *e.ErrorMessage
	}
	return fmt.Sprintf("API error with status code %d", e.HTTPStatusCode)
}
