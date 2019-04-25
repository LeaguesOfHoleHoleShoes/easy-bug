package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// GenRandomToken 生成 project token
func GenRandomToken() (string, error) {
	tmpB := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, tmpB); err != nil {
		return "", err
	}
	sb := sha256.Sum256(tmpB)
	return hex.EncodeToString(sb[:]), nil
}
