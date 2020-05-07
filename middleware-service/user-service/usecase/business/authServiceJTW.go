package business

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"user-service/core/meta"
	"user-service/core/model"
	"user-service/core/service"
	"user-service/env"
)

type authService struct {
	userService service.UserService
	expirationDelta int
	privateKey *rsa.PrivateKey
	publicKey *rsa.PublicKey
}

func NewAuthServiceJWT(userService service.UserService, authConfig env.AuthConfig) service.AuthService {
	return &authService{
		userService,
		authConfig.TokenExpirationDelta,
		getPrivateKey(authConfig.PrivateKeyPath),
		getPublicKey(authConfig.PublicKeyPath),
	}
}

func (s *authService) SingUp(user *model.User) (*meta.AuthToken, error) {
	if err := s.userService.Register(user); err != nil {
		return nil, err
	}
	return s.GenerateToken(user)
}

func (s *authService) Login(user *model.User)(*meta.AuthToken, error) {
	registered, err := s.userService.FetchOne(user.Username); if err != nil {
		return nil, err
	}

	if registered.PasswordHash != user.PasswordHash {
		return nil, service.ErrPasswordInvalid
	}
	if !registered.Confirmed {
		return nil, service.ErrNotConfirmed
	}

	return s.GenerateToken(user)
}

func (s *authService) GenerateToken(user *model.User) (*meta.AuthToken, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	expiresAt := time.Now().Add(time.Hour * time.Duration(s.expirationDelta))
	token.Claims = &jwt.StandardClaims {
		ExpiresAt: expiresAt.Unix(),
		IssuedAt: time.Now().Unix(),
		Issuer: user.UniqueId,
	}
	tokenString, err := token.SignedString(s.privateKey)
	if err != nil {
		return nil, err
	}
	return meta.NewAuthToken(tokenString, expiresAt), nil
}

func (s *authService) PublicKey() *rsa.PublicKey {
	return s.publicKey
}

func (s *authService) Logout(token string) error {
	panic("implement me")
}


func getPrivateKey(keyPath string) *rsa.PrivateKey {
	privateKeyFile, err := os.Open(keyPath)
	if err != nil {
		panic(err)
	}

	pemFileInfo, _ := privateKeyFile.Stat()
	var size int64 = pemFileInfo.Size()
	pemBytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pemBytes)

	data, _ := pem.Decode(pemBytes)

	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS8PrivateKey(data.Bytes); if err != nil {
		panic(err)
	}
	privateKey, ok := privateKeyImported.(*rsa.PrivateKey); if !ok {
		return nil
	}
	return privateKey
}

func getPublicKey(keyPath string) *rsa.PublicKey {
	publicKeyFile, err := os.Open(keyPath)
	if err != nil {
		panic(err)
	}

	pemFileInfo, _ := publicKeyFile.Stat()
	var size int64 = pemFileInfo.Size()
	pemBytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pemBytes)

	data, _ := pem.Decode(pemBytes)

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes); if err != nil {
		panic(err)
	}

	publicKey, ok := publicKeyImported.(*rsa.PublicKey); if !ok {
		return nil
	}

	return publicKey
}