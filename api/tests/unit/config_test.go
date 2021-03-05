package unit

import (
	"fmt"
	"github.com/hi-hi-ray/desafio-sw-go/api/config"
	"github.com/hi-hi-ray/desafio-sw-go/api/errors"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestConfigValidatorFieldsSuccess(t *testing.T) {

	configMock := &config.ConfigVariables{
		Database: config.DatabaseMongo{
			Server:     "MockServerDB",
			Database:   "MockDatabaseName",
			Collection: "MockCollectionName",
			Port:       0001,
			Username:   "",
			Password:   "",
			Timeout:    30,
		},
		Servers: config.Servers{
			Port: 0001,
		},
		Swapi: config.Swapi{
			Urlbase:  "https://swapi.co/",
			Endpoint: "api/planets",
		},
	}
	assert.Equal(t, 0, len(configMock.ConfigValidatorFields()))
}

func TestConfigValidatorFieldsFailDatabase(t *testing.T) {

	configMock := &config.ConfigVariables{
		Database: config.DatabaseMongo{
			Server:     "",
			Database:   "",
			Collection: "",
			Port:       0,
			Username:   "",
			Password:   "",
			Timeout:    0,
		},
		Servers: config.Servers{
			Port: 0001,
		},
		Swapi: config.Swapi{
			Urlbase:  "https://swapi.co/",
			Endpoint: "api/planets",
		},
	}
	assert.Equal(t, 5, len(configMock.ConfigValidatorFields()))
}

func TestConfigValidatorFieldsFailServers(t *testing.T) {

	configMock := &config.ConfigVariables{
		Database: config.DatabaseMongo{
			Server:     "MockServerDB",
			Database:   "MockDatabaseName",
			Collection: "MockCollectionName",
			Port:       0001,
			Username:   "",
			Password:   "",
			Timeout:    30,
		},
		Servers: config.Servers{
			Port: 0,
		},
		Swapi: config.Swapi{
			Urlbase:  "https://swapi.co/",
			Endpoint: "api/planets",
		},
	}
	assert.Equal(t, 1, len(configMock.ConfigValidatorFields()))
}

func TestConfigValidatorFieldsFailSwapi(t *testing.T) {

	configMock := &config.ConfigVariables{
		Database: config.DatabaseMongo{
			Server:     "MockServerDB",
			Database:   "MockDatabaseName",
			Collection: "MockCollectionName",
			Port:       0001,
			Username:   "",
			Password:   "",
			Timeout:    30,
		},
		Servers: config.Servers{
			Port: 0001,
		},
		Swapi: config.Swapi{
			Urlbase:  "",
			Endpoint: "",
		},
	}
	assert.Equal(t, 2, len(configMock.ConfigValidatorFields()))
}

func TestConfigNullExceptionFailDatabase(t *testing.T) {

	configMock := &config.ConfigVariables{
		Database: config.DatabaseMongo{
			Server:     "",
			Database:   "",
			Collection: "",
			Port:       0,
			Username:   "",
			Password:   "",
			Timeout:    30,
		},
		Servers: config.Servers{
			Port: 0001,
		},
		Swapi: config.Swapi{
			Urlbase:  "https://swapi.co/",
			Endpoint: "api/planets",
		},
	}
	needToFill := configMock.ConfigValidatorFields
	expectedErrorMessage := fmt.Sprintln(errors.MissingConfigVariable, strings.Join(needToFill(), ", "))
	assert.Panics(t, func() { config.ConfigNullException(needToFill()) })
	assert.PanicsWithValue(t, expectedErrorMessage, func() { config.ConfigNullException(needToFill()) })
}

func TestConfigNullExceptionFailServers(t *testing.T) {

	configMock := &config.ConfigVariables{
		Database: config.DatabaseMongo{
			Server:     "MockServerDB",
			Database:   "MockDatabaseName",
			Collection: "MockCollectionName",
			Port:       0001,
			Username:   "",
			Password:   "",
			Timeout:    30,
		},
		Servers: config.Servers{
			Port: 0,
		},
		Swapi: config.Swapi{
			Urlbase:  "https://swapi.co/",
			Endpoint: "api/planets",
		},
	}
	needToFill := configMock.ConfigValidatorFields
	expectedErrorMessage := fmt.Sprintln(errors.MissingConfigVariable, strings.Join(needToFill(), ", "))
	assert.Panics(t, func() { config.ConfigNullException(needToFill()) })
	assert.PanicsWithValue(t, expectedErrorMessage, func() { config.ConfigNullException(needToFill()) })
}

func TestConfigNullExceptionFailSwapi(t *testing.T) {

	configMock := &config.ConfigVariables{
		Database: config.DatabaseMongo{
			Server:     "MockServerDB",
			Database:   "MockDatabaseName",
			Collection: "MockCollectionName",
			Port:       0001,
			Username:   "",
			Password:   "",
			Timeout:    30,
		},
		Servers: config.Servers{
			Port: 0,
		},
		Swapi: config.Swapi{
			Urlbase:  "",
			Endpoint: "",
		},
	}
	needToFill := configMock.ConfigValidatorFields
	expectedErrorMessage := fmt.Sprintln(errors.MissingConfigVariable, strings.Join(needToFill(), ", "))
	assert.Panics(t, func() { config.ConfigNullException(needToFill()) })
	assert.PanicsWithValue(t, expectedErrorMessage, func() { config.ConfigNullException(needToFill()) })
}

