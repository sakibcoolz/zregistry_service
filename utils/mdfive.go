package utils

import (
	"crypto/md5"
	"encoding/hex"

	"go.uber.org/zap"
)

func MD5Encryption(log *zap.Logger, data string) string {
	// Create an MD5 hash object
	hash := md5.New()

	// Write the input string to the hash object
	_, err := hash.Write([]byte(data))
	if err != nil {
		log.Error("failed to encryption", zap.Error(err))

		return ""
	}

	// Get the MD5 hash as a byte slice
	hashBytes := hash.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	return hex.EncodeToString(hashBytes)
}
