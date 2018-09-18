package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/soeyusuke/url_shorter/storage/mysql"
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

	if err := c.Bind(&long_url); err != nil {
		return c.JSON(http.StatusBadRequest, CustomBadResponse(err))
	}

	if err := c.Validate(&long_url); err != nil {
		return c.JSON(http.StatusBadRequest, CustomBadResponse(err))
	}

	m := mysql.Init()
	b, err := m.Save(long_url.Url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, CustomBadResponse(err))
	}

	var shorten_url string

	return c.JSON(http.StatusOK, CustomSuccessResponse(shorten_url, long_url.Url))
}
