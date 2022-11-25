package astroServer

import (
	"context"
	"time"

	"github.com/dronasys-com/astro/pkg/astro"
	"github.com/dronasys-com/astro/pkg/dataLayer/mongo"
	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	metaData := &astroV1.Metadata{
		Uid:  uuid.New().String(),
		Name: in.Metadata.Name,
		// Owner, err = authServer.GetUserUuidFromContext(ctx)
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
		ChartType: astroV1.Metadata_Global,
	}

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

	// insert new chart in mongo
	mongo, err := mongo.NewDSMongoSingleton()

	if err != nil {
		logrus.Error("Cannot connect to DB")
		return nil, err
	}

	astroObject := &astroV1.AstroObject{
		Metadata:  metaData,
		BirthData: chart.BirthData,
		Planets: &astroV1.Planets{
			PlanetsMap: chart.Planets.PlanetsMap,
		},
		Dashas: chart.DashaBhukti,
	}

	err = mongo.InsertOne(ASTRO_BUCKET_NAME, astroObject)
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

	// insert new chart in mongo
	mongo, err := mongo.NewDSMongoSingleton()

	if err != nil {
		logrus.Error("Cannot connect to DB")
		return nil, err
	}

	astroObject := &astroV1.AstroObject{}
	err = mongo.FindOne(ASTRO_BUCKET_NAME, filter, astroObject)
	if err != nil {
		logrus.Error("Cannot Get Item ", err)
		return nil, err
	}

	// update := bson.D{{"$set", bson.D{
	// 	{"birthdata.tob.d", in.BirthData.Tob.D},
	// 	{"birthdata.tob.m", in.BirthData.Tob.M},
	// 	{"birthdata.tob.y", in.BirthData.Tob.Y},
	// 	{"birthdata.tob.hour", in.BirthData.Tob.Hour},
	// 	{"birthdata.tob.min", in.BirthData.Tob.Min},
	// 	{"birthdata.tob.sec", in.BirthData.Tob.Sec},
	// 	{"birthdata.tob.zone", in.BirthData.Tob.Zone},
	// 	{"status.podsdown", in.Cluster.Status.PodsDown},
	// 	{"status.volumesup", in.Cluster.Status.VolumesUp},
	// 	{"status.volumesdown", in.Cluster.Status.VolumesDown},
	// 	{"status.eventsnormal", in.Cluster.Status.EventsNormal},
	// 	{"status.eventswarning", in.Cluster.Status.EventsWarning},
	// 	{"status.events", in.Cluster.Status.Events},
	// 	{"status.updatedat", timestamppb.New(time.Now())},
	// }}}

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

	astroObject.BirthData = in.BirthData
	astroObject.Planets = chart.Planets
	astroObject.Dashas = chart.DashaBhukti

	updateFields := &astroV1.AstroObject{}
	conv, _ := bson.Marshal(astroObject)
	bson.Unmarshal(conv, updateFields)
	update := bson.M{
		"$set": updateFields,
	}

	err = mongo.UpdateOne(ASTRO_BUCKET_NAME, filter, update, true)
	if err != nil {
		logrus.Error("Cannot Update collection entry")
		return nil, err
	}

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

	// db654581-5870-4cb7-bfa8-bcd239499b2e
	mongo, err := mongo.NewDSMongoSingleton()
	if err != nil {
		logrus.Error("Cannot connect to DB")
		return nil, err
	}

	// ownerUid, err := helpers.GetUserUuidFromContext(ctx)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"metadata.uid", in.Uuid}}

	astroObject := &astroV1.AstroObject{}
	err = mongo.FindOne(ASTRO_BUCKET_NAME, filter, astroObject)
	if err != nil {
		logrus.Error("Cannot Get Item ", err)
		return nil, err
	}

	// chart, err := astro.NewAstroChart(&astroV1.BirthData{
	// 	Tob: &astroV1.TobData{
	// 		D:    04,
	// 		M:    02,
	// 		Y:    1978,
	// 		Hour: 10,
	// 		Min:  44,
	// 		Sec:  0,
	// 		Zone: "Asia/Kolkata",
	// 	},
	// 	Pob: &astroV1.PobData{
	// 		City: "Chennai",
	// 		Lat:  13.08,
	// 		Long: 80.7,
	// 	},
	// })
	// if err != nil {
	// 	logrus.Errorf("Cannot prepare chart err(%v)", err)
	// }
	// chart.PrintChartTable()

	return &astroV1.UserGetResponse{
		Astro: astroObject,
	}, nil

}

func (s *grpcAstroServer) List(ctx context.Context, in *astroV1.UserListRequest) (*astroV1.UserListResponse, error) {

	logrus.Info("Received: Listing request")
	charts := []*astroV1.AstroObject{}

	mongo, err := mongo.NewDSMongoSingleton()
	if err != nil {
		logrus.Error("Cannot connect to DB")
		return nil, err
	}

	// ownerUid, err := helpers.GetUserUuidFromContext(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	//filter := bson.D{{"metadata.owner", ownerUid}}
	filter := bson.D{{}}
	sortType := bson.M{}
	projection := bson.D{
		{"metadata", 1},
	}
	cursor, err := mongo.FindWithProjection(ASTRO_BUCKET_NAME, filter, projection, sortType, 1, 1000)
	if err != nil {
		logrus.Error("Cannot List collection ", err)
		return nil, err
	}

	//Map result to slice
	for cursor.Next(context.TODO()) {
		t := astroV1.AstroObject{}
		err := cursor.Decode(&t)
		if err != nil {
			return nil, err
		}
		charts = append(charts, &t)
	}
	// once exhausted, close the cursor
	cursor.Close(context.TODO())

	return &astroV1.UserListResponse{
		Astro: charts,
	}, nil

}
