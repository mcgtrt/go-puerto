package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	ts := "true"
	// Test HTTP Config
	c, err := NewDefaultConfig()
	errInvalidPortString := "http port must be a valid port number"
	assertConfigNil(t, c, err, errInvalidPortString)

	os.Setenv(HTTP_PORT, "invalid")
	defer os.Unsetenv(HTTP_PORT)
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, errInvalidPortString)

	errInvalidPortRange := "only registered and dynamic ports are allowed (1000 - 65535)"
	os.Setenv(HTTP_PORT, "15")
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, errInvalidPortRange)

	os.Setenv(HTTP_PORT, "1512312")
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, errInvalidPortRange)

	os.Setenv(HTTP_PORT, "3000")
	c, err = NewDefaultConfig()
	assert.Equal(t, 3000, c.HTTP.Port, "expected the same http port")
	assert.Nil(t, err, "expected no errors")

	os.Setenv(FILE_SERVER_PATH, "non url safe path")
	defer os.Unsetenv(FILE_SERVER_PATH)
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "file server path is not URL safe")

	os.Setenv(FILE_SERVER_PATH, "static")
	os.Setenv(USE_JS_ALPINE, ts)
	defer os.Unsetenv(USE_JS_ALPINE)
	c, err = NewDefaultConfig()
	assert.True(t, c.HTTP.ImportAlpineJS, "should be true")
	assert.Equal(t, "static", c.HTTP.FileServerPath)
	assert.Nil(t, err, "expected no errors")
	assert.Nil(t, c.Mongo, "expected nil mongo config")
	assert.Nil(t, c.Postgres, "expected nil postgres config")
	assert.Nil(t, c.Valkey, "expected nil valkey config")

	// Test Middleware Config
	mws := []string{
		USE_MW_LOCALISATION,
		USE_MW_SECURE_HEADERS,
		USE_MW_RATE_LIMIT,
		USE_MW_LOG_AND_MONITOR_HEADERS,
		USE_MW_CORS,
		USE_MW_ETAG,
		USE_MW_VALIDATE_SANITISE_HEADERS,
		USE_MW_METHOD_OVERRIDE,
	}
	for _, mw := range mws {
		os.Setenv(mw, ts)
		defer os.Unsetenv(mw)
	}
	c, err = NewDefaultConfig()
	assert.True(t, c.Middleware.Localisation, "expected localisation mw true")
	assert.True(t, c.Middleware.SecureHeaders, "expected secure headers mw true")
	assert.True(t, c.Middleware.RateLimit, "expected rate limit mw true")
	assert.True(t, c.Middleware.LogAndMonitorHeaders, "expected log and monitor mw true")
	assert.True(t, c.Middleware.CORS, "expected cors mw true")
	assert.True(t, c.Middleware.ETAG, "expected etag mw true")
	assert.True(t, c.Middleware.ValidateSanitiseHeaders, "expected validate and sanitise headers mw true")
	assert.True(t, c.Middleware.MethodOverride, "expected method override mw true")
	assert.Nil(t, err, "expected no errors")

	// Test Mongo Config
	os.Setenv(USE_DB_MONGO, "invalid")
	defer os.Unsetenv(USE_DB_MONGO)
	c, err = NewDefaultConfig()
	assert.Nil(t, c.Mongo, "expected nil mongo config")
	assert.Nil(t, err, "expected no errors")

	os.Setenv(USE_DB_MONGO, ts)
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "mongo database name cannot be empty")

	os.Setenv(MONGO_DB_NAME, "testname")
	defer os.Unsetenv(MONGO_DB_NAME)
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "mongo username cannot be empty")

	os.Setenv(MONGO_USERNAME, "testusername")
	defer os.Unsetenv(MONGO_USERNAME)
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "mongo password cannot be empty")

	os.Setenv(MONGO_PASSWORD, "password")
	defer os.Unsetenv(MONGO_PASSWORD)
	c, err = NewDefaultConfig()
	assert.Equal(t, "testname", c.Mongo.DBName, "expected the same mongo database name")
	assert.Equal(t, "testusername", c.Mongo.Username, "expected the same usernames")
	assert.Equal(t, "password", c.Mongo.Password, "expected the same passwords")
	assert.Equal(t, "27017", c.Mongo.Port, "expected the same mongo ports")
	assert.Equal(t, "localhost", c.Mongo.Host, "expected the same hosts")
	assert.Nil(t, err, "expected no errors")

	os.Setenv(MONGO_PORT, "invalid port")
	defer os.Unsetenv(MONGO_PORT)
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "invalid port for mongo connection")

	os.Setenv(MONGO_PORT, "123456")
	os.Setenv(MONGO_HOST, "someotherhost")
	defer os.Unsetenv(MONGO_HOST)
	c, err = NewDefaultConfig()
	assert.Equal(t, "123456", c.Mongo.Port, "expected the same ports")
	assert.Equal(t, "someotherhost", c.Mongo.Host, "expected the same hosts")
	assert.Nil(t, err, "expected no errors")

	// Test Postgres Config
	os.Setenv(USE_DB_POSTGRES, "invalid")
	defer os.Unsetenv(USE_DB_POSTGRES)
	c, err = NewDefaultConfig()
	assert.Nil(t, c.Postgres, "expected nil postgres config")
	assert.Nil(t, err, "expected no errors")

	os.Setenv(USE_DB_POSTGRES, ts)
	c, err = NewDefaultConfig()
	assert.NotNil(t, c.Postgres, "expected postgres config")
	assert.Nil(t, err, "expected no errors")

	// Test Valkey Config
	os.Setenv(USE_DB_VALKEY, "invalid")
	defer os.Unsetenv(USE_DB_VALKEY)
	c, err = NewDefaultConfig()
	assert.Nil(t, c.Valkey, "expected nil valkey config")
	assert.Nil(t, err, "expected no errors")

	os.Setenv(USE_DB_VALKEY, ts)
	c, err = NewDefaultConfig()
	assert.NotNil(t, c.Valkey, "expected valkey cnofiguration")
	assert.Nil(t, err, "expected no errors")
}

func assertConfigNil(t *testing.T, c *Config, err error, msg string) {
	assert.Nil(t, c, "expected nil config")
	assert.Error(t, err, "expected an error")
	assert.Equal(t, err.Error(), msg, "expected the same error message")
}
