package service

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

func (ku KeycloakUser) ToPersonModel() model.Person {
	return model.Person{
		Id:    ku.Id,
		Name:  fmt.Sprintf("%s %s", ku.FirstName, ku.LastName),
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

func (i IAM) authedRequest(request *http.Request) (response *http.Response, err error) {
	client := oauth2.NewClient(context.Background(), i.AuthConfig.TokenSource(context.Background()))
	response, err = client.Do(request)
	if err != nil {
		return
	}
	return
}

func (i IAM) GetUsers() (users []KeycloakUser, err error) {
	// NOTE: Possible improvments cache tese requests
	usersEndpoint := fmt.Sprintf("admin/realms/%s/users", i.AppConfig.KeycloakRealm)
	req, err := http.NewRequest("GET", i.getUrl(usersEndpoint), nil)
	if err != nil {
		return
	}

	resp, err := i.authedRequest(req)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &users)
	return
}

// TODO: GetUser name for
