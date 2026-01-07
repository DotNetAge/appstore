package internal

import (
	"encoding/asn1"
	"encoding/base64"
	"fmt"
	"regexp"
)

const (
	// PKCS7OID is the object identifier for PKCS#7
	PKCS7OID = "1.2.840.113549.1.7.2"
	// InAppArray is the tag for the in-app array in the receipt
	InAppArray = 17
	// TransactionIdentifier is the tag for the transaction identifier in the receipt
	TransactionIdentifier = 1703
	// OriginalTransactionIdentifier is the tag for the original transaction identifier in the receipt
	OriginalTransactionIdentifier = 1705
)

// ReceiptUtility provides utility methods for working with App Store receipts
type ReceiptUtility struct{}

// NewReceiptUtility creates a new ReceiptUtility
func NewReceiptUtility() *ReceiptUtility {
	return &ReceiptUtility{}
}

// ExtractTransactionIDFromAppReceipt extracts a transaction ID from an encoded App Receipt
// *NO validation* is performed on the receipt, and any data returned should only be used to call the App Store Server API.
// https://developer.apple.com/documentation/appstorereceipts/verifyreceipt
func (r *ReceiptUtility) ExtractTransactionIDFromAppReceipt(appReceipt string) (string, error) {
	// Decode the base64-encoded receipt
	receiptBytes, err := base64.StdEncoding.DecodeString(appReceipt)
	if err != nil {
		return "", fmt.Errorf("failed to decode app receipt: %w", err)
	}

	// Parse the ASN.1 data
	var receipt asn1.RawValue
	if _, err := asn1.Unmarshal(receiptBytes, &receipt); err != nil {
		return "", fmt.Errorf("failed to unmarshal receipt: %w", err)
	}

	// Check if it's a constructed sequence
	if !receipt.IsCompound || receipt.Tag != asn1.TagSequence {
		return "", fmt.Errorf("invalid receipt format: expected constructed sequence")
	}

	// Parse the PKCS#7 content
	var pkcs7 asn1.RawValue
	if _, err := asn1.Unmarshal(receipt.Bytes, &pkcs7); err != nil {
		return "", fmt.Errorf("failed to unmarshal PKCS#7: %w", err)
	}

	// TODO: Implement full PKCS#7 parsing to extract the transaction ID
	// This is a simplified implementation that needs to be completed
	// For now, we'll return an empty string
	return "", nil
}

// ExtractTransactionIDFromTransactionReceipt extracts a transaction ID from an encoded transactional receipt
// *NO validation* is performed on the receipt, and any data returned should only be used to call the App Store Server API.
func (r *ReceiptUtility) ExtractTransactionIDFromTransactionReceipt(transactionReceipt string) (string, error) {
	// Decode the base64-encoded receipt
	receiptBytes, err := base64.StdEncoding.DecodeString(transactionReceipt)
	if err != nil {
		return "", fmt.Errorf("failed to decode transaction receipt: %w", err)
	}

	receiptStr := string(receiptBytes)

	// Use regex to find the purchase-info field
	purchaseInfoRegex := regexp.MustCompile(`"purchase-info"\s+=\s+"([a-zA-Z0-9+/=]+)";`)
	matches := purchaseInfoRegex.FindStringSubmatch(receiptStr)
	if len(matches) != 2 {
		return "", fmt.Errorf("purchase-info not found in transaction receipt")
	}

	// Decode the purchase-info field
	purchaseInfo, err := base64.StdEncoding.DecodeString(matches[1])
	if err != nil {
		return "", fmt.Errorf("failed to decode purchase-info: %w", err)
	}

	purchaseInfoStr := string(purchaseInfo)

	// Use regex to find the transaction-id field
	transactionIDRegex := regexp.MustCompile(`"transaction-id"\s+=\s+"([a-zA-Z0-9+/=]+)";`)
	transactionMatches := transactionIDRegex.FindStringSubmatch(purchaseInfoStr)
	if len(transactionMatches) != 2 {
		return "", fmt.Errorf("transaction-id not found in purchase-info")
	}

	return transactionMatches[1], nil
}
