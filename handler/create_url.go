package handler

import (
	"net/url"
	"os"
)

func CreateURL(encoded string) string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	u := &url.URL{
		Scheme: "http",
		Host:   "localhost:" + port,
		Path:   encoded,
	}

	return u.String()
}
