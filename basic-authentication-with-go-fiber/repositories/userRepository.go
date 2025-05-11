package repositories

import (
	"basic-authentication-with-go-fiber/config"
	"basic-authentication-with-go-fiber/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.DB.Collection("users")
var roleCollection *mongo.Collection = config.DB.Collection("roles")

// Fungsi untuk menambahkan user baru
func CreateUser(user models.User) error {
	_, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}