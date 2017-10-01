package challenge1

import (
	"encoding/base64"
	"encoding/hex"
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
