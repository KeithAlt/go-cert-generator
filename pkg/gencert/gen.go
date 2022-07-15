package gencert

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"time"
)

// getKeyType returns the key type
func getKeyType(key interface{}) interface{} {
	switch k := key.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}

// readPemForKey returns the pem block for a key
func readPemForKey(key interface{}) (*pem.Block, error) {
	switch k := key.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}, nil
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal ECDSA private key: %w", err)
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}, nil
	default:
		return nil, errors.New("failed to determine key type")
	}
}

// GenerateKey generates a new key
func GenerateKey() (*ecdsa.PrivateKey, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// CreateCertificate creates a new certificate from a provided key
func CreateCertificate(key *ecdsa.PrivateKey) (*[]byte, error) {
	template := getTemplate()
	certBytes, err := x509.CreateCertificate(rand.Reader, template, template, getKeyType(key), key)
	if err != nil {
		return nil, err
	}
	return &certBytes, nil
}

// getTemplate returns our key data
func getTemplate() *x509.Certificate {
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(time.Hour * 24 * 180),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	return &template
}

// encodePem encodes a certification bytes into a byte buffer
func encodePem(pemBuf *bytes.Buffer, certBytes []byte) error {
	err := pem.Encode(pemBuf, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	if err != nil {
		return err
	}
	return nil
}

// encodeKey encodes a pem block into a byte buffer
func encodeKey(pemBuf *bytes.Buffer, key *ecdsa.PrivateKey) error {
	pemBlock, err := readPemForKey(key)
	if err != nil {
		return err
	}
	err = pem.Encode(pemBuf, pemBlock)
	if err != nil {
		return err
	}
	return nil
}
