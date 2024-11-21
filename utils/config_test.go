package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	var (
		fileserverPath = "FILE_SERVER_PATH"
		httpPort       = "HTTP_PORT"
		importAlpine   = "IMPORT_ALPINE_JS"
		incMongo       = "INCLUDE_MONGO"
		incPostgres    = "INCLUDE_POSTGRES"
		incValkey      = "INCLUDE_VALKEY"
		mongodbName    = "MONGO_DB_NAME"
		mongoUsername  = "MONGO_USERNAME"
		mongoPassword  = "MONGO_PASSWORD"
		mongoHost      = "MONGO_HOST"
		mongoPort      = "MONGO_PORT"
	)

	// Test HTTP Config
	c, err := NewDefaultConfig()
	errInvalidPortString := "http port must be a valid port number"
	assertConfigNil(t, c, err, errInvalidPortString)

	os.Setenv(httpPort, "invalid")
	defer os.Unsetenv(httpPort)
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, errInvalidPortString)

	errInvalidPortRange := "only registered and dynamic ports are allowed (1000 - 65535)"
	os.Setenv(httpPort, "15")
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, errInvalidPortRange)

	os.Setenv(httpPort, "1512312")
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, errInvalidPortRange)

	os.Setenv(httpPort, "3000")
	c, err = NewDefaultConfig()
	assert.Equal(t, 3000, c.HTTP.Port, "expected the same http port")
	assert.Nil(t, err, "expected no errors")

	os.Setenv(fileserverPath, "non url safe path")
	defer os.Unsetenv(fileserverPath)
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "file server path is not URL safe")

	os.Setenv(fileserverPath, "static")
	os.Setenv(importAlpine, "true")
	defer os.Unsetenv(importAlpine)
	c, err = NewDefaultConfig()
	assert.True(t, c.HTTP.ImportAlpineJS, "should be true")
	assert.Equal(t, "static", c.HTTP.FileServerPath)
	assert.Nil(t, err, "expected no errors")
	assert.Nil(t, c.Mongo, "expected nil mongo config")
	assert.Nil(t, c.Postgres, "expected nil postgres config")
	assert.Nil(t, c.Valkey, "expected nil valkey config")

	// Test Mongo Config
	os.Setenv(incMongo, "invalid")
	defer os.Unsetenv(incMongo)
	c, err = NewDefaultConfig()
	assert.Nil(t, c.Mongo, "expected nil mongo config")
	assert.Nil(t, err, "expected no errors")

	os.Setenv(incMongo, "true")
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "mongo database name cannot be empty")

	os.Setenv(mongodbName, "testname")
	defer os.Unsetenv(mongodbName)
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "mongo username cannot be empty")

	os.Setenv(mongoUsername, "testusername")
	defer os.Unsetenv(mongoUsername)
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "mongo password cannot be empty")

	os.Setenv(mongoPassword, "password")
	defer os.Unsetenv(mongoPassword)
	c, err = NewDefaultConfig()
	assert.Equal(t, "testname", c.Mongo.DBName, "expected the same mongo database name")
	assert.Equal(t, "testusername", c.Mongo.Username, "expected the same usernames")
	assert.Equal(t, "password", c.Mongo.Password, "expected the same passwords")
	assert.Equal(t, "27017", c.Mongo.Port, "expected the same mongo ports")
	assert.Equal(t, "localhost", c.Mongo.Host, "expected the same hosts")
	assert.Nil(t, err, "expected no errors")

	os.Setenv(mongoPort, "invalid port")
	defer os.Unsetenv(mongoPort)
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "invalid port for mongo connection")

	os.Setenv(mongoPort, "123456")
	os.Setenv(mongoHost, "someotherhost")
	defer os.Unsetenv(mongoHost)
	c, err = NewDefaultConfig()
	assert.Equal(t, "123456", c.Mongo.Port, "expected the same ports")
	assert.Equal(t, "someotherhost", c.Mongo.Host, "expected the same hosts")
	assert.Nil(t, err, "expected no errors")

	// Test Postgres Config
	os.Setenv(incPostgres, "invalid")
	defer os.Unsetenv(incPostgres)
	c, err = NewDefaultConfig()
	assert.Nil(t, c.Postgres, "expected nil postgres config")
	assert.Nil(t, err, "expected no errors")

	os.Setenv(incPostgres, "true")
	c, err = NewDefaultConfig()
	assert.NotNil(t, c.Postgres, "expected postgres config")
	assert.Nil(t, err, "expected no errors")

	// Test Valkey Config
	os.Setenv(incValkey, "invalid")
	defer os.Unsetenv(incValkey)
	c, err = NewDefaultConfig()
	assert.Nil(t, c.Valkey, "expected nil valkey config")
	assert.Nil(t, err, "expected no errors")

	os.Setenv(incValkey, "true")
	c, err = NewDefaultConfig()
	assert.NotNil(t, c.Valkey, "expected valkey cnofiguration")
	assert.Nil(t, err, "expected no errors")
}

func assertConfigNil(t *testing.T, c *Config, err error, msg string) {
	assert.Nil(t, c, "expected nil config")
	assert.Error(t, err, "expected an error")
	assert.Equal(t, err.Error(), msg, "expected the same error message")
}
