package utils

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
)

// Holds global configuration sourced from local .env file
type Config struct {
	Mongo    *MongoConfig
	Postgres *PostgresConfig
	Valkey   *ValkeyConfig
}

func NewDefaultConfig() (*Config, error) {
	config := &Config{}

	if os.Getenv("INCLUDE_MONGO") == "true" {
		mongo, err := newDefaultMongoConfig()
		if err != nil {
			return nil, err
		}
		config.Mongo = mongo
	}
	if os.Getenv("INCLUDE_POSTGRES") == "true" {
		postgres, err := newDefaultPostgresConfig()
		if err != nil {
			return nil, err
		}
		config.Postgres = postgres
	}
	if os.Getenv("INCLUDE_VALKEY") == "true" {
		valkey, err := newDefaultValkeyConfig()
		if err != nil {
			return nil, err
		}
		config.Valkey = valkey
	}

	return config, nil
}

type MongoConfig struct {
	Client *mongo.Client
	DBName string
}

func newDefaultMongoConfig() (*MongoConfig, error) {
	return nil, nil
}

type PostgresConfig struct{}

func newDefaultPostgresConfig() (*PostgresConfig, error) {
	return nil, nil
}

type ValkeyConfig struct{}

func newDefaultValkeyConfig() (*ValkeyConfig, error) {
	return nil, nil
}
