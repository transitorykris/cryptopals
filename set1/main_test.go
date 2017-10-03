package set1

import (
	"testing"
)

func TestHexToBase64(t *testing.T) {
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	res, err := HexToBase64(hex)
	if err != nil {
		t.Errorf("Conversion unexpectedly returned an error: %s", err.Error())
	}
	if res != b64 {
		t.Errorf("Base64 was incorrect, got: %s, wanted: %s", res, b64)
	}

	hex2 := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f"
	res, err = HexToBase64(hex2)
	if err != nil {
		t.Errorf("Conversion unexpectedly returned an error: %s", err.Error())
	}
	if res == b64 {
		t.Errorf("Base64 was incorrect, did not expect to get: %s", b64)
	}

	res, err = HexToBase64("This is not hex")
	if err == nil {
		t.Errorf("An invalid string did not generate an error")
	}
	if res != "" {
		t.Errorf("Unexpectedly got a result: %s", res)
	}
}

func TestXOR(t *testing.T) {
	a := "1c0111001f010100061a024b53535009181c"
	b := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"

	res, err := XOR(a, b)
	if err != nil {
		t.Errorf("Unexpected got error: %s", err.Error())
	}
	if res != expected {
		t.Errorf("Did not compute correct XOR, got: %s, expected: %s", res, expected)
	}

	res, err = XOR("1234", "12345")
	if err == nil {
		t.Errorf("Strings of two different lengths should have produced an error")
	}
	if res != "" {
		t.Errorf("Should not have gotten a result for strings of different lengths: %s", res)
	}

	res, err = XOR("This is not hex", "123456789012345")
	if err == nil {
		t.Errorf("An invalid string did not generate an error")
	}
	if res != "" {
		t.Errorf("Unexpectedly got a result: %s", res)
	}

	res, err = XOR("123456789012345", "This is not hex")
	if err == nil {
		t.Errorf("An invalid string did not generate an error")
	}
	if res != "" {
		t.Errorf("Unexpectedly got a result: %s", res)
	}
}

func TestCrackSingleByteXOR(t *testing.T) {
	cipherText := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	text := "Cooking MC's like a pound of bacon"

	res, err := CrackSingleByteXOR(cipherText)
	if err != nil {
		t.Errorf("Unexpectedly got an error: %s", err.Error())
	}
	if res != text {
		t.Errorf("Failed to find original text: %s, got: %s", text, res)
	}

	res, err = CrackSingleByteXOR("This is not hex")
	if err == nil {
		t.Errorf("An invalid string did not generate an error")
	}
	if res != "" {
		t.Errorf("Unexpectedly got a result: %s", res)
	}
}

func TestScanFile(t *testing.T) {
	text := "Now that the party is jumping\n"

	res, err := scanFile("4.txt")
	if err != nil {
		t.Errorf("Unexpectedly got an error: %s", err.Error())
	}
	if res != text {
		t.Errorf("Did not get correct text, expected: %s, got: %s", text, res)
	}

	res, err = scanFile("doesnotexist.txt")
	if err == nil {
		t.Errorf("Expected an error, got: %s", err.Error())
	}
	if res != "" {
		t.Errorf("Should have gotten empty result, got: %s", res)
	}
}

func TestRepeatingXOR(t *testing.T) {
	key := "ICE"
	text := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	cipherText := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	res := RepeatingXOR(key, text)
	if res != cipherText {
		t.Errorf("Incorrect encryption, got: %s, expected: %s", res, cipherText)
	}

	res = RepeatingXOR("WrongKey", text)
	if res == cipherText {
		t.Errorf("Incorrect encryption, did not expect expected: %s", res)
	}

	res = RepeatingXOR(key, "This is not the text we want to encrypt")
	if res == cipherText {
		t.Errorf("Incorrect encryption, did not expect expected: %s", res)
	}
}
