package utils

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// GenerateOrderID returns a generated order ID with prefix "ORDER-<UUID>".
func GenerateOrderID() string {
	newUUID := uuid.New()
	cleanUUID := strings.Replace(newUUID.String(), "-", "", -1)
	return fmt.Sprintf("ORDER-%s", cleanUUID)
}
