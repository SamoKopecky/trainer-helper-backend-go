package fetcher

type KeycloakAttributes struct {
	Nickname  []string `json:"nickname"`
	TrainerId []string `json:"trainerId"`
}

type KeycloakUser struct {
	Id         string             `json:"id"`
	Email      string             `json:"email"`
	Attributes KeycloakAttributes `json:"attributes"`
	FirstName  *string            `json:"firstName"`
	LastName   *string            `json:"lastName"`
}

type NewKeycloakUser struct {
	Email           string   `json:"email"`
	Username        string   `json:"username"`
	Enabled         bool     `json:"enabled"`
	EmailVerified   bool     `json:"emailVerified"`
	RequiredActions []string `json:"requiredActions"`
}

type KeycloakRole struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserLocation string
