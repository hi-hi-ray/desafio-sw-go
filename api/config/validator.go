package config

import (
	"strings"
)

func (configVar *ConfigVariables) ConfigValidatorFields() []string {
	var needToFill []string
	if len(strings.TrimSpace(configVar.Database.Server)) == 0 {
		needToFill = append(needToFill, "Database Server")
	}
	if len(strings.TrimSpace(configVar.Database.Database)) == 0 {
		needToFill = append(needToFill, "Database Database")
	}
	if len(strings.TrimSpace(configVar.Database.Collection)) == 0 {
		needToFill = append(needToFill, "Database Collection")
	}
	if len(strings.TrimSpace(string(configVar.Database.Port))) == 0 || configVar.Database.Port == 0 {
		needToFill = append(needToFill, "Database Port")
	}
	if len(strings.TrimSpace(string(configVar.Database.Timeout))) == 0 || configVar.Database.Timeout == 0 {
		needToFill = append(needToFill, "Database Timeout")
	}
	if len(strings.TrimSpace(string(configVar.Servers.Port))) == 0 || configVar.Servers.Port == 0 {
		needToFill = append(needToFill, "Server Port")
	}
	if len(strings.TrimSpace(configVar.Swapi.Urlbase)) == 0 {
		needToFill = append(needToFill, "Swapi Urlbase")
	}
	if len(strings.TrimSpace(configVar.Swapi.Endpoint)) == 0 {
		needToFill = append(needToFill, "Swapi Planet Endpoint")
	}
	return needToFill
}
