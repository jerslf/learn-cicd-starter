package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name:    "no authorization header",
			headers: http.Header{},
			want:    "",
			wantErr: true,
		},
		{
			name: "malformed authorization header",
			headers: http.Header{
				"Authorization": []string{"Bearer sometoken"},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "valid api key",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-secret-key"},
			},
			want:    "my-secret-key",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := auth.GetAPIKey(tt.headers)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tt.want {
				t.Errorf("GetAPIKey() = %q, want %q", got, tt.want)
			}
		})
	}
}
