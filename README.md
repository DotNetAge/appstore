# App Store Server Library for Go

This library provides a convenient way to interact with the Apple App Store Server API in Go. It wraps the core API calls and provides a simplified interface for common operations like transaction history retrieval, notification handling, and subscription management.

## Quick Start

### Installation

```bash
go get github.com/DotNetAge/appstore
```

### Client Initialization

To use the library, you first need to create an `AppStoreServerClient` instance with your Apple credentials and certificates.

```go
package main

import (
	"fmt"
	"os"

	"github.com/DotNetAge/appstore"
	"github.com/DotNetAge/appstore/internal/models"
)

func main() {
	// Load your private key
	signingKey, err := os.ReadFile("path/to/your/key.p8")
	if err != nil {
		panic(fmt.Sprintf("Failed to read private key: %v", err))
	}

	// Create client
	client := NewAppStoreServerClient(
		"YOUR_KEY_ID",          // Key ID from Apple Developer Account
		"YOUR_ISSUER_ID",       // Issuer ID from Apple Developer Account
		"YOUR_BUNDLE_ID",       // App Bundle ID
		123456789,               // App ID
		signingKey,              // Private key bytes
		"path/to/certs",         // Path to root certificates directory
		models.EnvironmentSandbox, // Environment (Sandbox or Production)
	)

	// Use client for API operations
	// ...
}
```

### Transaction Operations

#### Get Transaction History

```go
import "time"

// Get transaction history with options
transactions, err := client.GetTransactionHistory(
	"ORIGINAL_TRANSACTION_ID",
	WithStartDate(time.Now().AddDate(0, -1, 0)), // Start date (1 month ago)
	WithEndDate(time.Now()),                      // End date (now)
	WithProductTypes(models.ProductTypeAutoRenewable), // Filter by product type
	WithSort(models.OrderAscending),             // Sort order
)

if err != nil {
	fmt.Printf("Error getting transaction history: %v\n", err)
	return
}

for _, transaction := range transactions {
	fmt.Printf("Transaction ID: %s\n", transaction.TransactionId)
	fmt.Printf("Product ID: %s\n", transaction.ProductId)
	fmt.Printf("Purchase Date: %d\n", transaction.PurchaseDate)
	// ...
}
```

#### Get Transaction Info

```go
// Get info for a specific transaction
transaction, err := client.GetTransactionInfo("TRANSACTION_ID")
if err != nil {
	fmt.Printf("Error getting transaction info: %v\n", err)
	return
}

fmt.Printf("Transaction ID: %s\n", transaction.TransactionId)
fmt.Printf("Status: %d\n", transaction.Status)
```

#### Look Up Order ID

```go
// Look up transactions by order ID
transactions, err := client.LookUpOrderID("ORDER_ID")
if err != nil {
	fmt.Printf("Error looking up order ID: %v\n", err)
	return
}

for _, transaction := range transactions {
	fmt.Printf("Transaction ID: %s\n", transaction.TransactionId)
	// ...
}
```

### Notification Operations

#### Get Notification History

```go
import "context"

// Get notification history with pagination
ctx := context.Background()
notifications, err := client.GetNotificationHistory(
	ctx,
	"", // Empty pagination token for first request
	WithNotificationStartDate(time.Now().AddDate(0, -1, 0).Unix()), // Start date
	WithNotificationEndDate(time.Now().Unix()),                      // End date
	WithNotificationType(models.NotificationTypeV2SubscriptionRenewal), // Filter by type
)

if err != nil {
	fmt.Printf("Error getting notification history: %v\n", err)
	return
}

for _, notification := range notifications {
	fmt.Printf("Notification Type: %s\n", notification.NotificationType)
	fmt.Printf("Subtype: %s\n", notification.Subtype)
	// ...
}
```

#### Test Notification

```go
// Test notification workflow
notification, err := client.TestNotification()
if err != nil {
	fmt.Printf("Error testing notification: %v\n", err)
	return
}

fmt.Printf("Test Notification Type: %s\n", notification.NotificationType)
```

### Subscription Operations

#### Get All Subscription Statuses

```go
import "context"

// Get subscription statuses
ctx := context.Background()
statuses, err := client.GetAllSubscriptionStatuses(
	ctx,
	"ORIGINAL_TRANSACTION_ID",
	[]models.Status{models.StatusActive, models.StatusExpired},
)

if err != nil {
	fmt.Printf("Error getting subscription statuses: %v\n", err)
	return
}

for _, data := range statuses.Data {
	for _, subscription := range data.Subscriptions {
		fmt.Printf("Subscription Status: %d\n", subscription.Status)
		fmt.Printf("Expiry Date: %d\n", subscription.ExpiresDate)
		// ...
	}
}
```

#### Extend Subscription Renewal Date

```go
import "context"

// Extend subscription renewal date
ctx := context.Background()
request := &models.ExtendRenewalDateRequest{
	ExtendByDays: 7,
	Reason:       "Customer service adjustment",
}

response, err := client.ExtendSubscriptionRenewalDate(
	ctx,
	"ORIGINAL_TRANSACTION_ID",
	request,
)

if err != nil {
	fmt.Printf("Error extending renewal date: %v\n", err)
	return
}

fmt.Printf("New Renewal Date: %d\n", *response.NewRenewalDate)
```

### Error Handling

The library returns standard Go errors. It's recommended to handle errors appropriately in your application:

```go
response, err := client.SomeMethod()
if err != nil {
	// Log the error
	fmt.Printf("Error: %v\n", err)
	
	// Handle specific error cases if needed
	// ...
	
	return
}

// Use the response
// ...
```

### Best Practices

1. **Security**: Keep your private key secure and never expose it in your codebase.
2. **Environment**: Use `models.EnvironmentSandbox` for testing and `models.EnvironmentProduction` for production.
3. **Pagination**: When working with paginated endpoints like `GetNotificationHistory`, the library automatically handles pagination for you.
4. **Rate Limiting**: Be aware of Apple's rate limits for API calls and implement appropriate backoff strategies if needed.
5. **Error Handling**: Always check for errors and handle them gracefully.
6. **Certificates**: Ensure your root certificates are up-to-date to avoid verification failures.

### Troubleshooting

- **Certificate Errors**: Make sure your root certificates directory contains valid `.cer` files.
- **Authentication Errors**: Verify your Key ID, Issuer ID, and private key are correct.
- **API Errors**: Check the Apple App Store Server API documentation for specific error codes and meanings.

## License

This library is licensed under the MIT License. See the LICENSE file for details.
