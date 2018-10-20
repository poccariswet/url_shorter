package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/poccariswet/url_shorter/base62"
	"github.com/poccariswet/url_shorter/storage/mysql"
)

type URL struct {
	Url string `json:"url" validate:"required,url,max=255"`
}

func RedirectHandler(c echo.Context) error {
	code := c.Param("id")
	decode, err := base62.Decode(code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, CustomBadResponse(err))
	}

	m := mysql.Init()
	url, err := m.LoadAndCountUp(decode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, CustomBadResponse(err))
	}

	return c.Redirect(http.StatusMovedPermanently, url)
}

func UrlShortenerStatusHandler(c echo.Context) error {
	code := c.Param("id")
	decode, err := base62.Decode(code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, CustomBadResponse(err))
	}

	m := mysql.Init()
	url, err := m.FetchInfo(decode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, CustomBadResponse(err))
	}

	return c.JSON(http.StatusOK, url)
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

	return c.JSON(http.StatusOK, CustomSuccessResponse(CreateURL(b), long_url.Url))
}
