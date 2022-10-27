package astroServer

import (
	"context"

	"github.com/dronasys-com/astro/pkg/astro"
	astroV1 "github.com/dronasys-com/astro/proto/v1/astrocalc/go"
	"github.com/dronasys-com/trance-microservices/pkg/dataLayer/mongo"
	"github.com/dronasys-com/trance-microservices/pkg/foundation/helpers"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
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

func (s *grpcAstroServer) GetPanchang(ctx context.Context, in *astroV1.PanchangGetRequest) (*astroV1.PanchangResponse, error) {

	logrus.Info("Received GetPanchang: ", in)

	return &astroV1.PanchangResponse{}, nil

}

func (s *grpcAstroServer) Update(ctx context.Context, in *astroV1.UserUpdateRequest) (*astroV1.UserUpdateResponse, error) {

	logrus.Info("Received Update: ", in)

	chart, err := astro.NewAstroChart(&astroV1.BirthData{
		Dob: &astroV1.TobData{
			D:    in.BirthData.Dob.D,
			M:    in.BirthData.Dob.M,
			Y:    in.BirthData.Dob.Y,
			Hour: in.BirthData.Dob.Hour,
			Min:  in.BirthData.Dob.Min,
			Sec:  in.BirthData.Dob.Sec,
			Zone: in.BirthData.Dob.Zone,
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
		logrus.Error("Cannot prepare chart err(%v)", err)
	}
	chart.PrintChartTable()

	return &astroV1.UserUpdateResponse{}, nil

}

func (s *grpcAstroServer) Get(ctx context.Context, in *astroV1.UserGetRequest) (*astroV1.UserGetResponse, error) {
	logrus.Info("Received: Get request ", in)
	astro := &astroV1.AstroObject{}

	mongo, err := mongo.NewDSMongoSingleton()
	if err != nil {
		logrus.Error("Cannot connect to DB")
		return nil, err
	}

	ownerUid, err := helpers.GetUserUuidFromContext(ctx)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"metadata.uid", in.Uuid}, {"metadata.owner", ownerUid}}
	err = mongo.FindOne(ASTRO_BUCKET_NAME, filter, astro)
	if err != nil {
		logrus.Error("Cannot Get Item ", err)
		return nil, err
	}

	return &astroV1.UserGetResponse{
		Astro: astro,
	}, nil

}

// func (s *grpcAstroServer) Update(ctx context.Context, in *protoV1.ClusterUpdateRequest) (*protoV1.ClusterUpdateResponse, error) {

// 	logrus.Info("Received Update: ", in)
// 	// insert cluster into DB
// 	mongo, err := mongo.NewDSMongoSingleton()
// 	if err != nil {
// 		logrus.Error("Cannot connect to DB")
// 		return nil, err
// 	}

// 	filter := bson.D{{"metadata.uid", in.Cluster.MetaData.Uid}}
// 	update := bson.D{{"$set", bson.D{
// 		{"info.region", in.Cluster.Info.Region},
// 		{"status.status", in.Cluster.Status.Status},
// 		{"status.statusdescription", in.Cluster.Status.StatusDescription},
// 		{"status.kubever", in.Cluster.Status.KubeVer},
// 		{"status.nodesup", in.Cluster.Status.NodesUp},
// 		{"status.nodesdown", in.Cluster.Status.NodesDown},
// 		{"status.podsup", in.Cluster.Status.PodsUp},
// 		{"status.podsdown", in.Cluster.Status.PodsDown},
// 		{"status.volumesup", in.Cluster.Status.VolumesUp},
// 		{"status.volumesdown", in.Cluster.Status.VolumesDown},
// 		{"status.eventsnormal", in.Cluster.Status.EventsNormal},
// 		{"status.eventswarning", in.Cluster.Status.EventsWarning},
// 		{"status.events", in.Cluster.Status.Events},
// 		{"status.updatedat", timestamppb.New(time.Now())},
// 	}}}

// 	err = mongo.UpdateOne(CLUSTER_BUCKET_NAME, filter, update, true)
// 	if err != nil {
// 		logrus.Error("Cannot Update collection entry")
// 		return nil, err
// 	}

// 	return &protoV1.ClusterUpdateResponse{}, nil

// }
