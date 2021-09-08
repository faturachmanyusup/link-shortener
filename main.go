package main

import (
	"github.com/faturachmanyusup/link-shortener/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.LoadHTMLGlob("templates/*")

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	link := r.Group("/")
	route.Link(link)

	r.Run()
}
