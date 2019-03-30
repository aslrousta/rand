package rand

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"

	xerrors "github.com/aslrousta/errors"
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

// RandomBytes generates a slice of random bytes.
func RandomBytes(n int) ([]byte, error) {
	if n < 1 {
		return nil, errors.New("invalid random bytes length")
	}

	data := make([]byte, n)
	if _, err := rand.Read(data); err != nil {
		return nil, xerrors.Wrap(err, "random bytes failed")
	}

	return data, nil
}

// RandomHex generates a random hexadecimal string.
func RandomHex(n int) (string, error) {
	bytes, err := RandomBytes((n + 1) / 2)
	if err != nil {
		return "", xerrors.Wrap(err, "random hex failed")
	}

	enc := hex.EncodeToString(bytes)
	return enc[:n], nil
}

// RandomString generates a random string.
func RandomString(n int, charset Charset) (string, error) {
	alphabet := []byte(makeCharset(charset))
	alphaLen := len(alphabet)
	if alphaLen == 0 {
		return "", errors.New("charset is empty")
	}

	bytes, err := RandomBytes(n)
	if err != nil {
		return "", xerrors.Wrap(err, "random string failed")
	}

	var sb strings.Builder
	for _, b := range bytes {
		sb.WriteByte(alphabet[int(b)%alphaLen])
	}

	return sb.String(), nil
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
