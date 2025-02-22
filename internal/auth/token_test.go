package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	header1 := make(http.Header)
	header2 := make(http.Header)
	header3 := make(http.Header)

	header1.Add("Authorization", "Bearer as324OIjfoksdjfoiepOIUR")
	header2.Add("Authorization", "")
	header3.Add("Authorization", "Spaghetti")
	tests := []struct {
		name    string
		header  http.Header
		wantErr bool
	}{
		{
			name:    "Proper token string",
			header:  header1,
			wantErr: false,
		},
		{
			name:    "Empty header",
			header:  header2,
			wantErr: true,
		},
		{
			name:    "Wrong string type",
			header:  header3,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GetBearerToken(tt.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("unexpected error state: got error %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil && token != "as324OIjfoksdjfoiepOIUR" {
				t.Errorf("unexpected token: got '%s', want '%s'", token, "as324OIjfoksdjfoiepOIUR")
			}
		})
	}
}
