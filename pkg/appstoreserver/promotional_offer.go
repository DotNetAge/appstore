// Copyright (c) 2023 Apple Inc. Licensed under MIT License.

package appstoreserver

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"strings"
)

// PromotionalOfferSignatureCreator creates signatures for promotional offers
// https://developer.apple.com/documentation/storekit/in-app_purchase/original_api_for_in-app_purchase/subscriptions_and_offers/generating_a_signature_for_promotional_offers
type PromotionalOfferSignatureCreator struct {
	signingKey *ecdsa.PrivateKey
	keyID      string
	bundleID   string
}

// NewPromotionalOfferSignatureCreator creates a new PromotionalOfferSignatureCreator
func NewPromotionalOfferSignatureCreator(signingKey []byte, keyID, bundleID string) (*PromotionalOfferSignatureCreator, error) {
	// Parse the PEM-encoded private key
	block, _ := pem.Decode(signingKey)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("invalid PEM block: expected EC PRIVATE KEY")
	}

	// Parse the EC private key
	privKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse EC private key: %w", err)
	}

	return &PromotionalOfferSignatureCreator{
		signingKey: privKey,
		keyID:      keyID,
		bundleID:   bundleID,
	}, nil
}

// CreateSignature creates a signature for a promotional offer
// https://developer.apple.com/documentation/storekit/in-app_purchase/original_api_for_in-app_purchase/subscriptions_and_offers/generating_a_signature_for_promotional_offers
func (c *PromotionalOfferSignatureCreator) CreateSignature(productIdentifier, subscriptionOfferID, applicationUsername string, nonce string, timestamp int64) (string, error) {
	// Create the payload
	payload := fmt.Sprintf("%s\u2063%s\u2063%s\u2063%s\u2063%s\u2063%s\u2063%d",
		c.bundleID,
		c.keyID,
		productIdentifier,
		subscriptionOfferID,
		strings.ToLower(applicationUsername),
		strings.ToLower(nonce),
		timestamp,
	)

	// Hash the payload
	hash := sha256.Sum256([]byte(payload))

	// Sign the hash
	r, s, err := ecdsa.Sign(nil, c.signingKey, hash[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign payload: %w", err)
	}

	// Combine the r and s values into a single byte slice
	signature := make([]byte, 64)
	r.FillBytes(signature[:32])
	s.FillBytes(signature[32:])

	// Base64 encode the signature
	return base64.StdEncoding.EncodeToString(signature), nil
}
