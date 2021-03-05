package unit

import (
	"github.com/hi-hi-ray/desafio-sw-go/api/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateErrorsSuccess(t *testing.T)  {
	 messageError := errors.Create("teste")
	 assert.EqualValues(t, "teste", messageError.Error())
}

func TestCreateErrorsEmpty(t *testing.T)  {
	messageError := errors.Create("")
	assert.EqualValues(t, "", messageError.Error())
}