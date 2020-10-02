package rest

import (
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go.kicksware.com/api/service-common/config"
	"go.kicksware.com/api/service-common/core"
	"go.kicksware.com/api/service-common/util"
)

type authService struct {
	publicKey *rsa.PublicKey
	authEndpoint string
}

func NewAuthService(authConfig config.AuthConfig) core.AuthService {
	return &authService{
		util.GetPublicKey(authConfig.PublicKeyPath),
		authConfig.AuthEndpoint,
	}
}

func (s *authService) PublicKey() *rsa.PublicKey {
	return s.publicKey
}

func (s *authService) Authenticate() (string, error)  {
	resp, err := http.Get(s.authEndpoint); if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body); if err != nil {
		return "", err
	}
	token := make(map[string]interface{})
	err = json.Unmarshal(bytes, &token); if err != nil {
		return "", err
	}
	return token["Token"].(string), nil
}