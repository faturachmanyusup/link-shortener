package route

import (
	linkController "github.com/faturachmanyusup/link-shortener/controller"
	"github.com/gin-gonic/gin"
)

func Link(ctx *gin.RouterGroup) {
	ctx.GET("/", linkController.ShowForm)
	ctx.GET("/:link", linkController.Find)
	ctx.POST("/", linkController.Create)
}
