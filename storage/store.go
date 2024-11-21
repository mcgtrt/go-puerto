package storage

import (
	mongo_store "github.com/mcgtrt/go-puerto/storage/mongo"
	postgres_store "github.com/mcgtrt/go-puerto/storage/postgres"
	valkey_store "github.com/mcgtrt/go-puerto/storage/valkey"
	"github.com/mcgtrt/go-puerto/utils"
)

// Holds all instances of set up databases to ease accessing data
// accross the entire project.
type Store struct {
	Mongo    *mongo_store.MongoStore
	Postgres *postgres_store.PostgresStore
	Valkey   *valkey_store.ValkeyStore
}

// Create new store based on the configuration provided
func NewStore(config *utils.Config) (*Store, error) {
	var (
		mongo    *mongo_store.MongoStore
		postgres *postgres_store.PostgresStore
		valkey   *valkey_store.ValkeyStore
	)
	if config.Mongo != nil {
		client, err := config.Mongo.Client()
		if err != nil {
			return nil, err
		}
		mongo = mongo_store.NewMongoStore(client, config.Mongo.DBName)
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
