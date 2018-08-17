package main

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/urlshortener", UrlShortener)
	r.Use(cors.Default())

	r.Run()
}
