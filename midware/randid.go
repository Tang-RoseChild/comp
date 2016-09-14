package midware

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateRandID if length is 0,than default length will be 5
// for very very import id, need to check it whether exist
func GenerateRandID(length int) (string, error) {
	buf := make([]byte, length)
	id := make([]byte, hex.EncodedLen(len(buf)))
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	hex.Encode(id, buf)
	return string(id), nil
}
