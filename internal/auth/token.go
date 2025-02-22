package auth

import (
	"errors"
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
