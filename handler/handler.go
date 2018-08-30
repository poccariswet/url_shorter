package url_shortener

import (
	"net/http"

	"github.com/labstack/echo"
)

func sampleHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "sample handler")
}

func UrlShortenerHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "urlShortenerHandler")
}
