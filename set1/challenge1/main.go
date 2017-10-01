package challenge1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// HexToBase64 converts a string given in hex to a base64 encoded string
// http://cryptopals.com/sets/1/challenges/1
func HexToBase64(in string) (string, error) {
	// Convert our string into a slice of hexadecimal bytes
	b, err := hex.DecodeString(in)
	if err != nil {
		return "", err
	}
	// Encode the result as base64
	return base64.RawStdEncoding.EncodeToString(b), nil
}

// XOR computes the XOR of two equal length hex strings
func XOR(a, b string) (string, error) {
	aBytes, err := hex.DecodeString(a)
	if err != nil {
		return "", err
	}
	bBytes, err := hex.DecodeString(b)
	if err != nil {
		return "", err
	}

	if len(aBytes) != len(bBytes) {
		return "", fmt.Errorf("length of strings %d and %d do not match", len(aBytes), len(bBytes))
	}

	res := make([]byte, len(aBytes))
	for i := range res {
		res[i] = aBytes[i] ^ bBytes[i]
	}

	return hex.EncodeToString(res), nil
}
