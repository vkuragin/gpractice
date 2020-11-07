package main

import (
	"flag"
	"github.com/vk23/gpractice"
	"github.com/vk23/gpractice/repo"
	"html/template"
	"log"
	"net/http"
)

func main() {
	log.Printf("starting\n")

	templatePath := flag.String("template", "", "path to custom html template")
	flag.Parse()

	sqlRepo := repo.MySQLRepo{}
	sqlRepo.Init()
	gPractice := gpractice.GPractice{&sqlRepo}

	http.HandleFunc("/", handleRoot(gPractice, *templatePath))

	log.Fatal(http.ListenAndServe(":8001", nil))
}

func handleRoot(gPractice gpractice.GPractice, templatePath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("---\nServing page: %v %v\n", r.Method, r.URL)

		//TODO: cache template
		tmplt := getTemplate(templatePath)

		var pageData repo.PageData
		var err error

		pageData.Items, err = gPractice.GetAll()
		if err != nil {
			log.Printf("Error: %v\n", err)
			return
		}

		pageData.Report, err = gPractice.GetReport()
		if err != nil {
			log.Printf("Error: %v\n", err)
			return
		}

		err = tmplt.Execute(w, pageData)
		if err != nil {
			log.Printf("Error template: %v\n", err)
			return
		}

		log.Printf("Serving page done\n---\n")
	}
}

func getTemplate(templatePath string) *template.Template {
	var result *template.Template
	var err error

	if templatePath != "" {
		result, err = template.ParseFiles(templatePath)
	} else {
		tpl := template.New("defaultHtmlTemplate")
		result, err = tpl.Parse(defaultHtmlTemplate)
	}

	if err != nil {
		log.Fatalf("Error template: %v\n", err)
	}

	return result
}
