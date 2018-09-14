package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type URL struct {
	Url string `json:"url" validate max=255`
}

func RedirectHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "sample handler")
}

func UrlShortenerStatusHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "UrlShortenerStatusHandler")
}

func UrlShortenerHandler(c echo.Context) error {
	var long_url URL

	// TODO: validationの追加
	if err := c.Bind(&long_url); err != nil {
		return err
	}

	if err := c.Validate(&long_url); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "urlShortenerHandler: "+long_url.Url)
}
