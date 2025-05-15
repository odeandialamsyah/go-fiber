package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"basic-authentication-with-go-fiber/models"
)

func SeedRole(db *mongo.Database) {
	rolesCollection := db.Collection("roles")

	// Menambahkan Role jika belum ada
	roles := []bson.M{
		{"name": "admin"},
		{"name": "user"},
	}

	for _, role := range roles {
		_, err := rolesCollection.UpdateOne(
			context.Background(),
			bson.M{"name": role["name"]},
			bson.M{"$setOnInsert": role},
			options.Update().SetUpsert(true),
		)
		if err != nil {
			log.Fatal("Error seeding roles: ", err)
		}
	}
}

func SeedAdmin(db *mongo.Database) {
	usersCollection := db.Collection("users")

	// Get admin role
	rolesCollection := db.Collection("roles")
	var adminRole models.Role
	err := rolesCollection.FindOne(context.Background(), bson.M{"name": "admin"}).Decode(&adminRole)
	if err != nil {
		log.Fatal("Error finding admin role: ", err)
	}

	// Create admin user if not exists
	adminPassword, err := HashPassword("admin123")
	if err != nil {
		log.Fatal("Error hashing password: ", err)
	}

	adminUser := models.User{
		Username: "admin",
		Password: adminPassword,
		Role:     adminRole,
	}

	_, err = usersCollection.UpdateOne(
		context.Background(),
		bson.M{"username": adminUser.Username},
		bson.M{"$setOnInsert": adminUser},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Fatal("Error seeding admin user: ", err)
	}
}
