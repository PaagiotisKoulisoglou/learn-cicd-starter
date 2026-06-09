package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError string
	}{
		{
			name:          "no auth header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: "no authorization header included",
		},
		{
			name:          "malformed auth header",
			headers:       http.Header{"Authorization": []string{"Bearer token123"}},
			expectedKey:   "",
			expectedError: "malformed authorization header",
		},
		{
			name:          "valid auth header",
			headers:       http.Header{"Authorization": []string{"ApiKey myapikey123"}},
			expectedKey:   "myapikey123",
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)
			if key != tt.expectedKey {
				t.Errorf("expected key %v, got %v", tt.expectedKey, key)
			}
			if tt.expectedError == "" && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if tt.expectedError != "" && (err == nil || err.Error() != tt.expectedError) {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
		})
	}
}
