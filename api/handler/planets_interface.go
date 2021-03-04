package handler


import (
"net/http"
)

type PlanetHandle interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
	GetPlanets(w http.ResponseWriter, r *http.Request)
	GetPlanet(w http.ResponseWriter, r *http.Request)
	CreatePlanet(w http.ResponseWriter, r *http.Request)
	UpdatePlanet(w http.ResponseWriter, r *http.Request)
	DeletePlanet(w http.ResponseWriter, r *http.Request)
}
