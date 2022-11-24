package astroServer

import (
	"context"

	"github.com/dronasys-com/astro/pkg/astro"
	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	ASTRO_BUCKET_NAME = "astro"
)

type grpcAstroServer struct {
	astroV1.UnimplementedAstroServer
}

func NewGrpcAstroServer(srv *grpc.Server) *grpcAstroServer {

	server := &grpcAstroServer{}
	astroV1.RegisterAstroServer(srv, server)
	return server
}

func (s *grpcAstroServer) Update(ctx context.Context, in *astroV1.UserUpdateRequest) (*astroV1.UserUpdateResponse, error) {

	logrus.Info("Received Update: ", in)

	chart, err := astro.NewAstroChart(&astroV1.BirthData{
		Tob: &astroV1.TobData{
			D:    in.BirthData.Tob.D,
			M:    in.BirthData.Tob.M,
			Y:    in.BirthData.Tob.Y,
			Hour: in.BirthData.Tob.Hour,
			Min:  in.BirthData.Tob.Min,
			Sec:  in.BirthData.Tob.Sec,
			Zone: in.BirthData.Tob.Zone,
		},
		Pob: &astroV1.PobData{
			City:    in.BirthData.Pob.City,
			State:   in.BirthData.Pob.State,
			Country: in.BirthData.Pob.Country,
			Lat:     in.BirthData.Pob.Lat,
			Long:    in.BirthData.Pob.Long,
		},
	})
	if err != nil {
		logrus.Errorf("Cannot prepare chart err(%v)", err)
	}
	chart.PrintChartTable()

	return &astroV1.UserUpdateResponse{
		Astro: &astroV1.AstroObject{
			BirthData: chart.BirthData,
			Planets: &astroV1.Planets{
				PlanetsMap: chart.Planets.PlanetsMap,
			},
		},
	}, nil

}

func (s *grpcAstroServer) Get(ctx context.Context, in *astroV1.UserGetRequest) (*astroV1.UserGetResponse, error) {
	logrus.Info("Received: Get request ", in)

	// mongo, err := mongo.NewDSMongoSingleton()
	// if err != nil {
	// 	logrus.Error("Cannot connect to DB")
	// 	return nil, err
	// }

	// ownerUid, err := helpers.GetUserUuidFromContext(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// filter := bson.D{{"metadata.uid", in.Uuid}, {"metadata.owner", ownerUid}}
	// err = mongo.FindOne(ASTRO_BUCKET_NAME, filter, astro)
	// if err != nil {
	// 	logrus.Error("Cannot Get Item ", err)
	// 	return nil, err
	// }

	chart, err := astro.NewAstroChart(&astroV1.BirthData{
		Tob: &astroV1.TobData{
			D:    04,
			M:    02,
			Y:    1978,
			Hour: 10,
			Min:  44,
			Sec:  0,
			Zone: "Asia/Kolkata",
		},
		Pob: &astroV1.PobData{
			City: "Chennai",
			Lat:  13.08,
			Long: 80.7,
		},
	})
	if err != nil {
		logrus.Errorf("Cannot prepare chart err(%v)", err)
	}
	chart.PrintChartTable()

	return &astroV1.UserGetResponse{
		Astro: &astroV1.AstroObject{
			BirthData: chart.BirthData,
			Planets: &astroV1.Planets{
				PlanetsMap: chart.Planets.PlanetsMap,
			},
		},
	}, nil

}
