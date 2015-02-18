package cmd

import (
	"log"
	"github.com/codegangsta/negroni"
	"github.com/spf13/viper"
	"github.com/gorilla/mux"

	"github.com/untoldone/bloomapi/middleware"
	"github.com/untoldone/bloomapi/handler"
	"github.com/untoldone/bloomapi/handler_compat"
)

func Server() {
	port := viper.GetString("bloomapiPort")
	n := negroni.Classic()

	// Middleware setup
	n.Use(middleware.NewAuthentication())

	// Router setup
	router := mux.NewRouter()

	// Current API
	router.HandleFunc("/api/sources", handler.SourcesHandler).Methods("GET")
	router.HandleFunc("/api/search/{source}", handler.SearchSourceHandler).Methods("GET")
	router.HandleFunc("/api/sources/{source}/{id}", handler.ItemHandler).Methods("GET")

	// For Backwards Compatibility Feb 13, 2015
	router.HandleFunc("/api/search", handler_compat.NpiSearchHandler).Methods("GET")
	router.HandleFunc("/api/search/npi", handler_compat.NpiSearchHandler).Methods("GET")
	router.HandleFunc("/api/npis/{npi:[0-9]+}", handler_compat.NpiItemHandler).Methods("GET")
	router.HandleFunc("/api/sources/npi/{npi:[0-9]+}", handler_compat.NpiItemHandler).Methods("GET")

	n.UseHandler(router)

	log.Println("Running Server")
	n.Run(":" + port)
}