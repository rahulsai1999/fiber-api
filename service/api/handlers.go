package api

import (
	"context"
	"log"

	"github.com/gofiber/fiber"
	"github.com/rahulsai1999/fiber-api/service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Ping Custom message
func Ping(ctx *fiber.Ctx) {
	ctx.Status(200).JSON(fiber.Map{
		"message": "Pong",
	})
}

// GetAllBlogs -> gets all blogs
func GetAllBlogs(ctx *fiber.Ctx) {
	filter := bson.D{{}}
	findOptions := options.Find()
	findOptions.SetLimit(5)

	var results []models.Blog
	cur, err := collectionBlogs.Find(context.Background(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem models.Blog
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.Status(200).JSON(fiber.Map{
			"result": results,
		})
	}
}

// GetBlog -> get specific blog
func GetBlog(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	docID, _ := primitive.ObjectIDFromHex(id)
	result := models.Blog{}
	filter := bson.M{"_id": docID}
	err := collectionBlogs.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		ctx.Status(200).JSON(fiber.Map{
			"result": result,
		})
	}
}

// InsertBlog -> insert specific blog
func InsertBlog(ctx *fiber.Ctx) {
	title := ctx.Body("title")
	author := ctx.Body("author")
	body := ctx.Body("body")

	blog := models.Blog{
		Title:  title,
		Author: author,
		Body:   body,
	}

	result, err := collectionBlogs.InsertOne(context.Background(), blog)
	if err != nil {
		log.Fatal(err)
	} else {
		ctx.Status(200).JSON(fiber.Map{
			"inserted_id": result.InsertedID,
		})
	}
}
