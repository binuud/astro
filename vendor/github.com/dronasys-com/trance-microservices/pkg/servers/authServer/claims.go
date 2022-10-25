package authServer

import (
	"github.com/dgrijalva/jwt-go"
)

// TODO Have single type of Claim, use generic construct like Sub, Sub1, Sub2
// And use entityType enum to differentiate between bot, user, agents etc
type UserClaims struct {
	jwt.StandardClaims
	Issuer string `json:"issuer"`
	Uuid   string `json:"uuid"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type ClusterClaims struct {
	jwt.StandardClaims
	Issuer string `json:"issuer"`
	Uuid   string `json:"uuid"`
	Name   string `json:"name"`
}
