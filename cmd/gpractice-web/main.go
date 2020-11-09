package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/vk23/gpractice"
	"github.com/vk23/gpractice/repo"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	log.Printf("starting\n")

	templatePath := flag.String("template", "", "path to custom html template")
	flag.Parse()

	sqlRepo := repo.MySQLRepo{}
	sqlRepo.Init()
	gPractice := gpractice.GPractice{&sqlRepo}

	router := mux.NewRouter()
	router.HandleFunc("/app", handleRoot(gPractice, *templatePath)).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/app/{id:[0-9]+}", handleItem(gPractice, *templatePath)).Methods(http.MethodGet, http.MethodDelete, http.MethodPut)
	log.Fatal(http.ListenAndServe(":3000", router))
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

func handleItem(gPractice gpractice.GPractice, templatePath string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("---\nServing page: %v %v\n", r.Method, r.URL)

		path := r.URL.Path
		vars := mux.Vars(r)
		log.Printf("path=%v, vars=%v\n", path, vars)

		//TODO: cache template
		tmplt := getTemplate(templatePath)

		var pageData repo.PageData
		var err error
		switch r.Method {
		case http.MethodGet:
			id, err := strconv.Atoi(vars["id"])
			if err != nil {
				log.Printf("Bad id error: %v\n", err)
				http.Error(w, "Bad id", http.StatusBadRequest)
				return
			}
			item, err := gPractice.Get(uint64(id))
			if err != nil {
				log.Printf("Failed to retrieve item by id %d, error: %v\n", path, err)
				http.Error(w, "Failed to retrieve item by id", http.StatusInternalServerError)
				return
			}

			pageData.Item = item
			err = tmplt.Execute(w, pageData)
			if err != nil {
				log.Printf("Error template: %v\n", err)
				return
			}
			return

		case http.MethodPut, http.MethodPost:
			log.Printf("PUT/POST")
			log.Printf("Unparsed form: %v\n", r.Form)
			err = r.ParseForm()
			if err != nil {
				log.Printf("Error parsing form: %v\n", err)
				http.Error(w, "invalid form", http.StatusBadRequest)
				return
			}
			log.Printf("Parsed form: %v\n", r.Form)

			idInput := r.Form.Get("idInput")
			id := 0
			if idInput != "" {
				id, err = strconv.Atoi(idInput)
				if err != nil {
					log.Printf("Error parsing form: %v\n", err)
					http.Error(w, "invalid form", http.StatusBadRequest)
					return
				}
			}
			date := r.Form.Get("dateInput")
			var validDate = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
			if !validDate.MatchString(date) {
				log.Printf("Error parsing form date: %v\n", date)
				http.Error(w, "invalid date", http.StatusBadRequest)
				return
			}
			duration, err := strconv.Atoi(r.Form.Get("durationInput"))
			if err != nil {
				log.Printf("Error parsing form: %v\n", err)
				http.Error(w, "invalid form", http.StatusBadRequest)
				return
			}
			item, err := gPractice.Save(repo.Item{Id: uint64(id), Date: date, Duration: uint64(duration)})
			if err != nil {
				log.Printf("Error saving item: %v\n", err)
				http.Error(w, "Failed to save item", http.StatusInternalServerError)
				return
			}
			log.Printf("Item saved: %v\n", item)
			log.Printf("Redirecting to root")
			http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		case http.MethodDelete:
			log.Printf("DELETE")
			id, err := strconv.Atoi(vars["id"])
			if err != nil {
				log.Printf("Bad id error: %v\n", err)
				http.Error(w, "Bad id", http.StatusBadRequest)
				return
			}
			var result = false
			result, err = gPractice.Delete(uint64(id))
			if err != nil {
				log.Printf("Error deleting item: %v\n", err)
				http.Error(w, "Failed to delete item", http.StatusInternalServerError)
				return
			}
			log.Printf("Delete result: %v\n", result)

			http.Redirect(w, r, "/", http.StatusSeeOther)

		default:
			http.Error(w, "", http.StatusMethodNotAllowed)
		}
		log.Printf("Serving page done\n---\n")
	}
}

func handleRoot(gPractice gpractice.GPractice, templatePath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("---\nServing page: %v %v\n", r.Method, r.URL)

		//TODO: cache template
		tmplt := getTemplate(templatePath)

		var pageData repo.PageData
		var err error

		switch r.Method {
		case http.MethodGet:
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

		case http.MethodPost:
			log.Printf("There will be a new item creation")

			log.Printf("Unparsed form: %v\n", r.Form)
			err = r.ParseForm()
			if err != nil {
				log.Printf("Error parsing form: %v\n", err)
				http.Error(w, "invalid form", http.StatusBadRequest)
				return
			}
			log.Printf("Parsed form: %v\n", r.Form)

			idInput := r.Form.Get("idInput")
			id := 0
			if idInput != "" {
				id, err = strconv.Atoi(idInput)
				if err != nil {
					log.Printf("Error parsing form: %v\n", err)
					http.Error(w, "invalid form", http.StatusBadRequest)
					return
				}
			}
			date := r.Form.Get("dateInput")
			var validDate = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
			if !validDate.MatchString(date) {
				log.Printf("Error parsing form date: %v\n", date)
				http.Error(w, "invalid date", http.StatusBadRequest)
				return
			}
			duration, err := strconv.Atoi(r.Form.Get("durationInput"))
			if err != nil {
				log.Printf("Error parsing form: %v\n", err)
				http.Error(w, "invalid form", http.StatusBadRequest)
				return
			}
			item, err := gPractice.Save(repo.Item{Id: uint64(id), Date: date, Duration: uint64(duration)})
			if err != nil {
				log.Printf("Error saving item: %v\n", err)
				http.Error(w, "Failed to save item", http.StatusInternalServerError)
				return
			}
			log.Printf("Item saved: %v\n", item)

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
		default:
			http.Error(w, "", http.StatusMethodNotAllowed)
		}

		log.Printf("Serving page done\n---\n")
	}
}
