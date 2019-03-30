package mongoservice

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"log"

	"github.com/JustSomeHack/one-oauth2-server/models"
	"github.com/JustSomeHack/one-oauth2-server/models/mongoconfig"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const usersCollection = "Users"

// MongoService interface for Users database.Collection(usersCollection)
type MongoService interface {
	Authenticate(username string, password string) (*models.User, error)
}

type mongoService struct {
	database *mongo.Database
}

// NewMongoService returns a new instance of MongoService
func NewMongoService(config *mongoconfig.MongoConfig) MongoService {
	var connectionString string
	if config.DBUser != "" && config.DBPass != "" {
		connectionString = fmt.Sprintf("mongodb://%s:%s@%s/admin", config.DBUser, config.DBPass, config.MongoURL)
	} else {
		connectionString = fmt.Sprintf("mongodb://%s/%s", config.MongoURL, config.DBName)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &mongoService{
		database: client.Database(config.DBName),
	}
}

// GetOne a document from database.Collection(usersCollection)
func (m *mongoService) Authenticate(username string, password string) (*models.User, error) {
	filter := bson.M{}
	filter["Username"] = username

	h := md5.New()
	io.WriteString(h, password)
	hashed := fmt.Sprintf("%x", h.Sum(nil))

	filter["Password"] = hashed

	opts := make([]*options.FindOneOptions, 0)

	result := m.database.Collection(usersCollection).FindOne(context.Background(), filter, opts...)
	elem := new(models.User)
	err := result.Decode(elem)
	if err != nil {
		return nil, err
	}
	return elem, nil
}
