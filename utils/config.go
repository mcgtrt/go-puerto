package utils

import (
	"context"
	"errors"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Holds global configuration sourced from local .env file
type Config struct {
	HTTP     *HTTPConfig
	Mongo    *MongoConfig
	Postgres *PostgresConfig
	Valkey   *ValkeyConfig
}

// Create new default config from the local .env file. If any part of the configuration
// will cause an error, will return an empty config (nil) and error message
func NewDefaultConfig() (*Config, error) {
	config := &Config{}

	http, err := newDefaultHTTPConfig()
	if err != nil {
		return nil, err
	}
	config.HTTP = http

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

// Configuration required for HTTP server
type HTTPConfig struct {
	FileServerPath string
	ImportAlpineJS bool
	Port           int
}

func newDefaultHTTPConfig() (*HTTPConfig, error) {
	config := &HTTPConfig{}
	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return nil, errors.New("http port must be a valid port number")
	}
	if port < 1000 || port > 65535 {
		return nil, errors.New("only registered and dynamic ports are allowed (1000 - 65535)")
	}
	config.Port = port
	if os.Getenv("IMPORT_ALPINE_JS") == "true" {
		config.ImportAlpineJS = true
	}
	path := os.Getenv("FILE_SERVER_PATH")
	if !IsURLSafe(path) {
		return nil, errors.New("file server path is not URL safe")
	}
	config.FileServerPath = path
	return config, nil
}

// Required configuration for creating mongodb connection
type MongoConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (c *MongoConfig) ConnectionString() string {
	if c.Username == "" || c.Password == "" || c.Host == "" || c.Port == "" {
		return ""
	}
	return "mongodb://" + c.Username + ":" + c.Password + "@" + c.Host + ":" + c.Port
}

func (c *MongoConfig) Client() (*mongo.Client, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.ConnectionString()))
	if err != nil {
		return nil, errors.New("error creating mongo client: " + err.Error())
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, errors.New("error connecting to mongodb - could not ping the database: " + err.Error())
	}
	return client, nil
}

func newDefaultMongoConfig() (*MongoConfig, error) {
	dbname := os.Getenv("MONGO_DB_NAME")
	if dbname == "" {
		return nil, errors.New("mongo database name cannot be empty")
	}
	username := os.Getenv("MONGO_USERNAME")
	if username == "" {
		return nil, errors.New("mongo username cannot be empty")
	}
	password := os.Getenv("MONGO_PASSWORD")
	if password == "" {
		return nil, errors.New("mongo password cannot be empty")
	}
	host := os.Getenv("MONGO_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("MONGO_PORT")
	if port == "" {
		port = "27017"
	}
	if _, err := strconv.Atoi(port); err != nil {
		return nil, errors.New("invalid port for mongo connection")
	}
	return &MongoConfig{
		DBName:   dbname,
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
	}, nil
}

// Required configuration for creating postgres connection
type PostgresConfig struct{}

func newDefaultPostgresConfig() (*PostgresConfig, error) {
	return &PostgresConfig{}, nil
}

// Required configuration for creating valkey connection
type ValkeyConfig struct{}

func newDefaultValkeyConfig() (*ValkeyConfig, error) {
	return &ValkeyConfig{}, nil
}
