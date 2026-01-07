package internal

import (
	"crypto"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/DotNetAge/appstore/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

// VerificationStatus represents the status of a verification operation
type VerificationStatus int

const (
	// VerificationStatusOK indicates the verification was successful
	VerificationStatusOK VerificationStatus = 0
	// VerificationStatusVerificationFailure indicates the verification failed
	VerificationStatusVerificationFailure VerificationStatus = 1
	// VerificationStatusInvalidAppIdentifier indicates the app identifier is invalid
	VerificationStatusInvalidAppIdentifier VerificationStatus = 2
	// VerificationStatusInvalidCertificate indicates the certificate is invalid
	VerificationStatusInvalidCertificate VerificationStatus = 3
	// VerificationStatusInvalidChainLength indicates the certificate chain length is invalid
	VerificationStatusInvalidChainLength VerificationStatus = 4
	// VerificationStatusInvalidChain indicates the certificate chain is invalid
	VerificationStatusInvalidChain VerificationStatus = 5
	// VerificationStatusInvalidEnvironment indicates the environment is invalid
	VerificationStatusInvalidEnvironment VerificationStatus = 6
)

// VerificationException represents an exception that occurs during verification
type VerificationException struct {
	Status VerificationStatus
	Err    error
}

// Error implements the error interface for VerificationException
func (e *VerificationException) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("verification failed with status %d: %v", e.Status, e.Err)
	}
	return fmt.Sprintf("verification failed with status %d", e.Status)
}

// SignedDataVerifier provides methods for verifying and decoding App Store signed data
type SignedDataVerifier struct {
	chainVerifier      *chainVerifier
	environment        models.Environment
	bundleID           string
	appAppleID         *int64
	enableOnlineChecks bool
}

// NewSignedDataVerifier creates a new SignedDataVerifier
func NewSignedDataVerifier(rootCertificates [][]byte, enableOnlineChecks bool, environment models.Environment, bundleID string, appAppleID *int64) (*SignedDataVerifier, error) {
	if environment == models.EnvironmentProduction && appAppleID == nil {
		return nil, fmt.Errorf("appAppleID is required when the environment is Production")
	}

	chainVerifier, err := newChainVerifier(rootCertificates)
	if err != nil {
		return nil, err
	}

	return &SignedDataVerifier{
		chainVerifier:      chainVerifier,
		environment:        environment,
		bundleID:           bundleID,
		appAppleID:         appAppleID,
		enableOnlineChecks: enableOnlineChecks,
	}, nil
}

// VerifyAndDecodeRenewalInfo verifies and decodes a signedRenewalInfo obtained from the App Store Server API
// https://developer.apple.com/documentation/appstoreserverapi/jwsrenewalinfo
func (v *SignedDataVerifier) VerifyAndDecodeRenewalInfo(signedRenewalInfo string) (*models.JWSRenewalInfoDecodedPayload, error) {
	// Decode the signed object
	decoded, err := v.decodeSignedObject(signedRenewalInfo)
	if err != nil {
		return nil, err
	}

	// Unmarshal the decoded data into the payload struct
	var payload models.JWSRenewalInfoDecodedPayload
	if err := json.Unmarshal(decoded, &payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JWSRenewalInfoDecodedPayload: %w", err)
	}

	// Verify the environment
	if payload.Environment == nil || *payload.Environment != v.environment {
		return nil, &VerificationException{
			Status: VerificationStatusInvalidEnvironment,
		}
	}

	return &payload, nil
}

// VerifyAndDecodeSignedTransaction verifies and decodes a signedTransaction obtained from the App Store Server API
// https://developer.apple.com/documentation/appstoreserverapi/jwstransaction
func (v *SignedDataVerifier) VerifyAndDecodeSignedTransaction(signedTransaction string) (*models.JWSTransactionDecodedPayload, error) {
	// Decode the signed object
	decoded, err := v.decodeSignedObject(signedTransaction)
	if err != nil {
		return nil, err
	}

	// Unmarshal the decoded data into the payload struct
	var payload models.JWSTransactionDecodedPayload
	if err := json.Unmarshal(decoded, &payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JWSTransactionDecodedPayload: %w", err)
	}

	// Verify the bundle ID
	if payload.BundleId == nil || *payload.BundleId != v.bundleID {
		return nil, &VerificationException{
			Status: VerificationStatusInvalidAppIdentifier,
		}
	}

	// Verify the environment
	if payload.Environment == nil || *payload.Environment != v.environment {
		return nil, &VerificationException{
			Status: VerificationStatusInvalidEnvironment,
		}
	}

	return &payload, nil
}

