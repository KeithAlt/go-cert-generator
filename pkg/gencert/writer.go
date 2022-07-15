package gencert

import (
	"fmt"
	"os"
)

// certPath defines our directory path to contain our ssl
var certPath string = "etc/ssl"

// createKeyFile creates a key file from bytes
func createKeyFile(b []byte) (*string, error) {
	if validateCertPath() != nil {
		err := createCertPath()
		if err != nil {
			return nil, err
		}
	}
	path := fmt.Sprintf("%s/server.key", certPath)
	err := os.WriteFile(path, b, 0644)
	if err != nil {
		return nil, err
	}
	return &path, nil
}

// createPemFile creates a pem file from bytes
func createPemFile(b []byte) (*string, error) {
	if validateCertPath() != nil {
		err := createCertPath()
		if err != nil {
			return nil, err
		}
	}
	path := fmt.Sprintf("%s/server.pem", certPath)
	err := os.WriteFile(path, b, 0644)
	if err != nil {
		return nil, err
	}
	return &path, nil
}

// validateCertPath checks that the required directories exist for containing our ssl
func validateCertPath() error {
	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		return err
	}
	return nil
}

// createCertPath creates required directories for containing our ssl
func createCertPath() error {
	if err := os.MkdirAll(certPath, os.ModePerm); err != nil {
		return err
	}
	return nil
}
