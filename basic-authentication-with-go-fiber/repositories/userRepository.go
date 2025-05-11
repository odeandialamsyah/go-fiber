package repositories

import (
	"basic-authentication-with-go-fiber/config"

	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.DB.Collection("users")
var roleCollection *mongo.Collection = config.DB.Collection("roles")
