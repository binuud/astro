package astroServer

import (
	"context"

	astroV1 "github.com/binuud/astro/gen/go/v1/astro"
	"github.com/sirupsen/logrus"
)

const (
	ASTRO_BUCKET_NAME = "charts"
)

type AstroGRPCServer struct {
	astroV1.UnimplementedAstroServer
}

func NewGrpcAstroServer() *AstroGRPCServer {
	return &AstroGRPCServer{}
}

func (s *AstroGRPCServer) Create(ctx context.Context, in *astroV1.UserCreateRequest) (*astroV1.UserCreateResponse, error) {

	logrus.Info("Received Create: ", in)

	astroObject, err := CreateRecord(in)
	if err != nil {
		logrus.Error("Cannot Create Chart entry")
		return nil, err
	}

	return &astroV1.UserCreateResponse{
		Astro: astroObject,
	}, nil

}

func (s *AstroGRPCServer) Update(ctx context.Context, in *astroV1.UserUpdateRequest) (*astroV1.UserUpdateResponse, error) {

	logrus.Info("Received Update: ", in)

	astroObj, err := UpdateRecord(in)
	if err != nil {
		logrus.Error("Cannot Update collection entry")
		return nil, err
	}

	return &astroV1.UserUpdateResponse{
		Astro: astroObj,
	}, nil

}

func (s *AstroGRPCServer) Get(ctx context.Context, in *astroV1.UserGetRequest) (*astroV1.UserGetResponse, error) {
	logrus.Info("Received: Get request ", in)
	astroObject, err := GetRecord(in.Uuid, in.Name)
	if err != nil {
		return nil, err
	}

	return &astroV1.UserGetResponse{
		Astro: astroObject,
	}, nil

}

func (s *AstroGRPCServer) List(ctx context.Context, in *astroV1.UserListRequest) (*astroV1.UserListResponse, error) {

	logrus.Info("Received: Listing request ", in)

	charts, err := ListRecords(in.Pagination)
	if err != nil {
		return nil, err
	}

	return &astroV1.UserListResponse{
		Astro:      charts,
		Pagination: in.Pagination,
	}, nil

}
