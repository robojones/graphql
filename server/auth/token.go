package auth

import (
	"crypto/rand"
	"encoding/base64"
)

const TokenLen = 48
const ByteCount = TokenLen / 4 * 3

func GenerateToken() string {
	b := make([]byte, ByteCount)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
