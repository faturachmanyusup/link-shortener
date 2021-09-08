package route

import (
	linkController "github.com/faturachmanyusup/link-shortener/controller"
	"github.com/gin-gonic/gin"
)

func Link(ctx *gin.RouterGroup) {
	ctx.GET("/:link", linkController.Find)
	ctx.POST("/link", linkController.Create)
}
