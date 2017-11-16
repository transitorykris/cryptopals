package set1

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/bits"
	"os"
	"strings"
)

// HexToBase64 converts a string given in hex to a base64 encoded string
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

// scanFile checks each of the strings in 4.txt and determines which one
// has been XOR'd by a single character
func scanFile(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	highScore := 0
	highMatch := ""
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cipherText := scanner.Text()
		possibleDecryptedText, err := CrackSingleByteXOR(cipherText)
		if err != nil {
			return "", err
		}
		score := scoreText(possibleDecryptedText)
		if score > highScore {
			highScore = score
			highMatch = possibleDecryptedText
		}
	}
	return highMatch, nil
}

// RepeatingXOR repeatedly XORs key against text by cycling through the bytes in key
func RepeatingXOR(key string, text string) string {
	var cipherText []byte
	for i := 0; i < len(text); i++ {
		cipherText = append(cipherText, byte(text[i]^key[i%3]))
	}
	return hex.EncodeToString(cipherText)
}

// HammingDistance compute the number of differing bits in two strings
func HammingDistance(a, b string) int {
	aBytes := []byte(a)
	bBytes := []byte(b)
	distance := 0

	// Make sure our strings are the same length
	aLen := len(aBytes)
	bLen := len(bBytes)
	if aLen > bLen {
		bBytes = append(bBytes, bytes.Repeat([]byte{0x00}, aLen-bLen)...)
	} else if bLen > aLen {
		aBytes = append(aBytes, bytes.Repeat([]byte{0x00}, bLen-aLen)...)
	}

	for i := range aBytes {
		distance += bits.OnesCount8(uint8(aBytes[i] ^ bBytes[i]))
	}
	return distance
}
