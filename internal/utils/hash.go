package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func IsValidMD5Hash(text string, textHash string) bool {
	if GetMD5Hash(text) == textHash {
		return true
	}
	return false
}
