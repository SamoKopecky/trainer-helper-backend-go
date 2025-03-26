package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"trainer-helper/config"
	"trainer-helper/model"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const MasterTokenEndpoint = "realms/master/protocol/openid-connect/token"

// TODO: Move to schemas
type KeycloakUser struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type RoleInfo struct {
	Id string `json:"id"`
}

func (ku KeycloakUser) FullName() string {
	return fmt.Sprintf("%s %s", ku.FirstName, ku.LastName)
}

func (ku KeycloakUser) ToPersonModel() model.Person {
	return model.Person{
		Id:    ku.Id,
		Name:  ku.FullName(),
		Email: ku.Email,
	}
}

type IAM struct {
	AppConfig  config.Config
	AuthConfig clientcredentials.Config
}

func CreateAuthConfig(appConfig config.Config) clientcredentials.Config {
	return clientcredentials.Config{
		ClientID:     appConfig.KeycloakAdminClientId,
		ClientSecret: appConfig.KeycloakAdminClientSecret,
		TokenURL:     fmt.Sprintf("%s/%s", appConfig.KeycloakBaseUrl, MasterTokenEndpoint),
	}
}

func (i IAM) getUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s", i.AppConfig.KeycloakBaseUrl, endpoint)
}

func (i IAM) getUserUrl() string {
	return i.getUrl(fmt.Sprintf("admin/realms/%s/users", i.AppConfig.KeycloakRealm))
}

func (i IAM) getRoleUrl(role string) string {
	return i.getUrl(fmt.Sprintf("admin/realms/%s/roles/%s", i.AppConfig.KeycloakRealm, role))
}

func (i IAM) authedRequest(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := oauth2.NewClient(context.Background(), i.AuthConfig.TokenSource(context.Background()))
	return client.Do(request)
}

func responseData[T any](response *http.Response) (data T, err error) {
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &data)
	return

}

func (i IAM) GetUsers() ([]KeycloakUser, error) {
	resp, err := i.authedRequest(i.getUserUrl())
	if err != nil {
		return nil, err
	}

	return responseData[[]KeycloakUser](resp)
}

func (i IAM) GetUserById(userId string) (KeycloakUser, error) {
	var user KeycloakUser

	resp, err := i.authedRequest(fmt.Sprintf("%s/%s", i.getUserUrl(), userId))
	if err != nil {
		return user, err
	}

	return responseData[KeycloakUser](resp)
}

func (i IAM) GetUsersByRole(role string) ([]KeycloakUser, error) {
	resp, err := i.authedRequest(fmt.Sprintf("%s/users", i.getRoleUrl(role)))

	if err != nil {
		return nil, err
	}

	return responseData[[]KeycloakUser](resp)
}
