package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func MakeJWT(userID uuid.UUID, tokenSecret string) (string, error) {
	if tokenSecret == "" {
		return "", errors.New("tokenSecret is empty")
	}

	now := time.Now()

	expirationTime := jwt.NewNumericDate(now.Add(time.Minute * 60))

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: expirationTime,
		Subject:   userID.String(),
	})

	signedToken, err := newToken.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return signedToken, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	if tokenString == "" {
		return uuid.UUID{}, errors.New("tokenString is empty")
	}
	if tokenSecret == "" {
		return uuid.UUID{}, errors.New("tokenSecret is empty")
	}

	claims := &jwt.RegisteredClaims{}

	var parsedToken *jwt.Token
	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error parsing token: %w", err)
	}

	if !parsedToken.Valid {
		return uuid.UUID{}, fmt.Errorf("invalid token")
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error parsing user ID: %w", err)
	}

	return userID, nil
}
