package controllers

import (
	"app/database"
	"app/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var endpoint models.Endpoint
	json.NewDecoder(r.Body).Decode(&endpoint)
	database.Instance.Create(&endpoint)
	json.NewEncoder(w).Encode(endpoint)
}

func GetEndpointById(w http.ResponseWriter, r *http.Request) {
	endpointId := mux.Vars(r)["id"]
	if checkIfEndpointExists(endpointId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var endpoint models.Endpoint
	database.Instance.First(&endpoint, endpointId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(endpoint)
}

func GetEndpoints(w http.ResponseWriter, r *http.Request) {
	var endpoint []models.Endpoint
	database.Instance.Find(&endpoint)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(endpoint)
}

func UpdateEndpoint(w http.ResponseWriter, r *http.Request) {
	endpointId := mux.Vars(r)["id"]
	if checkIfEndpointExists(endpointId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var endpoint models.Endpoint
	database.Instance.First(&endpoint, endpointId)
	json.NewDecoder(r.Body).Decode(&endpoint)
	database.Instance.Save(&endpoint)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(endpoint)
}

func DeleteEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	endpointId := mux.Vars(r)["id"]
	if checkIfEndpointExists(endpointId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var endpoint models.Endpoint
	database.Instance.Delete(&endpoint, endpointId)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}

func checkIfEndpointExists(endpointId string) bool {
	var endpoint models.Endpoint
	database.Instance.First(&endpoint, endpointId)
	if endpoint.EndpointID == "null" { //change to work with guids
		return false
	}
	return true
}
