package utils

import (
	"github.com/google/uuid"
)

func CreateID() string {
	return uuid.New().String()
}