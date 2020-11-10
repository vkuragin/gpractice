package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vk23/gpractice"
	"github.com/vk23/gpractice/repo"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func restAll(gPractice gpractice.GPractice) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pageData := getAll(w, gPractice)
		jsonBytes, err := json.Marshal(pageData)
		if err != nil {
			log.Printf("Json error: %v\n", err)
			http.Error(w, "Json error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}
}

func appAll(gPractice gpractice.GPractice, tmplt *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pageData := getAll(w, gPractice)
		err := tmplt.Execute(w, pageData)
		if err != nil {
			log.Printf("Template error: %v\n", err)
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}
	}
}

func getAll(w http.ResponseWriter, gPractice gpractice.GPractice) repo.PageData {
	pageData := repo.PageData{}
	var err error

	log.Printf("getAll\n")

	pageData.Items, err = gPractice.GetAll()
	if err != nil {
		log.Printf("Error: %v\n", err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return pageData
	}

	pageData.Report, err = gPractice.GetReport()
	if err != nil {
		log.Printf("Error: %v\n", err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return pageData
	}
	return pageData
}

func restAdd(gPractice gpractice.GPractice) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pageData := addItem(w, gPractice, r)
		jsonBytes, err := json.Marshal(pageData)
		if err != nil {
			log.Printf("Json error: %v\n", err)
			http.Error(w, "Json error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}
}

func appAdd(gPractice gpractice.GPractice, tmplt *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pageData := addItem(w, gPractice, r)
		err := tmplt.Execute(w, pageData)
		if err != nil {
			log.Printf("Template error: %v\n", err)
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}
	}
}

func addItem(w http.ResponseWriter, gPractice gpractice.GPractice, r *http.Request) repo.PageData {
	pageData := repo.PageData{}
	var err error

	// validate form
	err = r.ParseForm()
	log.Printf("addItem: form=%v\n", r.Form)
	if err != nil {
		log.Printf("Error parsing form: %v\n", err)
		http.Error(w, "invalid form", http.StatusBadRequest)
		return pageData
	}
	idInput := r.Form.Get("idInput")
	id := 0
	if idInput != "" {
		id, err = strconv.Atoi(idInput)
		if err != nil {
			log.Printf("Error parsing form: %v\n", err)
			http.Error(w, "invalid form", http.StatusBadRequest)
			return pageData
		}
	}
	date := r.Form.Get("dateInput")
	var validDate = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !validDate.MatchString(date) {
		log.Printf("Error parsing form date: %v\n", date)
		http.Error(w, "invalid date", http.StatusBadRequest)
		return pageData
	}
	duration, err := strconv.Atoi(r.Form.Get("durationInput"))
	if err != nil {
		log.Printf("Error parsing form: %v\n", err)
		http.Error(w, "invalid form", http.StatusBadRequest)
		return pageData
	}

	// save item
	item, err := gPractice.Save(repo.Item{Id: uint64(id), Date: date, Duration: uint64(duration)})
	if err != nil {
		log.Printf("Error saving item: %v\n", err)
		http.Error(w, "Failed to save item", http.StatusInternalServerError)
		return pageData
	}
	log.Printf("Item saved: %v\n", item)

	return getAll(w, gPractice)
}

func restGet(gPractice gpractice.GPractice) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pageData := getItem(w, gPractice, r)
		jsonBytes, err := json.Marshal(pageData)
		if err != nil {
			log.Printf("Json error: %v\n", err)
			http.Error(w, "Json error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}
}

func appGet(gPractice gpractice.GPractice, tmplt *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pageData := getItem(w, gPractice, r)
		err := tmplt.Execute(w, pageData)
		if err != nil {
			log.Printf("Template error: %v\n", err)
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}
	}
}

func getItem(w http.ResponseWriter, gPractice gpractice.GPractice, r *http.Request) repo.PageData {
	pageData := repo.PageData{}
	vars := mux.Vars(r)
	var err error

	log.Printf("getItem: vars=%v\n", vars)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Bad id error: %v\n", err)
		http.Error(w, "Bad id", http.StatusBadRequest)
		return pageData
	}
	pageData.Item, err = gPractice.Get(uint64(id))
	if err != nil {
		log.Printf("Failed to retrieve item by id %d, error: %v\n", id, err)
		http.Error(w, "Failed to retrieve item by id", http.StatusInternalServerError)
		return pageData
	}

	return pageData
}

func restUpdate(gPractice gpractice.GPractice) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pageData := addItem(w, gPractice, r)
		jsonBytes, err := json.Marshal(pageData)
		if err != nil {
			log.Printf("Json error: %v\n", err)
			http.Error(w, "Json error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}
}

func appUpdate(gPractice gpractice.GPractice, tmplt *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pageData := addItem(w, gPractice, r)
		err := tmplt.Execute(w, pageData)
		if err != nil {
			log.Printf("Template error: %v\n", err)
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}
	}
}

func restDelete(gPractice gpractice.GPractice) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		deleteItem(w, gPractice, r)
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("{}"))
	}
}

func appDelete(gPractice gpractice.GPractice, tmplt *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		deleteItem(w, gPractice, r)
		w.WriteHeader(http.StatusNoContent)
	}
}
func deleteItem(w http.ResponseWriter, gPractice gpractice.GPractice, r *http.Request) {
	vars := mux.Vars(r)
	var err error

	log.Printf("delete: vars=%v\n", vars)

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
}
