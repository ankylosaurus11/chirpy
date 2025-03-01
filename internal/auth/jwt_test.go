package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	userID := uuid.New()
	secret := "testingsecret1"

	tests := []struct {
		name        string
		userID      uuid.UUID
		tokenSecret string
		expiresIn   time.Duration
		wantErr     bool
	}{
		{
			name:        "Valid token with correct secret and 3-minute expiry",
			userID:      userID,
			tokenSecret: secret,
			expiresIn:   3 * time.Minute,
			wantErr:     false,
		},
		{
			name:        "Invalid token due to empty secret",
			userID:      userID,
			tokenSecret: "",
			expiresIn:   1 * time.Second,
			wantErr:     true,
		},
		{
			name:        "Invalid token with 0 expiry time",
			userID:      userID,
			tokenSecret: secret,
			expiresIn:   0,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := MakeJWT(tt.userID, tt.tokenSecret)

			if (err != nil) != tt.wantErr {
				t.Errorf("MakeJWT() error = %v, wantErr = %v", err, tt.wantErr)
			}

			if tt.wantErr {
				if token != "" {
					t.Errorf("Expected empty token, got %v", token)
				}
				return
			}

			validateUserID, err := ValidateJWT(token, tt.tokenSecret)
			if err != nil {
				t.Errorf("ValidateJWT() error = %v, did not expect error", err)
				return
			}

			if validateUserID != tt.userID {
				t.Errorf("ValidateJWT() userID = %v, want %v", validateUserID, tt.userID)
			}
		})
	}
}
