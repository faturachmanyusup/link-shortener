package main

import (
	"github.com/gin-gonic/gin"

	"github.com/faturachmanyusup/link-shortener/route"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	link := r.Group("/")
	route.Link(link)

	r.Run(":80")
}
