package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
)

// GenerateRandomString generates a cryptographically secure random string
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

// GenerateAuthCode generates a random authorization code
func GenerateAuthCode() (string, error) {
	return GenerateRandomString(32)
}

// GenerateRandomAccessToken generates a random access token string
func GenerateRandomAccessToken() (string, error) {
	return GenerateRandomString(64)
}

// GenerateRefreshTokenString generates a random refresh token string
func GenerateRefreshTokenString() (string, error) {
	return GenerateRandomString(64)
}

// GenerateRegistrationNumber generates a student registration number
// Format: YYCCNNNNNNNNN (YY=year, CC=college, N=sequential)
func GenerateRegistrationNumber(year int, collegeCode int, sequence int) string {
	return fmt.Sprintf("%02d%02d%09d", year%100, collegeCode, sequence)
}

// GenerateAdmissionNumber generates an admission number
// Format: MUST/TYPE/YEAR/NUMBER
func GenerateAdmissionNumber(programType string, year int, number int) string {
	return fmt.Sprintf("MUST/%s/%d/%05d", programType, year, number)
}

// GenerateInvoiceNumber generates an invoice number
// Format: MB + 10 digits
func GenerateInvoiceNumber(sequence int) string {
	return fmt.Sprintf("MB%010d", sequence)
}

// GenerateControlNumber generates a payment control number
// Format: 12-13 digit number
func GenerateControlNumber() (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(9999999999999))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%013d", n.Int64()), nil
}
