package util

import "github.com/google/uuid"

func GetUUID() string {
	return uuid.New().String()
}

func IsValidUUID(parameter string) bool {
	_, err := uuid.Parse(parameter)
	if err != nil {
		return false
	}

	return true
}
