package app

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"trainer-helper/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/jwk"
)

func getKey(cfg config.Config, token *jwt.Token) (any, error) {
	// TODO: Use cache
	url := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/certs", cfg.KeycloakBaseUrl, cfg.KeycloakRealm)
	fmt.Println(url)
	keySet, err := jwk.Fetch(context.Background(), url)
	if err != nil {
		return rsa.PublicKey{}, err
	}

	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return rsa.PublicKey{}, errors.New("expecting JWT header to have a key ID in the kid field")
	}

	key, found := keySet.LookupKeyID(keyID)
	if !found {
		return rsa.PublicKey{}, fmt.Errorf("unable to find key %q", keyID)
	}
	var pubkey rsa.PublicKey
	if err := key.Raw(&pubkey); err != nil {
		return rsa.PublicKey{}, fmt.Errorf("Unable to get the public key. Error: %s", err.Error())
	}
	return &pubkey, nil
}
