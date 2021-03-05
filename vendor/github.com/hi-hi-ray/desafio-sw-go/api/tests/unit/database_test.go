package unit

import (
	"github.com/hi-hi-ray/desafio-sw-go/api/database"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUtilitiesCreateTimeoutSuccess(t *testing.T)  {
	connectionMock := database.DatabaseConnection{
		Server:     "MockServerDB",
		Database:   "MockDatabaseName",
		Collection: "MockCollectionName",
		Port:       0001,
		Username:   "",
		Password:   "",
		Timeout:    5,
	}
	expectedAnwser := 5 * time.Second
	anwser := connectionMock.CreateTimeout()
	assert.Equal(t, expectedAnwser, anwser)
}

