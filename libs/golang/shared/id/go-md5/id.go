package md5id

import (
	"crypto/md5"
	"encoding/hex"
	"log"

	typetools "libs/golang/shared/type-tools"
)

// ID type definition
type ID string

// NewID generates an ID from various types of data
func NewID(data interface{}) ID {
	str, err := typetools.ToString(data)
	if err != nil {
		log.Fatalf("Error converting data to string: %v", err)
	}
	return md5Hash(str)
}

// md5Hash generates an MD5 hash from a string
func md5Hash(data string) ID {
	hasher := md5.New()
	hasher.Write([]byte(data))
	hash := hasher.Sum(nil)
	return ID(hex.EncodeToString(hash))
}
