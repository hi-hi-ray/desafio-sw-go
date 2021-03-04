package handler

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/hi-hi-ray/desafio-sw-go/api/errors"
	"github.com/hi-hi-ray/desafio-sw-go/api/models"
	"github.com/hi-hi-ray/desafio-sw-go/api/services"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		pingStatus := services.HealthCheckDatabaseService()
		responseApi := HealthyCheckResponse{
			Code:             "200",
			ApiMessageStatus: "I`m Well and Alive, Thank you for checking me",
			DatabaseStatus:   pingStatus,
		}
		HealthCheckResponseApi(responseApi, w)
	} else {
		responseApi := GenericResponse{Code: "405", Description: "Method Not Allowed: To check if i`m fine, trying use GET"}
		GenericResponseApi(responseApi, w)
	}
}

func GetPlanets(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		planets, err := services.GetPlanetsService()
		if err != nil {
			responseApi := GenericResponse{Code: "500", Description: "An error happened while trying to get all planets",
				Error: err.Error()}
			GenericResponseApi(responseApi, w)
		} else {
			responseApi := PlanetsResponse{Code: "200", Count: len(planets), Planets: planets}
			PlanetsResponseApi(responseApi, w)
		}
	} else {
		responseApi := GenericResponse{Code: "405", Description: "Method Not Allowed: To get all planets, trying use GET"}
		GenericResponseApi(responseApi, w)
	}
}

func GetPlanet(params martini.Params, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if params["id"] != "" {
			planet, err := services.GetPlanetService(params["id"])
			if err != nil {
				if err.Error() == "mongo: no documents in result" {
					responseApi := GenericResponse{Code: "404",
						Description: "An error happened while trying to get the planet.",
						Error:       err.Error()}
					GenericResponseApi(responseApi, w)
				} else {
					responseApi := GenericResponse{Code: "500",
						Description: "An error happened while trying to get the planet.",
						Error:       err.Error()}
					GenericResponseApi(responseApi, w)
				}
			} else {
				responseApi := PlanetResponse{Code: "200", Planet: planet}
				PlanetResponseApi(responseApi, w)
			}
		} else {
			responseApi := GenericResponse{Code: "422", Description: "Id is required"}
			GenericResponseApi(responseApi, w)
		}
	} else {
		responseApi := GenericResponse{Code: "405", Description: "Method Not Allowed: To get, trying use GET"}
		GenericResponseApi(responseApi, w)
	}
}

func CreatePlanet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		defer r.Body.Close()
		var planet models.Planet

		if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
			responseApi := GenericResponse{Code: "400", Description: "Invalid request", Error: err.Error()}
			GenericResponseApi(responseApi, w)
		} else {
			var err error
			var itPassed bool

			err, itPassed = services.PlanetModelValidator(&planet)
			planetExists := services.PlanetExistsValidator(&planet)
			if itPassed != true {
				responseApi := GenericResponse{Code: "406", Description: "An error happened while trying to create the planet", Error: err.Error()}
				GenericResponseApi(responseApi, w)
			} else if planetExists != false {
				responseApi := GenericResponse{Code: "409", Description: "An error happened while trying to create the planet",
					Error: errors.Create(errors.PlanetExist).Error()}
				GenericResponseApi(responseApi, w)
			} else {
				err = services.CreatePlanetService(&planet)
				if err != nil {
					responseApi := GenericResponse{Code: "500", Description: "An error happened while trying to create the planet",
						Error: err.Error()}
					GenericResponseApi(responseApi, w)
				} else {
					responseApi := GenericResponse{Code: "201", Description: "Planet inserted with success"}
					GenericResponseApi(responseApi, w)
				}
			}
		}
	} else {
		responseApi := GenericResponse{Code: "405", Description: "Method Not Allowed: To create, trying use POST"}
		GenericResponseApi(responseApi, w)
	}
}

func UpdatePlanet(params martini.Params, w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		if params["id"] != "" {
			var planet models.Planet
			defer r.Body.Close()
			if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
				responseApi := GenericResponse{Code: "400", Description: "Invalid request", Error: err.Error()}
				GenericResponseApi(responseApi, w)
			} else {
				var err error
				var itPassed bool

				err, itPassed = services.PlanetModelValidator(&planet)
				if itPassed != true {
					responseApi := GenericResponse{Code: "406", Description: "An error happened while trying to create the planet", Error: err.Error()}
					GenericResponseApi(responseApi, w)
				} else {
					err = services.UpdatePlanetService(params["id"], planet)
					if err != nil {
						responseApi := GenericResponse{Code: "500", Description: "An error happened while trying to create the planet",
							Error: err.Error()}
						GenericResponseApi(responseApi, w)
					} else {
						responseApi := GenericResponse{Code: "200", Description: "Planet Updated with success"}
						GenericResponseApi(responseApi, w)
					}
				}
			}
		} else {
			responseApi := GenericResponse{Code: "422", Description: "Id is required"}
			GenericResponseApi(responseApi, w)
		}
	} else {
		responseApi := GenericResponse{Code: "405", Description: "Method Not Allowed: To update, trying use PUT"}
		GenericResponseApi(responseApi, w)
	}
}

func DeletePlanet(params martini.Params, w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		var err error
		if params["id"] != "" {
			err = services.DeletePlanetService(params["id"])
			if err != nil {
				responseApi := GenericResponse{Code: "500", Description: "An error happened while trying to delete the planet",
					Error: err.Error()}
				GenericResponseApi(responseApi, w)
			} else if err.Error() != errors.PlanetDoesNotExist {
				responseApi := GenericResponse{Code: "404", Description: "Planet not found.", Error: err.Error()}
				GenericResponseApi(responseApi, w)
			} else {
				responseApi := GenericResponse{Code: "200", Description: "Planet Deleted with success"}
				GenericResponseApi(responseApi, w)
			}
		} else {
			responseApi := GenericResponse{Code: "422", Description: "Id is required"}
			GenericResponseApi(responseApi, w)
		}
	} else {
		responseApi := GenericResponse{Code: "405", Description: "Method Not Allowed: To delete, trying use DELETE"}
		GenericResponseApi(responseApi, w)
	}
}

func GetPlanetByName(params martini.Params, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if params["name"] != "" {
			planet, err := services.GetPlanetByNameService(params["name"])
			if err != nil {
				if err.Error() == "mongo: no documents in result" {
					responseApi := GenericResponse{Code: "404",
						Description: "An error happened while trying to get the planet.",
						Error:       err.Error()}
					GenericResponseApi(responseApi, w)
				} else {
					responseApi := GenericResponse{Code: "500",
						Description: "An error happened while trying to get the planet.",
						Error:       err.Error()}
					GenericResponseApi(responseApi, w)
				}
			} else {
				responseApi := PlanetResponse{Code: "200", Planet: planet}
				PlanetResponseApi(responseApi, w)
			}
		} else {
			responseApi := GenericResponse{Code: "422", Description: "Name is required"}
			GenericResponseApi(responseApi, w)
		}
	} else {
		responseApi := GenericResponse{Code: "405", Description: "Method Not Allowed: To get, trying use GET"}
		GenericResponseApi(responseApi, w)
	}
}
