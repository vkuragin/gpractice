package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vkuragin/gpractice"
	"github.com/vkuragin/gpractice/repo"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type restHandler struct {
	gp gpractice.GPractice
}

func (h *restHandler) restAll() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pageData := getAll(w, h.gp)
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

func (h *restHandler) restAdd() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		item, ok := parseJson(w, r)
		if !ok {
			return
		}
		item = addItem(w, h.gp, item)
		jsonBytes, err := json.Marshal(item)
		if err != nil {
			log.Printf("Json error: %v\n", err)
			http.Error(w, "Json error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}
}

func (h *restHandler) restGet() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		item := getItem(w, h.gp, r)
		jsonBytes, err := json.Marshal(item)
		if err != nil {
			log.Printf("Json error: %v\n", err)
			http.Error(w, "Json error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}
}

func (h *restHandler) restUpdate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		item, ok := parseJson(w, r)
		if !ok {
			return
		}
		item = addItem(w, h.gp, item)
		jsonBytes, err := json.Marshal(item)
		if err != nil {
			log.Printf("Json error: %v\n", err)
			http.Error(w, "Json error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}
}

func (h *restHandler) restDelete() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		deleteItem(w, h.gp, r)
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("{}"))
	}
}

type appHandler struct {
	gp     gpractice.GPractice
	holder tplHolder
}

func (h *appHandler) appAll() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pageData := getAll(w, h.gp)
		tpl, err := h.holder.getTemplate()
		if err != nil {
			log.Printf("Template error: %v\n", err)
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}
		err = tpl.Execute(w, pageData)
		if err != nil {
			log.Printf("Template error: %v\n", err)
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}
	}
}

func (h *appHandler) appAdd() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		item, ok := parseForm(w, r)
		if !ok {
			return
		}
		item = addItem(w, h.gp, item)
		pageData := getAll(w, h.gp)

		tpl, err := h.holder.getTemplate()
		if err != nil {
			log.Printf("Template error: %v\n", err)
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}
		err = tpl.Execute(w, pageData)
		if err != nil {
			log.Printf("Template error: %v\n", err)
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}
	}
}

func (h *appHandler) appGet() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pageData := repo.PageData{}
		item := getItem(w, h.gp, r)
		pageData.Item = item

		tpl, err := h.holder.getTemplate()
		if err != nil {
			log.Printf("Template error: %v\n", err)
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}
		err = tpl.Execute(w, pageData)
		if err != nil {
			log.Printf("Template error: %v\n", err)
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}
	}
}

func (h *appHandler) appUpdate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		item, ok := parseForm(w, r)
		if !ok {
			return
		}
		item = addItem(w, h.gp, item)
		pageData := getAll(w, h.gp)

		tpl, err := h.holder.getTemplate()
		if err != nil {
			log.Printf("Template error: %v\n", err)
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}
		err = tpl.Execute(w, pageData)
		if err != nil {
			log.Printf("Template error: %v\n", err)
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}
	}
}

func (h *appHandler) appDelete() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		deleteItem(w, h.gp, r)
		w.WriteHeader(http.StatusNoContent)
	}
}

// ---- helper functions ----//

func getAll(w http.ResponseWriter, gPractice gpractice.GPractice) repo.PageData {
	item := repo.ItemDto{Date: time.Now().Format(repo.DateFormat)}
	pageData := repo.PageData{Item: item}
	var err error

	log.Printf("getAll\n")

	report, err := gPractice.GetReport()
	if err != nil {
		log.Printf("Error: %v\n", err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return pageData
	}

	dtos := make([]repo.ItemDto, len(report.Items))
	for i, item := range report.Items {
		dtos[i] = repo.ItemToDto(item)
	}
	pageData.Items = dtos
	pageData.Report = repo.ReportToDto(report)
	return pageData
}

func parseJson(w http.ResponseWriter, r *http.Request) (repo.ItemDto, bool) {
	item := repo.ItemDto{}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Json error: %v\n", err)
		http.Error(w, "Json error", http.StatusInternalServerError)
		return item, false
	}
	err = json.Unmarshal(bytes, &item)
	if err != nil {
		log.Printf("Json error: %v\n", err)
		http.Error(w, "Json error", http.StatusBadRequest)
		return item, false
	}

	var validDate = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !validDate.MatchString(item.Date) {
		log.Printf("Error parsing form date: %v\n", item.Date)
		http.Error(w, "invalid date", http.StatusBadRequest)
		return item, false
	}
	return item, true
}

func parseForm(w http.ResponseWriter, r *http.Request) (repo.ItemDto, bool) {
	var err error
	var dto repo.ItemDto

	err = r.ParseForm()
	if err != nil {
		log.Printf("Error parsing form: %v\n", err)
		http.Error(w, "invalid form", http.StatusBadRequest)
		return dto, false
	}

	idInput := r.Form.Get("idInput")
	id := 0
	if idInput != "" {
		id, err = strconv.Atoi(idInput)
		if err != nil {
			log.Printf("Error parsing form: %v\n", err)
			http.Error(w, "invalid form", http.StatusBadRequest)
			return dto, false
		}
	}
	date := r.Form.Get("dateInput")
	var validDate = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !validDate.MatchString(date) {
		log.Printf("Error parsing form date: %v\n", date)
		http.Error(w, "invalid date", http.StatusBadRequest)
		return dto, false
	}
	duration, err := time.ParseDuration(r.Form.Get("durationInput"))
	if err != nil {
		log.Printf("Error parsing form: %v\n", err)
		http.Error(w, "invalid form", http.StatusBadRequest)
		return dto, false
	}

	return repo.ItemDto{
		int(id),
		date,
		int(duration.Seconds()),
		r.Form.Get("durationInput"),
	}, true
}

func addItem(w http.ResponseWriter, gPractice gpractice.GPractice, dto repo.ItemDto) repo.ItemDto {
	// save item
	item := repo.DtoToItem(dto)
	item, err := gPractice.Save(item)
	if err != nil {
		log.Printf("Error saving item: %v\n", err)
		http.Error(w, "Failed to save item", http.StatusInternalServerError)
		return repo.ItemDto{}
	}
	log.Printf("Item saved: %v\n", item)

	return repo.ItemToDto(item)
}

func getItem(w http.ResponseWriter, gPractice gpractice.GPractice, r *http.Request) repo.ItemDto {
	dto := repo.ItemDto{}
	vars := mux.Vars(r)
	var err error

	log.Printf("getItem: vars=%v\n", vars)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Bad id error: %v\n", err)
		http.Error(w, "Bad id", http.StatusBadRequest)
		return dto
	}
	item, err := gPractice.Get(id)
	if err != nil {
		log.Printf("Failed to retrieve item by id %d, error: %v\n", id, err)
		http.Error(w, "Failed to retrieve item by id", http.StatusInternalServerError)
		return dto
	}

	return repo.ItemToDto(item)
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
	result, err = gPractice.Delete(id)
	if err != nil {
		log.Printf("Error deleting item: %v\n", err)
		http.Error(w, "Failed to delete item", http.StatusInternalServerError)
		return
	}
	log.Printf("Delete result: %v\n", result)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
