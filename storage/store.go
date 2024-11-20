package storage

import (
	"errors"

	mongo_store "github.com/mcgtrt/go-puerto/storage/mongo"
	postgres_store "github.com/mcgtrt/go-puerto/storage/postgres"
	valkey_store "github.com/mcgtrt/go-puerto/storage/valkey"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	Mongo    *mongo_store.MongoStore
	Postgres *postgres_store.PostgresStore
	Valkey   *valkey_store.ValkeyStore
}

func NewStore(config *Config) (*Store, error) {
	var (
		mongo    *mongo_store.MongoStore
		postgres *postgres_store.PostgresStore
		valkey   *valkey_store.ValkeyStore
	)
	if config.WithMongo {
		if config.MongoClient == nil {
			return nil, errors.New("mongo client cannot be empty")
		}
		if config.MongoDBName == "" {
			return nil, errors.New("mongo database name cannot be empty")
		}
		mongo = mongo_store.NewMongoStore(config.MongoClient, config.MongoDBName)
	}
	if config.WithPostgres {
		postgres = postgres_store.NewPostgresStore()
	}
	if config.WithValkey {
		valkey = valkey_store.NewValkeyStore()
	}

	return &Store{
		Mongo:    mongo,
		Postgres: postgres,
		Valkey:   valkey,
	}, nil
}

type Config struct {
	MongoClient *mongo.Client
	MongoDBName string

	WithMongo    bool
	WithPostgres bool
	WithValkey   bool
}
