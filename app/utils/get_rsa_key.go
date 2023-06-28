package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/pkg/errors"
	"golang-blueprint-clean/app/env"
)

func GetPublicKey() (*rsa.PublicKey, error) {
	data, _ := pem.Decode([]byte(env.RsaPublicKey))
	if data == nil {
		return nil, errors.New("public key file is not in pem format")
	}

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		return nil, errors.Wrap(err, "public key is not in the right format")
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		return nil, errors.New("public key is not in the right format")
	}

	return rsaPub, nil
}

func GetPrivateKey() (*rsa.PrivateKey, error) {
	data, _ := pem.Decode([]byte(env.RsaPrivateKey))
	if data == nil {
		return nil, errors.New("private key file is not in pem format")
	}

	privateKeyImported, err := x509.ParsePKCS8PrivateKey(data.Bytes)

	if err != nil {
		return nil, errors.Wrap(err, "private key is not in the right format")
	}

	return privateKeyImported.(*rsa.PrivateKey), nil
}
