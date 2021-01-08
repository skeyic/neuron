package utils

import "github.com/google/uuid"

// GenerateUUID ...
func GenerateUUID() string {
	newUUID, _ := uuid.NewRandom()
	return newUUID.String()
}
