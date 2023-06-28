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