// VerifyAndDecodeNotification verifies and decodes an App Store Server Notification signedPayload
// https://developer.apple.com/documentation/appstoreservernotifications/signedpayload
func (v *SignedDataVerifier) VerifyAndDecodeNotification(signedPayload string) (*models.ResponseBodyV2DecodedPayload, error) {
	// Decode the signed object
	decoded, err := v.decodeSignedObject(signedPayload)
	if err != nil {
		return nil, err
	}

	// Unmarshal the decoded data into the payload struct
	var payload models.ResponseBodyV2DecodedPayload
	if err := json.Unmarshal(decoded, &payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ResponseBodyV2DecodedPayload: %w", err)
	}

	// Verify the bundle ID, app Apple ID, and environment
	var bundleID string
	var appAppleID *int64
	var environment models.Environment

	if payload.Data != nil {
		if payload.Data.BundleId != nil {
			bundleID = *payload.Data.BundleId
		}
		appAppleID = payload.Data.AppAppleId
		if payload.Data.Environment != nil {
			environment = *payload.Data.Environment
		}
	} else if payload.Summary != nil {
		if payload.Summary.BundleId != nil {
			bundleID = *payload.Summary.BundleId
		}
		appAppleID = payload.Summary.AppAppleId
		if payload.Summary.Environment != nil {
			environment = *payload.Summary.Environment
		}
	} else if payload.ExternalPurchaseToken != nil {
		if payload.ExternalPurchaseToken.BundleId != nil {
			bundleID = *payload.ExternalPurchaseToken.BundleId
		}
		appAppleID = payload.ExternalPurchaseToken.AppAppleId
		if payload.ExternalPurchaseToken.ExternalPurchaseId != nil && strings.HasPrefix(*payload.ExternalPurchaseToken.ExternalPurchaseId, "SANDBOX") {
			environment = models.EnvironmentSandbox
		} else {
			environment = models.EnvironmentProduction
		}
	}

	if err := v.verifyNotification(bundleID, appAppleID, environment); err != nil {
		return nil, err
	}

	return &payload, nil
}

// VerifyAndDecodeAppTransaction verifies and decodes a signed AppTransaction
// https://developer.apple.com/documentation/storekit/apptransaction
func (v *SignedDataVerifier) VerifyAndDecodeAppTransaction(signedAppTransaction string) (*models.AppTransaction, error) {
	// Decode the signed object
	decoded, err := v.decodeSignedObject(signedAppTransaction)
	if err != nil {
		return nil, err
	}

	// Unmarshal the decoded data into the app transaction struct
	var appTransaction models.AppTransaction
	if err := json.Unmarshal(decoded, &appTransaction); err != nil {
		return nil, fmt.Errorf("failed to unmarshal AppTransaction: %w", err)
	}

	// Verify the bundle ID and app Apple ID
	if appTransaction.BundleId == nil || *appTransaction.BundleId != v.bundleID {
		return nil, &VerificationException{
			Status: VerificationStatusInvalidAppIdentifier,
		}
	}

	if v.environment == models.EnvironmentProduction && (appTransaction.AppAppleId == nil || *appTransaction.AppAppleId != *v.appAppleID) {
		return nil, &VerificationException{
			Status: VerificationStatusInvalidAppIdentifier,
		}
	}

	// Verify the environment
	if appTransaction.ReceiptType == nil || *appTransaction.ReceiptType != v.environment {
		return nil, &VerificationException{
			Status: VerificationStatusInvalidEnvironment,
		}
	}

	return &appTransaction, nil
}

