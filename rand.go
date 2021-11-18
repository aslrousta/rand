package rand

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

// Charset is the data type for random alphabet charset.
type Charset int

// List of charsets.
const (
	All = Uppercase | Lowercase | Digit

	Uppercase = 1 << iota
	Lowercase
	Digit
)

// Bytes generates a slice of random bytes.
func Bytes(n int) ([]byte, error) {
	if n < 1 {
		return nil, errors.New("invalid random bytes length")
	}
	data := make([]byte, n)
	if _, err := rand.Read(data); err != nil {
		return nil, fmt.Errorf("random bytes failed: %v", err)
	}
	return data, nil
}

// MustBytes generates a slice of random bytes, but panics on error.
func MustBytes(n int) []byte {
	bytes, err := Bytes(n)
	if err != nil {
		panic(err)
	}
	return bytes
}

// Hex generates a random hexadecimal string.
func Hex(n int) (string, error) {
	bytes, err := Bytes((n + 1) / 2)
	if err != nil {
		return "", fmt.Errorf("random hex failed: %v", err)
	}
	enc := hex.EncodeToString(bytes)
	return enc[:n], nil
}

// MustHex generates a random hexadecimal string, but panics on error.
func MustHex(n int) string {
	hex, err := Hex(n)
	if err != nil {
		panic(err)
	}
	return hex
}

// String generates a random string.
func String(n int, charset Charset) (string, error) {
	alphabet := []byte(makeCharset(charset))
	alphaLen := len(alphabet)
	if alphaLen == 0 {
		return "", errors.New("charset is empty")
	}
	bytes, err := Bytes(n)
	if err != nil {
		return "", fmt.Errorf("random string failed: %v", err)
	}
	var sb strings.Builder
	for _, b := range bytes {
		sb.WriteByte(alphabet[int(b)%alphaLen])
	}
	return sb.String(), nil
}

// MustString generates a random string, but panics on error.
func MustString(n int, charset Charset) string {
	str, err := String(n, charset)
	if err != nil {
		panic(err)
	}
	return str
}

func makeCharset(charset Charset) string {
	var sb strings.Builder
	if charset&Uppercase > 0 {
		sb.WriteString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	if charset&Lowercase > 0 {
		sb.WriteString("abcdefghijklmnopqrstuvwxyz")
	}
	if charset&Digit > 0 {
		sb.WriteString("0123456789")
	}
	return sb.String()
}
