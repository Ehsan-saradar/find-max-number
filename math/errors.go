package math

import "errors"

var (
	PublicKeyNotFound = errors.New("Public key not found")
	PrivateKeyNotFound       = errors.New("Private key not found")
	InvalidPublicKey        = errors.New("Invalid public key")
	InvalidPrivateKey         = errors.New("Invalid private key")
)

