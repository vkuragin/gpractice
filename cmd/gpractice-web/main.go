package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/vk23/gpractice"
	"github.com/vk23/gpractice/repo"
	"log"
	"net/http"
)

func main() {
	log.Printf("starting\n")

	// parse flags
	templatePath := flag.String("template", "", "path to custom html template")
	flag.Parse()

	// initialize services
	sqlRepo := repo.MySQLRepo{}
	sqlRepo.Init()
	gPractice := gpractice.GPractice{&sqlRepo}

	//TODO: cache template
	tmplt := getTemplate(*templatePath)

	router := mux.NewRouter()

	// rest api endpoints
	router.HandleFunc("/rest", restAll(gPractice)).Methods(http.MethodGet)
	router.HandleFunc("/rest", restAdd(gPractice)).Methods(http.MethodPost)
	router.HandleFunc("/rest/{id}", restGet(gPractice)).Methods(http.MethodGet)
	router.HandleFunc("/rest/{id}", restUpdate(gPractice)).Methods(http.MethodPost, http.MethodPut)
	router.HandleFunc("/rest/{id}", restDelete(gPractice)).Methods(http.MethodDelete)

	// web app endpoints
	router.HandleFunc("/app", appAll(gPractice, tmplt)).Methods(http.MethodGet)
	router.HandleFunc("/app", appAdd(gPractice, tmplt)).Methods(http.MethodPost)
	router.HandleFunc("/app/{id}", appGet(gPractice, tmplt)).Methods(http.MethodGet)
	router.HandleFunc("/app/{id}", appUpdate(gPractice, tmplt)).Methods(http.MethodPost, http.MethodPut)
	router.HandleFunc("/app/{id}", appDelete(gPractice, tmplt)).Methods(http.MethodDelete)

	// run server
	log.Fatal(http.ListenAndServe(":3000", router))
}
