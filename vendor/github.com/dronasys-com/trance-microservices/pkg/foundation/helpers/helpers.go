package helpers

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// randStr
// generates random string of given len
// used to key gen, secrets, passwords etc
func RandStr(len int) string {
	buff := make([]byte, len)
	rand.Read(buff)
	str := base64.StdEncoding.EncodeToString(buff)
	return str[:len]
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomStringURLSafe(n int) string {
	b := RandStr(n)
	return base64.URLEncoding.EncodeToString([]byte(b))
}

func GetUserUuidFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logrus.Errorf("metadata is not provided")
		return "", status.Errorf(codes.Unauthenticated, "Access token is invalid: ") // force relogin on code Unauthenticated
	}
	uuid := md["uuid"][0]
	return uuid, nil
	// insert cluster into DB
}
