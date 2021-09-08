package main

import (
	"os"

	"github.com/faturachmanyusup/link-shortener/route"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT IS EMPTY")
	}
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	link := r.Group("/")
	route.Link(link)

	r.Run(port)
}
