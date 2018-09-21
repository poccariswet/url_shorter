package base62_test

import (
	"testing"

	"github.com/soeyusuke/url_shorter/base62"
)

func TestEncode(t *testing.T) {
	testData := []struct {
		n int
		s string
	}{
		{1, "1"},
		{12, "c"},
		{123, "1Z"},
		{1234, "jU"},
		{12345, "3d7"},
		{123456, "w7e"},
	}

	for _, v := range testData {
		en := base62.Encode(v.n)
		if en != v.s {
			t.Fatalf("%d encoded expexted %s, but %s", v.n, v.s, en)
		}
	}

}

func TestDecode(t *testing.T) {
	testData := []struct {
		s string
		n int
	}{
		{"1", 1},
		{"c", 12},
		{"1Z", 123},
		{"jU", 1234},
		{"3d7", 12345},
		{"w7e", 123456},
	}

	for _, v := range testData {
		i, err := base62.Decode(v.s)
		if err != nil {
			t.Fatal(err)
		}

		if i != v.n {
			t.Fatalf("%s decodeed expexted %d, but %d", v.s, v.n, i)
		}
	}
}
