package link

import (
	"context"
	"time"

	"github.com/faturachmanyusup/link-shortener/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	OriginalLink string             `bson:"originalLink,omitempty"`
	ShortLink    string             `bson:"shortLink,omitempty"`
	CreatedAt    time.Time          `bson:"createdAt,omitempty"`
}

var connect, _ = config.Connect()
var coll = connect.Collection("links")

func Find(ctx context.Context, filter Link) ([]Link, error) {
	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	links := []Link{}

	if err = cur.All(ctx, &links); err != nil {
		return nil, err
	}

	return links, nil
}

func FindOne(ctx context.Context, filter Link) (Link, error) {
	cur := coll.FindOne(ctx, filter)
	link := Link{}

	if err := cur.Decode(&link); err != nil {
		return link, err
	}

	return link, nil
}

func Create(ctx context.Context, newData Link) (Link, error) {
	_, err := coll.InsertOne(ctx, newData)
	if err != nil {
		return Link{}, err
	}

	return newData, nil
}
