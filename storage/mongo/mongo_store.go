package mongo_store

import "go.mongodb.org/mongo-driver/mongo"

type MongoStore struct {
	Client *mongo.Client
	DBName string
}

func NewMongoStore(client *mongo.Client, dbname string) *MongoStore {
	return &MongoStore{
		Client: client,
		DBName: dbname,
	}
}