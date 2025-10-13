package mongo

import (
	"context"
	"fmt"
	"log"
	"os"

	astroV1 "github.com/dronasys-com/astro/gen/go/v1/astro"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DataLayer_Mongo *MongoConnection

type MongoConfig struct {
	Endpoint      string /* Endpoint of the Mongo DB, this is a minimum requirement */
	URI           string
	DBName        string
	User          string
	Password      string
	EncryptionKey string
}

type Pagination struct {
	CurPage int
	Total   int
	PerPage int
}

type MongoConnection struct {
	Client     *mongo.Client
	Config     *MongoConfig
	DBConnect  *mongo.Database
	Collection string
}

func NewDSMongoSingleton() (*MongoConnection, error) {

	// if singleton exists return same
	if DataLayer_Mongo != nil {
		logrus.Trace("Singleton exists, returned existing Mongo DB Connection")
		return DataLayer_Mongo, nil
	}

	logrus.Info("Initializing Mongo Connection")
	mongoConfig := initMongoConfig()

	// Create new connection, if called for the first time
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// logrus.Infof("Datalayer connecting to - from config (%s)", mongoConfig.Endpoint)
	logrus.Infof("Datalayer connecting to (%s)", mongoConfig.URI)
	clientOptions := options.Client().ApplyURI(mongoConfig.URI)
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		logrus.Fatalf("DATALAYER: mongo client couldn't connect with background context: %v", err)
		return nil, fmt.Errorf("todo: mongo client couldn't connect with background context: %v", err)
	}

	logrus.Info("DATALAYER: New Mongo Singleton Created ")

	DataLayer_Mongo = &MongoConnection{
		Client: client,
		Config: mongoConfig,
	}

	DataLayer_Mongo.Ping()

	return DataLayer_Mongo, nil

}

/* Check whether the required env variables are set */
func initMongoConfig() *MongoConfig {

	config := &MongoConfig{
		Endpoint:      os.Getenv("MONGO_ENDPOINT"),
		URI:           "",
		DBName:        os.Getenv("MONGO_DB"),
		User:          os.Getenv("MONGO_USER"),
		Password:      os.Getenv("MONGO_PASSWORD"),
		EncryptionKey: os.Getenv("MONGO_ENCRYPTION_KEY"),
	}

	if config.Endpoint == "" {
		logrus.Panic("Env variables MONGO_ENDPOINT/MONGO_DB/MONGO_USER/MONGO_PASSWORD/MONGO_ENCRYPTION_KEY not set")
	}

	if config.User != "" {
		config.URI = fmt.Sprintf(`mongodb://%s:%s@%s/`,
			config.User,
			config.Password,
			config.Endpoint,
		)
	} else {
		config.URI = fmt.Sprintf(`mongodb://%s/`,
			config.Endpoint,
		)
	}

	return config
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
	return m.Client.Database(m.Config.DBName).Collection(collectionName)
}

func (m *MongoConnection) UpdateOne(collectionName string, query interface{}, doc interface{}, upsert bool) error {

	collection := m.Client.Database(m.Config.DBName).Collection(collectionName)
	opts := options.UpdateOne()
	opts.SetUpsert(upsert)
	_, err := collection.UpdateOne(context.TODO(), query, doc, opts)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (m *MongoConnection) InsertMany(collectionName string, docs ...interface{}) error {

	coll := m.Client.Database(m.Config.DBName).Collection(collectionName)
	opts := options.InsertMany()
	_, err := coll.InsertMany(context.TODO(), docs, opts)
	if err != nil {
		logrus.Fatal("Cannot insert docs into database ", err)
	}
	return err
}

func (m *MongoConnection) InsertOne(collectionName string, doc interface{}) error {

	coll := m.Client.Database(m.Config.DBName).Collection(collectionName)
	opts := options.InsertOne()
	_, err := coll.InsertOne(context.TODO(), doc, opts)
	if err != nil {
		logrus.Fatal("Cannot insert into database ", err)
	}
	return err
}

func (m *MongoConnection) Find(collectionName string, query interface{}, sortType interface{}, currentPage int64, page_size int64) (*mongo.Cursor, error) {
	// Pass these options to the Find method

	logrus.Info(" Database ", m.Config.DBName, " Collection ", collectionName)
	findOptions := options.Find()
	if currentPage > 1 {
		findOptions.SetSkip((currentPage - 1) * page_size)
	}

	findOptions.SetLimit(page_size)
	findOptions.SetSort(sortType)

	collection := m.Client.Database(m.Config.DBName).Collection(collectionName)
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

	logrus.Info(" Database ", m.Config.DBName, " Collection ", collectionName)
	findOptions := options.Find()
	if pagination.GetCurrent() > 1 {
		findOptions.SetSkip((pagination.GetCurrent() - 1) * pagination.GetPerPage())
	}

	findOptions.SetLimit(pagination.GetPerPage())
	findOptions.SetSort(sortType)
	findOptions.SetProjection(projection)

	collection := m.Client.Database(m.Config.DBName).Collection(collectionName)
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

	collection := m.Client.Database(m.Config.DBName).Collection(collectionName)

	err := collection.FindOne(context.TODO(), query, findOptions).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoConnection) DeleteOne(collectionName string, query interface{}, result interface{}) error {

	deleteOptions := options.DeleteOne()

	collection := m.Client.Database(m.Config.DBName).Collection(collectionName)

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
