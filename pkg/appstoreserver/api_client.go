// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

// Package appstoreserver provides the App Store Server Library for Go
package appstoreserver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/apple/app-store-server-library-go/pkg/appstoreserver/models"
)

// GetTransactionHistoryVersion represents the version of the Get Transaction History endpoint to use
type GetTransactionHistoryVersion string

const (
	// GetTransactionHistoryVersionV1 represents version 1 of the Get Transaction History endpoint
	GetTransactionHistoryVersionV1 GetTransactionHistoryVersion = "v1"
	// GetTransactionHistoryVersionV2 represents version 2 of the Get Transaction History endpoint
	GetTransactionHistoryVersionV2 GetTransactionHistoryVersion = "v2"
)

// BaseAppStoreServerAPIClient represents the base API client for the App Store Server API
type BaseAppStoreServerAPIClient struct {
	baseURL    string
	signingKey []byte
	keyID      string
	issuerID   string
	bundleID   string
	environment models.Environment
}

// NewBaseAppStoreServerAPIClient creates a new BaseAppStoreServerAPIClient
func NewBaseAppStoreServerAPIClient(signingKey []byte, keyID, issuerID, bundleID string, environment models.Environment) (*BaseAppStoreServerAPIClient, error) {
	if environment == models.EnvironmentXcode {
		return nil, fmt.Errorf("Xcode is not a supported environment for an AppStoreServerAPIClient")
	}

	baseURL := "https://api.storekit-sandbox.itunes.apple.com"
	if environment == models.EnvironmentProduction {
		baseURL = "https://api.storekit.itunes.apple.com"
	} else if environment == models.EnvironmentLocalTesting {
		baseURL = "https://local-testing-base-url"
	}

	return &BaseAppStoreServerAPIClient{
		baseURL:    baseURL,
		signingKey: signingKey,
		keyID:      keyID,
		issuerID:   issuerID,
		bundleID:   bundleID,
		environment: environment,
	}, nil
}

// generateToken generates a JWT token for authenticating with the App Store Server API
func (c *BaseAppStoreServerAPIClient) generateToken() (string, error) {
	// Parse the signing key
	key, err := jwt.ParseECPrivateKeyFromPEM(c.signingKey)
	if err != nil {
		return "", err
	}

	// Create the claims
	claims := jwt.MapClaims{
		"bid": c.bundleID,
		"iss": c.issuerID,
		"aud": "appstoreconnect-v1",
		"exp": time.Now().Add(5 * time.Minute).Unix(),
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	
	// Set the kid header
	token.Header["kid"] = c.keyID

	// Sign the token
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// getFullURL returns the full URL for a given path
func (c *BaseAppStoreServerAPIClient) getFullURL(path string) string {
	return c.baseURL + path
}

// getHeaders returns the headers for a request
func (c *BaseAppStoreServerAPIClient) getHeaders() (map[string]string, error) {
	token, err := c.generateToken()
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"User-Agent":    "app-store-server-library/go/1.5.0",
		"Authorization": "Bearer " + token,
		"Accept":        "application/json",
	}, nil
}

// parseResponse parses a response from the App Store Server API
func (c *BaseAppStoreServerAPIClient) parseResponse(statusCode int, body []byte, destination interface{}) error {
	if 200 <= statusCode && statusCode < 300 {
		if destination == nil {
			return nil
		}
		return json.Unmarshal(body, destination)
	}

	// Parse error response
	var errorResponse struct {
		ErrorCode    int    `json:"errorCode"`
		ErrorMessage string `json:"errorMessage"`
	}

	if err := json.Unmarshal(body, &errorResponse); err != nil {
		return &models.APIException{
			HTTPStatusCode: statusCode,
		}
	}

	return &models.APIException{
		HTTPStatusCode: statusCode,
		RawAPIError:    &errorResponse.ErrorCode,
		APIError:       models.APIError(errorResponse.ErrorCode),
		ErrorMessage:   &errorResponse.ErrorMessage,
	}
}

// AppStoreServerAPIClient represents the synchronous API client for the App Store Server API
type AppStoreServerAPIClient struct {
	*BaseAppStoreServerAPIClient
	httpClient *http.Client
}

// NewAppStoreServerAPIClient creates a new AppStoreServerAPIClient
func NewAppStoreServerAPIClient(signingKey []byte, keyID, issuerID, bundleID string, environment models.Environment) (*AppStoreServerAPIClient, error) {
	baseClient, err := NewBaseAppStoreServerAPIClient(signingKey, keyID, issuerID, bundleID, environment)
	if err != nil {
		return nil, err
	}

	return &AppStoreServerAPIClient{
		BaseAppStoreServerAPIClient: baseClient,
		httpClient:                  &http.Client{},
	}, nil
}

// makeRequest makes a request to the App Store Server API
func (c *AppStoreServerAPIClient) makeRequest(ctx context.Context, path, method string, queryParams url.Values, body interface{}, destination interface{}) error {
	// Get the full URL
	fullURL := c.getFullURL(path)

	// Get the headers
	headers, err := c.getHeaders()
	if err != nil {
		return err
	}

	// Create the request body
	var requestBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return err
		}
		requestBody = bytes.NewBuffer(jsonBody)
		headers["Content-Type"] = "application/json"
	}

	// Create the request
	req, err := http.NewRequestWithContext(ctx, method, fullURL, requestBody)
	if err != nil {
		return err
	}

	// Add headers
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// Add query parameters
	if len(queryParams) > 0 {
		req.URL.RawQuery = queryParams.Encode()
	}

	// Send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse the response
	return c.parseResponse(resp.StatusCode, respBody, destination)
}
