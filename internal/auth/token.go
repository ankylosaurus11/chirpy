package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	authHeader = strings.TrimSpace(authHeader)
	if strings.HasPrefix(authHeader, "Bearer ") {
		authHeader = strings.TrimPrefix(authHeader, "Bearer ")
		if authHeader == "" {
			return "", errors.New("invalid authorization header")
		}
		return authHeader, nil
	}

	return "", errors.New("invalid authorization header")
}

func MakeRefreshToken() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", fmt.Errorf("error reading key, %w", err)
	}
	encodedStr := hex.EncodeToString(key)

	return encodedStr, nil
}