// verifyNotification verifies the notification's bundle ID, app Apple ID, and environment
func (v *SignedDataVerifier) verifyNotification(bundleID string, appAppleID *int64, environment models.Environment) error {
	// Verify the bundle ID
	if bundleID != v.bundleID {
		return &VerificationException{
			Status: VerificationStatusInvalidAppIdentifier,
		}
	}

	// Verify the app Apple ID for production environment
	if v.environment == models.EnvironmentProduction && (appAppleID == nil || *appAppleID != *v.appAppleID) {
		return &VerificationException{
			Status: VerificationStatusInvalidAppIdentifier,
		}
	}

	// Verify the environment
	if environment != v.environment {
		return &VerificationException{
			Status: VerificationStatusInvalidEnvironment,
		}
	}

	return nil
}

// decodeSignedObject decodes a signed object from the App Store
func (v *SignedDataVerifier) decodeSignedObject(signedObj string) ([]byte, error) {
	// Parse the JWT without verification to get the headers and claims
	parser := jwt.NewParser()
	token, err := parser.Parse(signedObj, func(token *jwt.Token) (interface{}, error) {
		// We'll verify the signature later
		return nil, nil
	})
	if err != nil {
		return nil, &VerificationException{
			Status: VerificationStatusVerificationFailure,
			Err:    err,
		}
	}

	// Get the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, &VerificationException{
			Status: VerificationStatusVerificationFailure,
			Err:    fmt.Errorf("invalid token claims"),
		}
	}

	// For Xcode or LocalTesting environments, skip verification
	if v.environment == models.EnvironmentXcode || v.environment == models.EnvironmentLocalTesting {
		// Marshal the claims back to JSON
		return json.Marshal(claims)
	}

	// Get the headers
	headers := token.Header

	// Check the algorithm
	alg, ok := headers["alg"].(string)
	if !ok || alg != "ES256" {
		return nil, &VerificationException{
			Status: VerificationStatusVerificationFailure,
			Err:    fmt.Errorf("unsupported algorithm: %s", alg),
		}
	}

	// Get the x5c header
	x5c, ok := headers["x5c"].([]interface{})
	if !ok || len(x5c) == 0 {
		return nil, &VerificationException{
			Status: VerificationStatusVerificationFailure,
			Err:    fmt.Errorf("missing or invalid x5c header"),
		}
	}

	// Convert x5c to []string
	certChain := make([]string, len(x5c))
	for i, cert := range x5c {
		certStr, ok := cert.(string)
		if !ok {
			return nil, &VerificationException{
				Status: VerificationStatusVerificationFailure,
				Err:    fmt.Errorf("invalid certificate in x5c header"),
			}
		}
		certChain[i] = certStr
	}

	// Get the signed date
	var signedDate int64
	if sd, ok := claims["signedDate"].(float64); ok {
		signedDate = int64(sd)
	} else if rcd, ok := claims["receiptCreationDate"].(float64); ok {
		signedDate = int64(rcd)
	} else {
		signedDate = time.Now().Unix() * 1000 // Use current time if no signed date
	}

	// Get the effective date for verification
	effectiveDate := time.Now().Unix()
	if !v.enableOnlineChecks {
		effectiveDate = signedDate / 1000
	}

	// Verify the certificate chain and get the signing key
	signingKey, err := v.chainVerifier.verifyChain(certChain, v.enableOnlineChecks, effectiveDate)
	if err != nil {
		return nil, err
	}

	// Parse and verify the token with the signing key
	verifiedToken, err := parser.Parse(signedObj, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, &VerificationException{
			Status: VerificationStatusVerificationFailure,
			Err:    err,
		}
	}

	// Get the verified claims
	verifiedClaims, ok := verifiedToken.Claims.(jwt.MapClaims)
	if !ok || !verifiedToken.Valid {
		return nil, &VerificationException{
			Status: VerificationStatusVerificationFailure,
			Err:    fmt.Errorf("invalid verified token"),
		}
	}

	// Marshal the verified claims to JSON
	return json.Marshal(verifiedClaims)
}

// chainVerifier handles certificate chain verification
type chainVerifier struct {
	r                  []*x509.Certificate
	enableStrictChecks bool
}

