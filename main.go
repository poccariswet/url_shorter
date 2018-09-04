package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/soeyusuke/url_shorter/handler"
	"github.com/soeyusuke/url_shorter/storage/mysql"
)

func main() {
	db, err := mysql.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer db.Close()

	e := echo.New()
	e.GET("/", handler.SampleHandler)
	e.GET("urlshortener/info", handler.UrlShortenerStatusHandler)
	e.POST("urlshortener", handler.UrlShortenerHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
