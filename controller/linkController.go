package linkController

import (
	"errors"
	"html/template"
	"strings"
	"time"

	link "github.com/faturachmanyusup/link-shortener/collection"
	random "github.com/faturachmanyusup/link-shortener/lib"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Form struct {
	Link string
}

type HTMLdata struct {
	link.Link
	error
}

var tmp, _ = template.ParseGlob("templates/*")

func Find(ctx *gin.Context) {
	linkForm := ctx.Param("link")
	filter := link.Link{ShortLink: linkForm}

	cur, err := link.FindOne(ctx, filter)

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

	ctx.Redirect(302, cur.OriginalLink)
}

func Create(ctx *gin.Context) {
	linkForm := ctx.PostForm("link")
	if linkForm == "" {
		handleView(ctx, link.Link{}, errors.New("link cannot be empty"))
		return
	}

	shortLink := random.Link(6)

	if !strings.Contains(linkForm, "https://") && !strings.Contains(linkForm, "http://") {
		linkForm = "http://" + linkForm
	}

	data := link.Link{
		Id:           primitive.NewObjectID(),
		OriginalLink: linkForm,
		ShortLink:    shortLink,
		CreatedAt:    time.Now(),
	}

	cur, err := link.Create(ctx, data)
	if err != nil {
		message := err.Error()

		if strings.Contains(message, "duplicate") {
			Create(ctx)
		} else {
			handleView(ctx, link.Link{}, errors.New("something went wrong"))
		}

		return
	}

	cur.ShortLink = "https://fierce-stream-80745.herokuapp.com/" + shortLink

	handleView(ctx, cur, nil)
}

func ShowForm(ctx *gin.Context) {
	handleView(ctx, link.Link{}, nil)
}

func handleView(ctx *gin.Context, data link.Link, err error) {
	var HTMLdata = struct {
		Link  link.Link
		Error error
		Test  string
	}{
		Link:  data,
		Error: err,
		Test:  "test",
	}

	tmp.ExecuteTemplate(ctx.Writer, "index.html", HTMLdata)
}
