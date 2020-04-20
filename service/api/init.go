package api

import (
	"github.com/rahulsai1999/fiber-api/service/db"
	"go.mongodb.org/mongo-driver/mongo"
)

// initialize database collection for api requests
var collectionBlogs *mongo.Collection

func init() {
	collectionBlogs = db.ConnectClient().Database("golang-api").Collection("blogs")
}
