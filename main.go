package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
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
	e.Logger.SetLevel(log.INFO)
	e.GET("/", handler.SampleHandler)
	e.GET("urlshortener/info", handler.UrlShortenerStatusHandler)
	e.POST("urlshortener", handler.UrlShortenerHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	go func() {
		if err := e.Start(":" + port); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
