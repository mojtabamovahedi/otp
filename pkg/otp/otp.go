package otp

import (
	"crypto/rand"
	"io"
)

const CodeLength = 6

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

// EncodeToString encodes a random byte slice to a string of digits.
func EncodeToString() (string, error) {
	b := make([]byte, CodeLength)
	n, err := io.ReadAtLeast(rand.Reader, b, CodeLength)
	if err != nil {
		return "", err
	}
	if n != CodeLength {
		return "", io.ErrUnexpectedEOF
	}
	for i := range b {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b), nil
}
