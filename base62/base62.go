package base62

import (
	"errors"
	"math"
	"strings"
)

const (
	Base62  = 62
	baseSet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func Encode(n int) (string, error) {
	if n == 0 {
		return "", errors.New("request num is invalid")
	}

	buf := make([]byte, 0)

	for ; n > 0; n = n / Base62 {
		r := math.Mod(float64(n), float64(Base62))
		buf = append([]byte{baseSet[int(r)]}, buf...)
	}

	return string(buf), nil
}

func Decode(s string) (int, error) {
	if s == "" {
		return 0, errors.New("non value")
	}

	r, pow := 0, 0
	for i, v := range s {
		pow = len(s) - (i + 1)

		pos := strings.IndexRune(baseSet, v)

		if pos == -1 {
			return pos, errors.New("invalid character: " + string(v))
		}

		r += pos * int(math.Pow(float64(Base62), float64(pow)))
	}

	return int(r), nil
}
