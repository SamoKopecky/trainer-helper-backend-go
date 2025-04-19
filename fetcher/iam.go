package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"trainer-helper/config"
	"trainer-helper/model"
	"trainer-helper/utils"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const MasterTokenEndpoint = "realms/master/protocol/openid-connect/token"

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

func (ku KeycloakUser) Nickname() *string {
	if len(ku.Attributes.Nickname) > 0 {
		return &ku.Attributes.Nickname[0]
	} else {
		return nil
	}
}

func (ku KeycloakUser) FullName() string {
	return fmt.Sprintf("%s %s",
		utils.UpperFirstChar(ku.FirstName),
		utils.UpperFirstChar(ku.LastName))
}

func (ku KeycloakUser) ToUserModel() model.User {
	return model.User{
		Id:       ku.Id,
		Name:     ku.FullName(),
		Email:    ku.Email,
		Nickname: ku.Nickname(),
	}
}

type IAM struct {
	AppConfig  *config.Config
	AuthConfig clientcredentials.Config
}

func CreateAuthConfig(appConfig *config.Config) clientcredentials.Config {
	return clientcredentials.Config{
		ClientID:     appConfig.KeycloakAdminClientId,
		ClientSecret: appConfig.KeycloakAdminClientSecret,
		TokenURL:     fmt.Sprintf("%s/%s", appConfig.KeycloakBaseUrl, MasterTokenEndpoint),
	}
}

func (i IAM) fromBaseUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s", i.AppConfig.KeycloakBaseUrl, endpoint)
}

func (i IAM) userUrl() string {
	return i.fromBaseUrl(fmt.Sprintf("admin/realms/%s/users", i.AppConfig.KeycloakRealm))
}

func (i IAM) roleUrl(role string) string {
	return i.fromBaseUrl(fmt.Sprintf("admin/realms/%s/roles/%s", i.AppConfig.KeycloakRealm, role))
}

func (i IAM) authedRequest(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := oauth2.NewClient(context.Background(), i.AuthConfig.TokenSource(context.Background()))
	return client.Do(request)
}

func (i IAM) GetUserById(userId string) (KeycloakUser, error) {
	var user KeycloakUser

	resp, err := i.authedRequest(fmt.Sprintf("%s/%s", i.userUrl(), userId))
	if err != nil {
		return user, err
	}

	return responseData[KeycloakUser](resp)
}

func (i IAM) GetUsersByRole(role string) ([]KeycloakUser, error) {
	resp, err := i.authedRequest(fmt.Sprintf("%s/users", i.roleUrl(role)))

	if err != nil {
		return nil, err
	}

	return responseData[[]KeycloakUser](resp)
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
