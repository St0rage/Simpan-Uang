package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateId() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
