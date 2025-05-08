package repositories

import (
	"context"
	"crud-mvc-with-go-fiber/config"
	"crud-mvc-with-go-fiber/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.DB.Collection("users")

func GetAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users []models.User
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return users, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}