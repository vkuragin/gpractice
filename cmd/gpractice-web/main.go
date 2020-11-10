package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/vkuragin/gpractice"
	"github.com/vkuragin/gpractice/repo"
	"log"
	"net/http"
)

func main() {
	log.Printf("starting\n")

	// parse flags
	templatePath := flag.String("tplpath", "", "path to custom html template")
	templateRefresh := flag.Bool("tplfresh", false, "parse template each time")
	flag.Parse()

	// initialize services
	sqlRepo := repo.MySQLRepo{}
	sqlRepo.Init()
	gPractice := gpractice.GPractice{&sqlRepo}
	holder := tplHolder{refresh: *templateRefresh, tplPath: *templatePath}

	router := mux.NewRouter()

	// rest api endpoints
	restHandler := restHandler{gPractice}
	router.HandleFunc("/rest", restHandler.restAll()).Methods(http.MethodGet)
	router.HandleFunc("/rest/", restHandler.restAll()).Methods(http.MethodGet)
	router.HandleFunc("/rest", restHandler.restAdd()).Methods(http.MethodPost)
	router.HandleFunc("/rest/{id}", restHandler.restGet()).Methods(http.MethodGet)
	router.HandleFunc("/rest/{id}", restHandler.restUpdate()).Methods(http.MethodPost, http.MethodPut)
	router.HandleFunc("/rest/{id}", restHandler.restDelete()).Methods(http.MethodDelete)

	// web app endpoints
	appHandler := appHandler{gPractice, holder}
	router.HandleFunc("/app", appHandler.appAll()).Methods(http.MethodGet)
	router.HandleFunc("/app/", appHandler.appAll()).Methods(http.MethodGet)
	router.HandleFunc("/app", appHandler.appAdd()).Methods(http.MethodPost)
	router.HandleFunc("/app/{id}", appHandler.appGet()).Methods(http.MethodGet)
	router.HandleFunc("/app/{id}", appHandler.appUpdate()).Methods(http.MethodPost, http.MethodPut)
	router.HandleFunc("/app/{id}", appHandler.appDelete()).Methods(http.MethodDelete)

	// run server
	log.Fatal(http.ListenAndServe(":3000", router))
}
