package astroServer

import (
	"context"

	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
)

const (
	ASTRO_BUCKET_NAME = "charts"
)

type grpcAstroServer struct {
	astroV1.UnimplementedAstroServer
}

func NewGrpcAstroServer(srv *grpc.Server) *grpcAstroServer {

	server := &grpcAstroServer{}
	astroV1.RegisterAstroServer(srv, server)
	return server
}

func (s *grpcAstroServer) Create(ctx context.Context, in *astroV1.UserCreateRequest) (*astroV1.UserCreateResponse, error) {

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

func (s *grpcAstroServer) Update(ctx context.Context, in *astroV1.UserUpdateRequest) (*astroV1.UserUpdateResponse, error) {

	logrus.Info("Received Update: ", in)

	filter := bson.D{{"metadata.uid", in.Uuid}}
	astroObj, err := UpdateRecord(in, filter)
	if err != nil {
		logrus.Error("Cannot Update collection entry")
		return nil, err
	}

	return &astroV1.UserUpdateResponse{
		Astro: astroObj,
	}, nil

}

func (s *grpcAstroServer) Get(ctx context.Context, in *astroV1.UserGetRequest) (*astroV1.UserGetResponse, error) {
	logrus.Info("Received: Get request ", in)

	filter := bson.D{{"metadata.uid", in.Uuid}}
	astroObject, err := GetRecord(filter)
	if err != nil {
		return nil, err
	}

	return &astroV1.UserGetResponse{
		Astro: astroObject,
	}, nil

}

func (s *grpcAstroServer) List(ctx context.Context, in *astroV1.UserListRequest) (*astroV1.UserListResponse, error) {

	logrus.Info("Received: Listing request ", in)
	filter := bson.D{{}}

	// filter := bson.D{
	// 	{"planets.planetsmap.moon.rasi", 2},
	// 	{"planets.planetsmap.sun.rasi", 1},
	// 	{"planets.planetsmap.mars.rasi", 10},

	// 	{"planets.planetsmap.jupiter.rasi", 4},
	// }
	projection := bson.D{
		{"metadata", 1},
		{"birthdata", 1},
	}

	charts, err := ListRecords(filter, projection, in.Pagination)
	if err != nil {
		return nil, err
	}

	return &astroV1.UserListResponse{
		Astro:      charts,
		Pagination: in.Pagination,
	}, nil

}
