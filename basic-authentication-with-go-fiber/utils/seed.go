package utils

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func SeedRole(db *mongo.Database) {
	rolesCollection := db.Collection("roles")

	// Menambahkan Role jika belum ada
	roles := []interface{}{
		bson.M{"name": "admin"},
		bson.M{"name": "user"},
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
