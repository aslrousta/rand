package rand_test

import (
	"math/rand"
	"testing"
	"time"

	xrand "github.com/aslrousta/rand"

	. "gopkg.in/go-playground/assert.v1"
)

var (
	uppercase = makeAlphabet("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowercase = makeAlphabet("abcdefghijklmnopqrstuvwxyz")
	digits    = makeAlphabet("0123456789")
	all       = joinAlphabets(uppercase, lowercase, digits)

	hexDigits = makeAlphabet("abcdefgh")
	hex       = joinAlphabets(digits, hexDigits)
)

func TestRandomBytes(t *testing.T) {
	check(t, nil, xrand.RandomBytes)
}

func TestRandomString(t *testing.T) {
	check(t, all,
		func(n int) ([]byte, error) {
			s, err := xrand.RandomString(n, xrand.All)
			if err != nil {
				return nil, err
			}
			return []byte(s), nil
		},
	)
	check(t, uppercase,
		func(n int) ([]byte, error) {
			s, err := xrand.RandomString(n, xrand.Uppercase)
			if err != nil {
				return nil, err
			}
			return []byte(s), nil
		},
	)
	check(t, lowercase,
		func(n int) ([]byte, error) {
			s, err := xrand.RandomString(n, xrand.Lowercase)
			if err != nil {
				return nil, err
			}
			return []byte(s), nil
		},
	)
	check(t, digits,
		func(n int) ([]byte, error) {
			s, err := xrand.RandomString(n, xrand.Digit)
			if err != nil {
				return nil, err
			}
			return []byte(s), nil
		},
	)
}

func TestRandomHex(t *testing.T) {
	check(t, hex,
		func(n int) ([]byte, error) {
			s, err := xrand.RandomHex(n)
			if err != nil {
				return nil, err
			}
			return []byte(s), nil
		},
	)
}

type genFunc func(int) ([]byte, error)

type alphabet []byte

func (alpha alphabet) Contains(b byte) bool {
	for _, a := range alpha {
		if b == a {
			return true
		}
	}
	return false
}

func makeAlphabet(s string) alphabet {
	return []byte(s)
}

func joinAlphabets(alphas ...alphabet) alphabet {
	var bytes []byte
	for _, alpha := range alphas {
		bytes = append(bytes, []byte(alpha)...)
	}
	return alphabet(bytes)
}

func check(t *testing.T, alpha alphabet, g genFunc) {
	const maxIter = 20
	const min = 1
	const max = 100

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < maxIter; i++ {
		n := min + rand.Intn(max-min)

		bytes, err := g(n)
		if err != nil {
			t.Fatal(err)
		}

		Equal(t, len(bytes), n)

		if alpha != nil {
			for _, b := range bytes {
				if !alpha.Contains(b) {
					t.Fatal("invalid character")
				}
			}
		}
	}
}
