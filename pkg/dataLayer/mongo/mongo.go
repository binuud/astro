package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dronasys-com/astro/pkg/foundation/config"
	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	mongolib "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DataLayer_Mongo *MongoConnection

type Pagination struct {
	CurPage int
	Total   int
	PerPage int
}

type MongoConnection struct {
	Client     *mongo.Client
	Endpoint   string
	URI        string
	DBName     string
	DBConnect  *mongo.Database
	Collection string
}

func NewDSMongoSingleton() (*MongoConnection, error) {

	// if singleton exists return same
	if DataLayer_Mongo != nil {
		logrus.Trace("Singleton exists, returned existing Mongo DB Connection")
		return DataLayer_Mongo, nil
	}

	// Create new connection, if called for the first time

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uri := ""
	if config.Config.MongoUser != "" {
		uri = fmt.Sprintf(`mongodb://%s:%s@%s/`,
			config.Config.MongoUser,
			config.Config.MongoPassword,
			config.Config.MongoEndpoint,
		)
	} else {
		uri = fmt.Sprintf(`mongodb://%s/`,
			config.Config.MongoEndpoint,
		)
	}

	logrus.Infof("Datalayer connecting to - from config (%s)", config.Config.MongoEndpoint)
	logrus.Infof("Datalayer connecting to (%s)", uri)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongolib.Connect(ctx, clientOptions)

	if err != nil {
		logrus.Fatal("DATALAYER: mongo client couldn't connect with background context: ", err)
		return nil, fmt.Errorf("todo: mongo client couldn't connect with background context: %v", err)
	}

	logrus.Info("DATALAYER: New Mongo Singleton Created ")

	DataLayer_Mongo = &MongoConnection{
		Client:   client,
		URI:      uri,
		Endpoint: config.Config.MongoEndpoint,
		DBName:   config.Config.MongoDB,
	}

	DataLayer_Mongo.Ping()

	return DataLayer_Mongo, nil

}

func (m *MongoConnection) SetCollection(collection string) {
	m.Collection = collection
	logrus.Trace("Collection set to ", m.Collection)
}

func (m *MongoConnection) Ping() {

	// filter := bson.D{{}}
	// list, err := m.Client.ListDatabaseNames(context.TODO(), filter)
	// if err != nil {
	// 	logrus.Error("Cannot connect to DB ", err)
	// }
	// logrus.Trace("DB List", list)

	// Check the connection
	err := m.Client.Ping(context.TODO(), nil)
	if err != nil {
		logrus.Fatal(err)
	}
}

func (m *MongoConnection) GetCollectionInterface(collectionName string) *mongo.Collection {
	return m.Client.Database(m.DBName).Collection(collectionName)
}

func (m *MongoConnection) UpdateOne(collectionName string, query interface{}, doc interface{}, upsert bool) error {

	collection := m.Client.Database(m.DBName).Collection(collectionName)
	opts := options.Update()
	opts.SetUpsert(upsert)
	_, err := collection.UpdateOne(context.TODO(), query, doc, opts)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (m *MongoConnection) InsertMany(collectionName string, docs ...interface{}) error {

	coll := m.Client.Database(m.DBName).Collection(collectionName)
	opts := options.InsertMany()
	_, err := coll.InsertMany(context.TODO(), docs, opts)
	if err != nil {
		logrus.Fatal("Cannot insert docs into database ", err)
	}
	return err
}

func (m *MongoConnection) InsertOne(collectionName string, doc interface{}) error {

	coll := m.Client.Database(m.DBName).Collection(collectionName)
	opts := options.InsertOne()
	_, err := coll.InsertOne(context.TODO(), doc, opts)
	if err != nil {
		logrus.Fatal("Cannot insert into database ", err)
	}
	return err
}

func (m *MongoConnection) Find(collectionName string, query interface{}, sortType interface{}, currentPage int64, page_size int64) (*mongo.Cursor, error) {
	// Pass these options to the Find method

	logrus.Info(" Database ", m.DBName, " Collection ", collectionName)
	findOptions := options.Find()
	if currentPage > 1 {
		findOptions.SetSkip((currentPage - 1) * page_size)
	}

	findOptions.SetLimit(page_size)
	findOptions.SetSort(sortType)

	collection := m.Client.Database(m.DBName).Collection(collectionName)
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), query, findOptions)
	if err != nil {
		return nil, err
	}

	return cur, nil
}

func (m *MongoConnection) FindWithProjection(collectionName string, query interface{}, projection interface{}, sortType interface{}, pagination *astroV1.Pagination) (*mongo.Cursor, error) {
	// Pass these options to the Find method

	if pagination.GetPerPage() > 50 {
		pagination.PerPage = 50
	}

	logrus.Info(" Database ", m.DBName, " Collection ", collectionName)
	findOptions := options.Find()
	if pagination.GetCurrent() > 1 {
		findOptions.SetSkip((pagination.GetCurrent() - 1) * pagination.GetPerPage())
	}

	findOptions.SetLimit(pagination.GetPerPage())
	findOptions.SetSort(sortType)
	findOptions.SetProjection(projection)

	collection := m.Client.Database(m.DBName).Collection(collectionName)
	// total, err := collection.CountDocuments(context.Background(), query)
	// if err == nil {
	// 	pagination.Total = int64(total)
	// }
	pagination.Total = 111147

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), query, findOptions)
	if err != nil {
		return nil, err
	}

	return cur, nil
}

func (m *MongoConnection) FindOne(collectionName string, query interface{}, result interface{}) error {

	findOptions := options.FindOne()

	collection := m.Client.Database(m.DBName).Collection(collectionName)

	err := collection.FindOne(context.TODO(), query, findOptions).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoConnection) DeleteOne(collectionName string, query interface{}, result interface{}) error {

	deleteOptions := options.Delete()

	collection := m.Client.Database(m.DBName).Collection(collectionName)

	res, err := collection.DeleteOne(context.TODO(), query, deleteOptions)
	if err != nil {
		return err
	}
	logrus.Info("Object deleted (%v) (%v)", query, res)

	return nil
}

// func (m MongoConnection) ConvertToBson(in interface{}) (interface{}, error) {
// 	marshaler := jsonpb.Marshaler{}
// 	var b bytes.Buffer
// 	err := marshaler.Marshal(&b, in)
// 	if err != nil {
// 		logrus.Errorf("Error in marshalling %v", err)
// 		return nil, err
// 	}

// 	var v interface{}
// 	if err = json.Unmarshal(b.Bytes(), &v); err != nil {
// 		logrus.Errorf("Error in un marshalling  %v", err)
// 		return nil, err
// 	}
// pByte, err := bson.Marshal(clusterObj)
// if err != nil {
// 	logrus.Error("Cannot Marshal to clusterObj", err)
// 	return err
// }
//}
// var updateBson bson.M
// err = bson.Unmarshal(pByte, &updateBson)
// if err != nil {
// 	logrus.Error("Cannot UnMarshal to Bson", err)
// 	return err
//
// 	return v, err
// }
