package astroServer

import (
	"context"
	"time"

	astroV1 "github.com/dronasys-com/astro/gen/go/v1/astro"
	"github.com/dronasys-com/astro/pkg/astro"
	"github.com/dronasys-com/astro/pkg/dataLayer/mongo"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateRecord(in *astroV1.UserCreateRequest) (*astroV1.AstroObject, error) {

	metaData := &astroV1.Metadata{
		Uid:  uuid.New().String(),
		Name: in.Metadata.Name,
		// Owner, err = authServer.GetUserUuidFromContext(ctx)
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
		ChartType: astroV1.Metadata_Global,
	}

	// logrus.Info("Received Create: ", in)

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
	// chart.PrintChartTable()

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
		Dashas: chart.Dashas,
	}

	err = mongo.InsertOne(ASTRO_BUCKET_NAME, astroObject)
	if err != nil {
		logrus.Error("Cannot Create Chart entry")
		return nil, err
	}

	return astroObject, nil
}

func UpdateRecord(in *astroV1.UserUpdateRequest, filter bson.D) (*astroV1.AstroObject, error) {

	// insert new chart in mongo
	mongo, err := mongo.NewDSMongoSingleton()

	if err != nil {
		logrus.Error("Cannot connect to DB")
		return nil, err
	}

	currentAstroObject := &astroV1.AstroObject{}
	err = mongo.FindOne(ASTRO_BUCKET_NAME, filter, currentAstroObject)
	if err != nil {
		logrus.Error("Cannot Get Item ", err)
		return nil, err
	}

	chart, err := astro.NewAstroChart(in.BirthData)
	if err != nil {
		return nil, err
	}

	astroObject := chart.AstroObject
	astroObject.Metadata = currentAstroObject.Metadata

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
	return &astroV1.AstroObject{
		BirthData: chart.BirthData,
		Planets: &astroV1.Planets{
			PlanetsMap: chart.Planets.PlanetsMap,
		},
	}, nil
}

func UpSertRecord(in *astroV1.UserCreateRequest, filter bson.D) (*astroV1.AstroObject, error) {

	// insert new chart in mongo
	mongo, err := mongo.NewDSMongoSingleton()

	if err != nil {
		logrus.Error("Cannot connect to DB")
		return nil, err
	}

	metaData := &astroV1.Metadata{
		Uid:  uuid.New().String(),
		Name: in.Metadata.Name,
		// Owner, err = authServer.GetUserUuidFromContext(ctx)
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
		ChartType: astroV1.Metadata_Global,
	}

	chart, err := astro.NewAstroChart(in.BirthData)
	if err != nil {
		return nil, err
	}

	astroObject := chart.AstroObject
	astroObject.Metadata = metaData

	updateFields := &astroV1.AstroObject{}
	conv, _ := bson.Marshal(chart.AstroObject)
	bson.Unmarshal(conv, updateFields)
	update := bson.M{
		"$set": updateFields,
	}

	err = mongo.UpdateOne(ASTRO_BUCKET_NAME, filter, update, true)
	if err != nil {
		logrus.Error("Cannot Update collection entry")
		return nil, err
	}
	return astroObject, nil
}

func BulkInsert(astroObjects []interface{}) error {
	// insert new chart in mongo
	mongo, err := mongo.NewDSMongoSingleton()

	if err != nil {
		logrus.Error("Cannot connect to DB")
		return err
	}
	return mongo.InsertMany(ASTRO_BUCKET_NAME, astroObjects...)

}

func CreateAstroObjectForTime(name string, t time.Time) (*astroV1.AstroObject, error) {

	birthData := CreateBirthDataForTime(t)
	return CreateAstroObject(name, birthData)

}

func CreateAstroObject(name string, birthData *astroV1.BirthData) (*astroV1.AstroObject, error) {
	metaData := &astroV1.Metadata{
		Uid:  uuid.New().String(),
		Name: name,
		// Owner, err = authServer.GetUserUuidFromContext(ctx)
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
		ChartType: astroV1.Metadata_Global,
	}

	chart, err := astro.NewAstroChart(birthData)
	if err != nil {
		return nil, err
	}

	astroObject := chart.AstroObject
	astroObject.Metadata = metaData

	return astroObject, nil
}

func CreateBirthDataForTime(t time.Time) *astroV1.BirthData {
	return &astroV1.BirthData{
		Tob: &astroV1.TobData{
			D:    int32(t.UTC().Day()),
			M:    int32(t.UTC().Month()),
			Y:    int32(t.UTC().Year()),
			Hour: int32(t.UTC().Hour()),
			Min:  int32(t.UTC().Minute()),
			Sec:  0,
			Zone: "UTC",
		},
		Pob: &astroV1.PobData{
			City: "Chennai",

			Lat:  13.079999923706055,
			Long: 80.69999694824219,
		},
	}
}

func GetRecord(filter bson.D) (*astroV1.AstroObject, error) {
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

	astroObject := &astroV1.AstroObject{}
	err = mongo.FindOne(ASTRO_BUCKET_NAME, filter, astroObject)
	if err != nil {
		return nil, err
	}
	return astroObject, nil
}

func ListRecords(filter bson.D, projection bson.D, pagination *astroV1.Pagination) ([]*astroV1.AstroObject, error) {
	charts := []*astroV1.AstroObject{}

	mongo, err := mongo.NewDSMongoSingleton()
	if err != nil {
		logrus.Error("Cannot connect to DB")
		return nil, err
	}

	sortType := bson.M{}

	cursor, err := mongo.FindWithProjection(ASTRO_BUCKET_NAME, filter, projection, sortType, pagination)
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

	return charts, nil
}
