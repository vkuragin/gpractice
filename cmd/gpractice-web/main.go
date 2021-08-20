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
	log.Print("starting web application")

	// parse flags
	templatePath := flag.String("tplpath", "", "path to custom html template")
	templateRefresh := flag.Bool("tplfresh", false, "parse template each time")
	flag.Parse()

	// initialize db
	sqlRepo := &repo.MySQLRepo{}
	sqlRepo.Init()

	// initialize services
	gPractice := gpractice.GPractice{sqlRepo}
	holder := tplHolder{refresh: *templateRefresh, tplPath: *templatePath}
	router := configureRouter(gPractice, holder)

	// run server
	log.Print("web application started")
	log.Fatal(http.ListenAndServe("localhost:3000", router))
}

func configureRouter(service gpractice.GPractice, holder tplHolder) http.Handler {
	router := mux.NewRouter()

	// rest api endpoints
	restHandler := restHandler{service}
	router.HandleFunc("/rest", restHandler.restAll()).Methods(http.MethodGet)
	router.HandleFunc("/rest/", restHandler.restAll()).Methods(http.MethodGet)
	router.HandleFunc("/rest", restHandler.restAdd()).Methods(http.MethodPost)
	router.HandleFunc("/rest/{id}", restHandler.restGet()).Methods(http.MethodGet)
	router.HandleFunc("/rest/{id}", restHandler.restUpdate()).Methods(http.MethodPost, http.MethodPut)
	router.HandleFunc("/rest/{id}", restHandler.restDelete()).Methods(http.MethodDelete)

	// web app endpoints
	appHandler := appHandler{service, holder}
	router.HandleFunc("/app", appHandler.appAll()).Methods(http.MethodGet)
	router.HandleFunc("/app/", appHandler.appAll()).Methods(http.MethodGet)
	router.HandleFunc("/app", appHandler.appAdd()).Methods(http.MethodPost)
	router.HandleFunc("/app/{id}", appHandler.appGet()).Methods(http.MethodGet)
	router.HandleFunc("/app/{id}", appHandler.appUpdate()).Methods(http.MethodPost, http.MethodPut)
	router.HandleFunc("/app/{id}", appHandler.appDelete()).Methods(http.MethodDelete)

	return router
}
