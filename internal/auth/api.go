package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	authHeader = strings.TrimSpace(authHeader)
	if strings.HasPrefix(authHeader, "ApiKey ") {
		authHeader = strings.TrimPrefix(authHeader, "ApiKey ")
		if authHeader == "" {
			return "", errors.New("invalid authorization header, missing api key")
		}
		return authHeader, nil
	}

	return "", errors.New("invalid authorization header, missing api key")
}
