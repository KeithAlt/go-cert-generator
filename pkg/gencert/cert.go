package gencert

import (
	"bytes"
)

// Cert defines our SSL certificate
type Cert struct {
	PemPath   string `binding:"required"`
	KeyPath   string `binding:"required"`
	PemBytes  []byte `binding:"required"`
	KeyBytes  []byte `binding:"required"`
	CertBytes []byte `binding:"required"`
}

// Generate creates a new SSL certificate
func Generate() (*Cert, error) {
	key, err := GenerateKey()
	if err != nil {
		return nil, err
	}

	certBytes, err := CreateCertificate(key)
	pemBuf := &bytes.Buffer{}
	err = encodePem(pemBuf, *certBytes)
	if err != nil {
		return nil, err
	}

	var cert Cert
	cert.CertBytes = *certBytes

	certPem := pemBuf
	pemPath, err := createPemFile(certPem.Bytes())
	if err != nil {
		return nil, err
	}

	cert.PemBytes = pemBuf.Bytes()
	cert.PemPath = *pemPath

	pemBuf.Reset()
	err = encodeKey(pemBuf, key)
	if err != nil {
		return nil, err
	}
	cert.PemBytes = pemBuf.Bytes()

	certKey := pemBuf
	keyPath, err := createKeyFile(certKey.Bytes())
	if err != nil {
		return nil, err
	}
	cert.KeyPath = *keyPath

	return &cert, nil
}
