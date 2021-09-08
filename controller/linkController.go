package linkController

import (
	"strings"
	"time"

	link "github.com/faturachmanyusup/link-shortener/collection"
	random "github.com/faturachmanyusup/link-shortener/lib"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Find(ctx *gin.Context) {
	linkForm := ctx.Param("link")
	filter := link.Link{ShortLink: linkForm}

	link, err := link.FindOne(ctx, filter)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			ctx.JSON(400, gin.H{
				"message": "link not found",
			})
		} else {
			ctx.JSON(500, gin.H{
				"message": "something went wrong",
			})
		}

		return
	}

	ctx.Redirect(302, link.OriginalLink)
}

func Create(ctx *gin.Context) {
	linkForm := ctx.PostForm("link")
	if linkForm == "" {
		ctx.JSON(401, gin.H{
			"message": "link cannot be null",
		})
		return
	}
	shortLink := random.Link(6)

	data := link.Link{
		Id:           primitive.NewObjectID(),
		OriginalLink: linkForm,
		ShortLink:    shortLink,
		CreatedAt:    time.Now(),
	}

	link, err := link.Create(ctx, data)
	if err != nil {
		message := err.Error()

		if strings.Contains(message, "duplicate") {
			Create(ctx)
		} else {
			ctx.JSON(500, gin.H{
				"message": "something went wrong",
			})
		}

		return
	}

	ctx.JSON(200, gin.H{
		"link": link,
	})
}
