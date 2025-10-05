package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// GenerateHMACSignature generates an HMAC-SHA256 signature for webhook payloads
func GenerateHMACSignature(payload []byte, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	signature := mac.Sum(nil)
	return hex.EncodeToString(signature)
}

// VerifyHMACSignature verifies an HMAC-SHA256 signature
func VerifyHMACSignature(payload []byte, signature, secret string) bool {
	expectedSignature := GenerateHMACSignature(payload, secret)
	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}
