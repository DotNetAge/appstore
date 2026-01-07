package pkg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/DotNetAge/appstore/internal"
	"github.com/DotNetAge/appstore/internal/models"
)

type AppStoreServerClient struct {
	client   *internal.AppStoreServerAPIClient
	verifier *internal.SignedDataVerifier
}

func NewAppStoreServerClient(
	keyID, issuerID, bundleID string, appID int64,
	signingKey []byte,
	rootCertPath string,
	environment models.Environment) *AppStoreServerClient {

	baseClient, err := internal.NewAppStoreServerAPIClient(
		signingKey,
		keyID,
		issuerID,
		bundleID,
		environment)
	if err != nil {
		panic(fmt.Sprintf("failed to create AppStoreServerAPIClient: %v", err))
	}

	// 加载根证书
	rootCerts, err := loadRootCertificates(rootCertPath)
	if err != nil {
		panic(fmt.Sprintf("failed to load root certificates: %v", err))
	}

	// 初始化验证器
	verifier, err := internal.NewSignedDataVerifier(rootCerts, false, environment, bundleID, &appID)
	if err != nil {
		panic(fmt.Sprintf("failed to create signed data verifier: %v", err))
	}

	return &AppStoreServerClient{
		client:   baseClient,
		verifier: verifier,
	}
}

// loadRootCertificates loads root certificates from a directory.
func loadRootCertificates(path string) ([][]byte, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var certs [][]byte
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".cer" {
			certPath := filepath.Join(path, file.Name())
			certData, err := os.ReadFile(certPath)
			if err != nil {
				return nil, fmt.Errorf("failed to read certificate file %s: %w", certPath, err)
			}
			certs = append(certs, certData)
		}
	}

	return certs, nil
}
