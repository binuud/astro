package astroServer

import (
	"context"
	"time"

	protoV1 "github.com/dronasys-com/astro/proto/v1/astrocalc/go"
	"github.com/dronasys-com/trance-microservices/pkg/dataLayer/mongo"
	"github.com/dronasys-com/trance-microservices/pkg/foundation/helpers"
	"github.com/dronasys-com/trance-microservices/pkg/servers/authServer"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	CLUSTER_BUCKET_NAME = "astro"
)

type grpcAstroServer struct {
	protoV1.UnimplementedAstroServer
}

func NewGrpcAstroServer() *grpcAstroServer {
	return &grpcAstroServer{}
}

func (s *grpcAstroServer) Update(ctx context.Context, in *protoV1.User) (*protoV1.ClusterCreateResponse, error) {

	var err error
	in.Cluster.MetaData.Uid = uuid.New().String()
	in.Cluster.MetaData.Owner, err = authServer.GetUserUuidFromContext(ctx)
	if err != nil {
		return nil, err
	}

	nowTime := timestamppb.New(time.Now())
	in.Cluster.MetaData.CreatedAt = nowTime
	in.Cluster.MetaData.UpdatedAt = nowTime

	logrus.Info("Received Create: ", in)

	newCluster := &protoV1.ClusterSecretObject{
		MetaData: in.Cluster.MetaData,
		Info:     in.Cluster.Info,
		Status: &protoV1.ClusterStatus{
			Events:    nil,
			UpdatedAt: timestamppb.New(time.Now()),
			Status:    protoV1.ClusterStatus_Status(protoV1.AppProvisionStatus_ProvisionRequest),
		},
		Secret: &protoV1.ClusterSecret{
			HandShake:   uuid.New().String() + helpers.GenerateRandomStringURLSafe(32),
			TokenSecret: helpers.GenerateRandomStringURLSafe(32),
		},
	}

	// insert cluster into DB
	mongo, err := mongo.NewDSMongoSingleton()

	if err != nil {
		logrus.Error("Cannot connect to DB")
		return nil, err
	}
	err = mongo.InsertOne(CLUSTER_BUCKET_NAME, newCluster)
	if err != nil {
		logrus.Error("Cannot Create cluster entry")
		return nil, err
	}

	cluster, err := s.GetFromDB(newCluster.MetaData.Uid)
	if err != nil {
		logrus.Error("Cannot read new entry entry")
		return nil, err
	}

	return &protoV1.ClusterCreateResponse{
		Cluster:   cluster,
		HandShake: newCluster.Secret.HandShake,
	}, nil

}

func (s *grpcAstroServer) Get(ctx context.Context, in *protoV1.ClusterGetRequest) (*protoV1.ClusterGetResponse, error) {
	logrus.Info("Received: Get request ", in)
	cluster := &protoV1.ClusterObject{}

	mongo, err := mongo.NewDSMongoSingleton()
	if err != nil {
		logrus.Error("Cannot connect to DB")
		return nil, err
	}

	ownerUid, err := helpers.GetUserUuidFromContext(ctx)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"metadata.uid", in.ClusterUid}, {"metadata.owner", ownerUid}}
	err = mongo.FindOne(CLUSTER_BUCKET_NAME, filter, cluster)
	if err != nil {
		logrus.Error("Cannot Get Item ", err)
		return nil, err
	}

	return &protoV1.ClusterGetResponse{
		Cluster: cluster,
	}, nil

}

func (s *grpcAstroServer) GetFromDB(uid string) (*protoV1.ClusterObject, error) {

	mongo, err := mongo.NewDSMongoSingleton()
	if err != nil {
		logrus.Error("Cannot connect to DB")
		return nil, err
	}
	cluster := &protoV1.ClusterObject{}

	filter := bson.D{{"metadata.uid", uid}}
	err = mongo.FindOne(CLUSTER_BUCKET_NAME, filter, cluster)
	if err != nil {
		logrus.Error("Cannot Get Item ", err)
		return nil, err
	}

	return cluster, nil

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
