package authServer

import (
	"context"
	"fmt"
	"time"

	"github.com/dronasys-com/trance-microservices/pkg/dataLayer/mongo"
	protoV1 "github.com/dronasys-com/trance-microservices/proto/v1/user/go"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	USER_BUCKET_NAME = "users"
	AUTH_SALT        = "TRANCE_SALT"   // TODO - externalize
	AUTH_PEPPER      = "TRANCE_PEPPER" // TODO - externalize
)

type grpcAuthServer struct {
	protoV1.UnimplementedAuthServer
}

func NewGrpcAuthServer() *grpcAuthServer {
	return &grpcAuthServer{}
}

func (s *grpcAuthServer) Login(ctx context.Context, in *protoV1.LoginRequest) (*protoV1.LoginResponse, error) {

	logrus.Info("Login Request: ", in)
	userWithSecret := &protoV1.UserSecretObject{}
	// insert cluster into DB
	mongo, err := mongo.NewDSMongoSingleton()

	if err != nil {
		logrus.Error("Cannot connect to DB")
		return nil, fmt.Errorf("Login Failed")
	}

	filter := bson.D{
		{"metadata.email", in.Email},
		{"status.status", protoV1.UserStatus_Active},
	}
	err = mongo.FindOne(USER_BUCKET_NAME, filter, userWithSecret)
	if err != nil {
		myErr := fmt.Errorf("Login failed for user (%s)", in.Email)
		logrus.Errorf("%v", myErr)
		return nil, myErr
	}

	if !CheckPassword(in.Password, userWithSecret.Secret.Password) {
		myErr := fmt.Errorf("Login failed for user (%s)", in.Email)
		logrus.Errorf("%v", myErr)
		return nil, myErr
	}

	// generate token
	token, err := s.GenerateJWT(&protoV1.UserObject{
		MetaData: userWithSecret.MetaData,
		Profile:  userWithSecret.Profile,
	})
	if err != nil {
		return nil, err
	}

	filter = bson.D{{"metadata.uuid", userWithSecret.MetaData.Uuid}}
	update := bson.D{{"$set", bson.D{
		{"status.lastlogin", timestamppb.New(time.Now())},
	}}}
	err = mongo.UpdateOne(USER_BUCKET_NAME, filter, update, true) // update last login date
	if err != nil {
		logrus.Error("Cannot Change password")
		return nil, err
	}

	return &protoV1.LoginResponse{
		Response: &protoV1.ResponseStatus{
			Status:  protoV1.ResponseStatus_Success,
			Message: "Login Success",
		},
		User:  userWithSecret.Profile,
		Token: token,
	}, nil

}

func (s *grpcAuthServer) GenerateJWT(user *protoV1.UserObject) (string, error) {
	jwt := NewTranceJWT(TRANCE_TOKEN_SECRET, 24*time.Hour)
	token, err := jwt.Generate(user)
	if err != nil {
		logrus.Errorf("Cannot generate Trance JWT Token (%v)", err)
		return "", err
	}
	return token, err
}

func CheckPassword(password string, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(AUTH_SALT+password+AUTH_PEPPER))
	return err == nil
}

func HashPassword(password string) (*string, error) {
	if len(password) < 8 {
		return nil, fmt.Errorf("Password length should be of min  character")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(AUTH_SALT+password+AUTH_PEPPER), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}
	hPassword := string(hashedPassword)
	return &hPassword, nil
}
