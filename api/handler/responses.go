package handler

import (
	"encoding/json"
	"fmt"
	"github.com/hi-hi-ray/desafio-sw-go/api/models"
	"log"
	"net/http"
	"strconv"
)

type GenericResponse struct {
	Code        string `json:"code"`
	Description string `json:"description, omitempty"`
	Error       string `json:"error, omitempty"`
}

type HealthyCheckResponse struct {
	Code             string `json:"code"`
	ApiMessageStatus string `json:"api_status, omitempty"`
	DatabaseStatus   string `json:"database_status, omitempty"`
}

type PlanetsResponse struct {
	Code    string           `json:"code"`
	Count   int              `json:"count, omitempty"`
	Planets []*models.Planet `json:"planets, omitempty"`
}

type PlanetResponse struct {
	Code   string        `json:"code"`
	Count  int           `json:"count, omitempty"`
	Planet models.Planet `json:"planets, omitempty"`
}

func GenericResponseApi(responseApi GenericResponse, w http.ResponseWriter) {
	statusCode, _ := strconv.Atoi(responseApi.Code)
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	responseStatus, err := json.Marshal(responseApi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, errWriter := w.Write(responseStatus)

	if errWriter != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PlanetResponseApi(responseApi PlanetResponse, w http.ResponseWriter) {
	statusCode, _ := strconv.Atoi(responseApi.Code)
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	responseStatus, err := json.Marshal(responseApi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, errWriter := w.Write(responseStatus)

	if errWriter != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PlanetsResponseApi(responseApi PlanetsResponse, w http.ResponseWriter) {
	statusCode, _ := strconv.Atoi(responseApi.Code)
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	responseStatus, err := json.Marshal(responseApi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, errWriter := w.Write(responseStatus)

	if errWriter != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HealthCheckResponseApi(responseApi HealthyCheckResponse, w http.ResponseWriter) {
	statusCode, _ := strconv.Atoi(responseApi.Code)
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	responseStatus, err := json.Marshal(responseApi)
	if err != nil {
		log.Println(fmt.Printf("An error appears: %v", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, errWriter := w.Write(responseStatus)

	if errWriter != nil {
		log.Println(fmt.Printf("An error appears: %v", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
