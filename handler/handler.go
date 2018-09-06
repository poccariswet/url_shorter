package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func RedirectHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "sample handler")
}

func UrlShortenerStatusHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "UrlShortenerStatusHandler")
}

func UrlShortenerHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "urlShortenerHandler")
}
