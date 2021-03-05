package config

import (
	"fmt"
	"github.com/hi-hi-ray/desafio-sw-go/api/errors"
	"log"
	"strings"
)

func ConfigNullException(needToFill []string) {
	if len(needToFill) != 0 {
		errorMessage := fmt.Sprintln(errors.MissingConfigVariable, strings.Join(needToFill, ", "))
		log.Panic(errorMessage)
	}
}
