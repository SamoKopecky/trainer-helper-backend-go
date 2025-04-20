package fetcher

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"trainer-helper/config"
	"trainer-helper/model"
	"trainer-helper/utils"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// TODO: Add status code to error
// TODO: Properly log api errors, create dynamic errors that can be read in logs
var ErrUserNotCreated = errors.New("iam: user not created due to invalid status code")
var ErrUserAlreadyExists = errors.New("iam: user already exists")
var ErrUserActionTriggerFailed = errors.New("iam: user trigger failed because of unknown status code")

const masterTokenEndpoint = "realms/master/protocol/openid-connect/token"

var newRequiredActions = []string{"UPDATE_PROFILE", "UPDATE_PASSWORD", "VERIFY_EMAIL"}

func (ku KeycloakUser) Nickname() *string {
	if len(ku.Attributes.Nickname) > 0 {
		return &ku.Attributes.Nickname[0]
	} else {
		return nil
	}
}

func (ku KeycloakUser) FullName() *string {
	if ku.FirstName != nil && ku.LastName != nil {
		name := fmt.Sprintf("%s %s",
			utils.UpperFirstChar(*ku.FirstName),
			utils.UpperFirstChar(*ku.LastName))
		return &name
	}
	return nil
}

func (ku KeycloakUser) ToUserModel() model.User {
	return model.User{
		Id:       ku.Id,
		Email:    ku.Email,
		Name:     ku.FullName(),
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
		TokenURL:     fmt.Sprintf("%s/%s", appConfig.KeycloakBaseUrl, masterTokenEndpoint),
	}
}

func (i IAM) fromBaseUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s", i.AppConfig.KeycloakBaseUrl, endpoint)
}

func (i IAM) editUserRoles(method string, userLocation UserLocation, kcRole KeycloakRole) error {
	buf := createParamsBuf([]KeycloakRole{kcRole})
	resp, err := i.authedRequest(method, fmt.Sprintf("%s/role-mappings/realm", userLocation), &buf)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return err
	}

	return nil
}

func (i IAM) userUrl() string {
	return i.fromBaseUrl(fmt.Sprintf("admin/realms/%s/users", i.AppConfig.KeycloakRealm))
}

func (i IAM) roleUrl(role string) string {
	return i.fromBaseUrl(fmt.Sprintf("admin/realms/%s/roles/%s", i.AppConfig.KeycloakRealm, role))
}

func (i IAM) authedRequest(method, url string, body *bytes.Buffer) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		reqBody = body
	} else {
		reqBody = nil
	}
	request, err := http.NewRequest(method, url, reqBody)

	if err != nil {
		return nil, err
	}
	client := oauth2.NewClient(context.Background(), i.AuthConfig.TokenSource(context.Background()))
	return client.Do(request)
}

func (i IAM) GetUserById(userId string) (KeycloakUser, error) {
	var user KeycloakUser

	resp, err := i.authedRequest(http.MethodGet, fmt.Sprintf("%s/%s", i.userUrl(), userId), nil)
	if err != nil {
		return user, err
	}

	return responseData[KeycloakUser](resp)
}

func (i IAM) GetUserLocationByEmail(email string) (UserLocation, error) {
	baseUrl := fmt.Sprintf("%s", i.userUrl())
	queryParams := map[string]string{"email": email, "exact": "true"}
	resp, err := i.authedRequest(
		http.MethodGet,
		utils.AddQueryParam(baseUrl, queryParams),
		nil)

	if err != nil {
		return "", err
	}

	users, err := responseData[[]KeycloakUser](resp)
	if err != nil {
		return "", err
	}

	// TODO: Properly check len
	user := users[0]

	userLocation := UserLocation(
		fmt.Sprintf("%s/%s", i.userUrl(), user.Id))
	return userLocation, nil
}

func (i IAM) GetUsersByRole(role string) ([]KeycloakUser, error) {
	resp, err := i.authedRequest(http.MethodGet, fmt.Sprintf("%s/users", i.roleUrl(role)), nil)

	if err != nil {
		return nil, err
	}

	return responseData[[]KeycloakUser](resp)
}

func (i IAM) CreateUser(email, username string) (userLocation UserLocation, err error) {
	newUser := NewKeycloakUser{
		Email:           email,
		Username:        username,
		Enabled:         true,
		EmailVerified:   false,
		RequiredActions: newRequiredActions,
	}
	buf := createParamsBuf(newUser)
	resp, err := i.authedRequest(http.MethodPost, i.userUrl(), &buf)
	if err != nil {
		return
	}
	if resp.StatusCode == http.StatusConflict {
		return "", ErrUserAlreadyExists
	}
	if resp.StatusCode != http.StatusCreated {
		return "", ErrUserNotCreated
	}

	userLocation = UserLocation(resp.Header.Get("Location"))
	return
}

func (i IAM) InvokeUserUpdate(userLocation UserLocation) error {
	buf := createParamsBuf(newRequiredActions)

	resp, err := i.authedRequest(http.MethodPut, fmt.Sprintf("%s/execute-actions-email", userLocation), &buf)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return ErrUserActionTriggerFailed
	}
	return nil
}

func (i IAM) AddUserRoles(userLocation UserLocation, kcRole KeycloakRole) error {
	return i.editUserRoles(http.MethodPost, userLocation, kcRole)
}

func (i IAM) RemoveUserRoles(userLocation UserLocation, kcRole KeycloakRole) error {
	return i.editUserRoles(http.MethodDelete, userLocation, kcRole)
}

func (i IAM) GetRole(roleName string) (KeycloakRole, error) {
	resp, err := i.authedRequest(http.MethodGet, i.roleUrl(roleName), nil)
	if err != nil {
		return KeycloakRole{}, err
	}

	return responseData[KeycloakRole](resp)
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

func createParamsBuf(data any) (buf bytes.Buffer) {
	err := json.NewEncoder(&buf).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
	return
}
