package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddQueryParam(t *testing.T) {
	testCases := []struct {
		name     string
		rawUrl   string
		expected string
	}{
		{
			name:     "Standard Keycloak URL",
			rawUrl:   "https://keycloak.example.com/auth/admin/realms/myrealm/users/f5b1f4a0-1b3a-4b3a-8b3a-1b3a4b3a8b3a",
			expected: "https://keycloak.example.com/auth/admin/realms/myrealm/users/f5b1f4a0-1b3a-4b3a-8b3a-1b3a4b3a8b3a?email=abc&strict=true",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := AddQueryParam(tc.rawUrl, map[string]string{
				"email": "abc", "strict": "true"})
			assert.Equal(t, tc.expected, actual)
		})
	}
}
