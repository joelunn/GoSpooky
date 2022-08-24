package main

import (
	"app/config"
	"app/controllers"
	"app/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/endpoints", controllers.GetEndpoints).Methods("GET")
	router.HandleFunc("/api/endpoints/{id}", controllers.GetEndpointById).Methods("GET")
	router.HandleFunc("/api/endpoints", controllers.CreateEndpoint).Methods("POST")
	router.HandleFunc("/api/endpoints/{id}", controllers.UpdateEndpoint).Methods("PUT")
	router.HandleFunc("/api/endpoints/{id}", controllers.DeleteEndpoint).Methods("DELETE")
}

func main() {

	config.LoadAppConfig()
	//config.DatabaseConfigurations.GetConnectString()
	fmt.Println("Database server connection string is \t", config.GetConnectString())

	database.Connect(config.GetConnectString())
	database.InitialiseDB()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)
	// Register Routes
	RegisterProductRoutes(router)

	log.Println(fmt.Sprintf("Starting Server on port %s", config.GetServerPort()), router)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.GetServerPort()), router))
}
