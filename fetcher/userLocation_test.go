package fetcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserLocation_UserId(t *testing.T) {
	testCases := []struct {
		name     string
		location UserLocation
		expected string
	}{
		{
			name:     "Standard Keycloak URL",
			location: UserLocation("https://keycloak.example.com/auth/admin/realms/myrealm/users/f5b1f4a0-1b3a-4b3a-8b3a-1b3a4b3a8b3a"),
			expected: "f5b1f4a0-1b3a-4b3a-8b3a-1b3a4b3a8b3a",
		},
		{
			name:     "URL with trailing slash",
			location: UserLocation("https://keycloak.example.com/auth/admin/realms/myrealm/users/e2c4g5b1-2c4b-5c4b-9c4b-2c4b5c4b9c4b/"),
			expected: "e2c4g5b1-2c4b-5c4b-9c4b-2c4b5c4b9c4b",
		},
		{
			name:     "Simpler path structure",
			location: UserLocation("/api/v1/users/another-user-id"),
			expected: "another-user-id",
		},
		{
			name:     "Just the ID (not a URL, but path.Base handles it)",
			location: UserLocation("just-an-id"),
			expected: "just-an-id",
		},
		{
			name:     "Empty string input",
			location: UserLocation(""),
			expected: ".", // Documenting path.Base behaviour for empty string
		},
		{
			name:     "Root path input",
			location: UserLocation("/"),
			expected: "/", // Documenting path.Base behaviour for "/"
		},
		{
			name:     "URL without specific path",
			location: UserLocation("https://host.example.com"),
			expected: "host.example.com", // Documenting path.Base behaviour for host only
		},
		{
			name:     "URL with only root path",
			location: UserLocation("https://host.example.com/"),
			expected: "host.example.com",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.location.UserId()
			assert.Equal(t, tc.expected, actual)
		})
	}
}
