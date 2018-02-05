package util

import (
	"github.com/google/uuid"
)

func IsValidUUID(uuidVal string) (bool, error) {
	_, err := uuid.Parse(uuidVal)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GenerateUUID() string {

	return " "
}
