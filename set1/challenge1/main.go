package challenge1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
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

// CrackSingleByteXOR attempts to crack a single byte XOR cipher text
// http://cryptopals.com/sets/1/challenges/3
func CrackSingleByteXOR(cipherText string) (string, error) {
	var key byte
	key = 0
	cipherBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	highScore := 0
	highMatch := ""
	for i := 0; i < 256; i++ {
		textBytes := make([]byte, len(cipherBytes))
		for i := range []byte(cipherBytes) {
			textBytes[i] = cipherBytes[i] ^ key
		}
		score := scoreText(string(textBytes))
		if score > highScore {
			highScore = score
			highMatch = string(textBytes)
		}
		key++
	}
	return highMatch, nil
}

// scoreText will return a score based on the frequency of the characters
// ETAOINSHRDLU occuring in the given string
func scoreText(text string) int {
	text = strings.ToLower(text)
	etaoin := "etaoinshrdlu "
	score := 0
	for i := 0; i < len(text); i++ {
		if strings.Contains(etaoin, string(text[i])) {
			score++
		}
	}
	return score
}
