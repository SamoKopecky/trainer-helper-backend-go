package fetcher

type KeycloakAttributes struct {
	Nickname []string `json:"nickname"`
}

type KeycloakUser struct {
	Id         string             `json:"id"`
	FirstName  string             `json:"firstName"`
	LastName   string             `json:"lastName"`
	Email      string             `json:"email"`
	Attributes KeycloakAttributes `json:"attributes"`
}

type NewKeycloakUser struct {
	Email           string   `json:"email"`
	Username        string   `json:"username"`
	Enabled         bool     `json:"enabled"`
	EmailVerified   bool     `json:"emailVerified"`
	RequiredActions []string `json:"requiredActions"`
}
