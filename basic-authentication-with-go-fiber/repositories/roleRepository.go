package repositories

import (
	"basic-authentication-with-go-fiber/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllRoles returns all roles
func GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	cursor, err := roleCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &roles); err != nil {
		return nil, err
	}

	return roles, nil
}

// GetUserCountByRole returns the number of users with a specific role
func GetUserCountByRole(roleID primitive.ObjectID) (int, error) {
	count, err := userCollection.CountDocuments(
		context.Background(),
		bson.M{"role_id": roleID},
	)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

// GetUserCountByRoleID returns the number of users with a specific role ID
func GetUserCountByRoleID(roleID string) (int, error) {
	objID, err := primitive.ObjectIDFromHex(roleID)
	if err != nil {
		return 0, err
	}
	return GetUserCountByRole(objID)
}

// DeleteRole deletes a role by ID
func DeleteRole(roleID string) error {
	objID, err := primitive.ObjectIDFromHex(roleID)
	if err != nil {
		return err
	}

	_, err = roleCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
	return err
}
