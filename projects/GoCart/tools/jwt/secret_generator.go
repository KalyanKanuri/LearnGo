package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// Hyphenated names are kept so existing invocations keep working.
var hmacKeySizes = map[string]int{
	"HS256":  32,
	"HS384":  48,
	"HS512":  64,
	"HS-256": 32,
	"HS-384": 48,
	"HS-512": 64,
}

func GenSecret(algo string) (string, error) {
	byteSize, ok := hmacKeySizes[algo]
	if !ok {
		return "", fmt.Errorf("unsupported algorithm %q: use HS256, HS384 or HS512", algo)
	}

	secret := make([]byte, byteSize)
	if _, err := rand.Read(secret); err != nil {
		return "", err
	}

	return hex.EncodeToString(secret), nil
}
