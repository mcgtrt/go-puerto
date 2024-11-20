package storage

import (
	"errors"

	mongo_store "github.com/mcgtrt/go-puerto/storage/mongo"
	postgres_store "github.com/mcgtrt/go-puerto/storage/postgres"
	valkey_store "github.com/mcgtrt/go-puerto/storage/valkey"
	"github.com/mcgtrt/go-puerto/utils"
)

// General store structure holding all database controllers
type Store struct {
	Mongo    *mongo_store.MongoStore
	Postgres *postgres_store.PostgresStore
	Valkey   *valkey_store.ValkeyStore
}

// Create new store from local config
func NewStore(config *utils.Config) (*Store, error) {
	var (
		mongo    *mongo_store.MongoStore
		postgres *postgres_store.PostgresStore
		valkey   *valkey_store.ValkeyStore
	)
	if config.Mongo != nil {
		if config.Mongo.Client == nil {
			return nil, errors.New("mongo client cannot be empty")
		}
		if config.Mongo.DBName == "" {
			return nil, errors.New("mongo database name cannot be empty")
		}
		mongo = mongo_store.NewMongoStore(config.Mongo.Client, config.Mongo.DBName)
	}
	if config.Postgres != nil {
		postgres = postgres_store.NewPostgresStore()
	}
	if config.Valkey != nil {
		valkey = valkey_store.NewValkeyStore()
	}

	return &Store{
		Mongo:    mongo,
		Postgres: postgres,
		Valkey:   valkey,
	}, nil
}
