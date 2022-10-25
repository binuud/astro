package authServer

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	clusterV1 "github.com/dronasys-com/trance-microservices/proto/v1/cluster/go"
	userV1 "github.com/dronasys-com/trance-microservices/proto/v1/user/go"
)

const (
	TRANCE_TOKEN_SECRET = "TRANCEGIBBERISHSECRET" // TODO Externalize later
	TRANCE_TOKEN_ISSUER = "dronasys.com"
)

// Create a JWT with UserClaim and expiry
// Allows JWT to be verified by client
// Import as required

type TranceJWT struct {
	secret        string
	tokenDuration time.Duration
}

type MicroServiceTokenReq struct {
	Uid  string
	Name string
}

func NewTranceJWT(secret string, tokenDuration time.Duration) *TranceJWT {
	return &TranceJWT{secret, tokenDuration}
}

func (tranceJwt *TranceJWT) Generate(user *userV1.UserObject) (string, error) {

	// Create Claims Object
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tranceJwt.tokenDuration).Unix(), // add token expiry duration to now time
		},
		Issuer: TRANCE_TOKEN_ISSUER,
		Uuid:   user.MetaData.Uuid,
		Email:  user.MetaData.Email,
		Name:   user.MetaData.Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tranceJwt.secret))

}

func (tranceJwt *TranceJWT) GenerateForAgent(cluster *clusterV1.ClusterSecretObject) (string, error) {

	// Create Claims Object
	claims := ClusterClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tranceJwt.tokenDuration).Unix(), // add token expiry duration to now time
		},
		Issuer: TRANCE_TOKEN_ISSUER,
		Uuid:   cluster.MetaData.Uid,
		Name:   cluster.MetaData.Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tranceJwt.secret))

}

func (tranceJwt *TranceJWT) GenerateForMicroService(microService *MicroServiceTokenReq) (string, error) {

	// Create Claims Object
	claims := ClusterClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tranceJwt.tokenDuration).Unix(), // add token expiry duration to now time
		},
		Issuer: TRANCE_TOKEN_ISSUER,
		Uuid:   microService.Uid,
		Name:   microService.Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tranceJwt.secret))

}

func (tranceJwt *TranceJWT) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Cannot read/parse token")
			}

			return []byte(tranceJwt.secret), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
