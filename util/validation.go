package util

import "go.mongodb.org/mongo-driver/bson/primitive"

func ValidateObjectID(id string) error {
	// ObjectIDFromHex
	_, err := primitive.ObjectIDFromHex(id)

	// if err
	if err != nil {
		return err
	}

	return nil
}
