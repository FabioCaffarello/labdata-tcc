package gouuid

import (
	"crypto/sha256"
	"encoding/json"

	"github.com/google/uuid"
)

// ID represents a unique identifier.
type ID = string

// hashData calculates the SHA-256 hash of the input data.
func hashData(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// generateUUIDFromHash generates a UUID from a hash.
func generateUUIDFromHash(hash []byte) string {
	combinedHash := sha256.Sum256(hash)
	return uuid.NewSHA1(uuid.Nil, combinedHash[:]).String()
}

// GenerateUUIDFromMap generates a UUID from a map.
// It serializes the map into JSON, calculates the hash of the JSON data,
// and generates a UUID from the hash.
func GenerateUUIDFromMap(data map[string]interface{}) (ID, error) {
	// Serialize the map into JSON
	serializedData, err := json.Marshal(data)
	if err != nil {
		return ID(""), err
	}

	// Calculate the hash of the JSON data
	hash := hashData(serializedData)

	// Generate a UUID from the hash
	resultID := generateUUIDFromHash(hash)

	return ID(resultID), nil
}
