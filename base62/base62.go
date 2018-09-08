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

func Encode(n int) string {
	buf := make([]byte, 0)

	for ; n > 0; n = n / Base62 {
		r := math.Mod(float64(n), float64(Base62))
		buf = append([]byte{baseSet[int(r)]}, buf...)
	}

	return string(buf)
}

func Decode(s string) (int, error) {
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
