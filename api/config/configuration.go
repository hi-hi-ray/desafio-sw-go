package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/hi-hi-ray/desafio-sw-go/api/errors"
	"log"
)

type ConfigVariables struct {
	Database DatabaseMongo `toml:"database"`
	Servers  Servers       `toml:"servers"`
	Swapi    Swapi         `toml:"swapi"`
}

type DatabaseMongo struct {
	Server     string `toml:"server"`
	Database   string `toml:"database"`
	Collection string `toml:"collection"`
	Port       int    `toml:"port"`
	Username   string `toml:"username"`
	Password   string `toml:"password"`
	Timeout    int    `toml:"timeout"`
}

type Servers struct {
	Port int `toml:"port"`
}

type Swapi struct {
	Urlbase  string `toml:"urlbase"`
	Endpoint string `toml:"endpoint"`
}

func (configVar *ConfigVariables) Read() {
	_, err := toml.DecodeFile("./config/config.toml", &configVar)
	if err != nil {
		log.Panicln(fmt.Sprintln(errors.ProblemAtReadingFile, err.Error()))
	}
}

func (configVar *ConfigVariables) GetConfigEnvoriments() *ConfigVariables {
	configs := ConfigVariables{}
	configs.Read()
	needToFill := configs.ConfigValidatorFields()
	ConfigNullException(needToFill)
	return &configs
}