// newChainVerifier creates a new chainVerifier
func newChainVerifier(rootCertificates [][]byte) (*chainVerifier, error) {
	r := make([]*x509.Certificate, 0, len(rootCertificates))

	for _, rootCertBytes := range rootCertificates {
		rootCert, err := x509.ParseCertificate(rootCertBytes)
		if err != nil {
			return nil, &VerificationException{
				Status: VerificationStatusInvalidCertificate,
				Err:    fmt.Errorf("failed to parse root certificate: %w", err),
			}
		}
		r = append(r, rootCert)
	}

	if len(r) == 0 {
		return nil, &VerificationException{
			Status: VerificationStatusInvalidCertificate,
			Err:    fmt.Errorf("no valid root certificates provided"),
		}
	}

	return &chainVerifier{
		r:                  r,
		enableStrictChecks: true,
	}, nil
}

// verifyChain verifies the certificate chain and returns the signing key
func (cv *chainVerifier) verifyChain(certificates []string, performOnlineChecks bool, effectiveDate int64) (crypto.PublicKey, error) {
	if len(certificates) != 3 {
		return nil, &VerificationException{
			Status: VerificationStatusInvalidChainLength,
			Err:    fmt.Errorf("expected 3 certificates in chain, got %d", len(certificates)),
		}
	}

	// Decode the certificates
	parsedCerts := make([]*x509.Certificate, len(certificates))
	for i, certStr := range certificates {
		certBytes, err := base64.StdEncoding.DecodeString(certStr)
		if err != nil {
			return nil, &VerificationException{
				Status: VerificationStatusInvalidCertificate,
				Err:    fmt.Errorf("failed to decode certificate %d: %w", i, err),
			}
		}

		cert, err := x509.ParseCertificate(certBytes)
		if err != nil {
			return nil, &VerificationException{
				Status: VerificationStatusInvalidCertificate,
				Err:    fmt.Errorf("failed to parse certificate %d: %w", i, err),
			}
		}

		parsedCerts[i] = cert
	}

	// Create a certificate pool with the trusted roots
	rootPool := x509.NewCertPool()
	for _, rootCert := range cv.r {
		rootPool.AddCert(rootCert)
	}

	// Verify the leaf certificate with the intermediate certificate
	verifyOpts := x509.VerifyOptions{
		Roots:         rootPool,
		Intermediates: x509.NewCertPool(),
		CurrentTime:   time.Unix(effectiveDate, 0),
		KeyUsages:     []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
	}

	// Add the intermediate certificate to the pool
	verifyOpts.Intermediates.AddCert(parsedCerts[1])

	// Verify the leaf certificate
	_, err := parsedCerts[0].Verify(verifyOpts)
	if err != nil {
		return nil, &VerificationException{
			Status: VerificationStatusVerificationFailure,
			Err:    fmt.Errorf("failed to verify leaf certificate: %w", err),
		}
	}

	// Check the OIDs
	if err := cv.checkOID(parsedCerts[0], "1.2.840.113635.100.6.11.1"); err != nil {
		return nil, err
	}

	if err := cv.checkOID(parsedCerts[1], "1.2.840.113635.100.6.2.1"); err != nil {
		return nil, err
	}

	// Perform OCSP checks if enabled
	if performOnlineChecks {
		// TODO: Implement OCSP checks
	}

	// Return the leaf certificate's public key
	return parsedCerts[0].PublicKey, nil
}

// checkOID checks if the certificate has the specified OID
func (cv *chainVerifier) checkOID(cert *x509.Certificate, oidStr string) error {
	// Parse OID string manually
	var oid asn1.ObjectIdentifier
	if _, err := fmt.Sscanf(oidStr, "%d.%d.%d.%d.%d.%d.%d.%d", &oid[0], &oid[1], &oid[2], &oid[3], &oid[4], &oid[5], &oid[6], &oid[7]); err != nil {
		return &VerificationException{
			Status: VerificationStatusVerificationFailure,
			Err:    fmt.Errorf("invalid OID %s: %w", oidStr, err),
		}
	}

	for _, ext := range cert.Extensions {
		if ext.Id.Equal(oid) {
			return nil
		}
	}

	return &VerificationException{
		Status: VerificationStatusVerificationFailure,
		Err:    fmt.Errorf("certificate missing required OID %s", oidStr),
	}
}
