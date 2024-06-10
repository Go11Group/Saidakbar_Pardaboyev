package handler

import (
	"encoding/json"
	"leetcode/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Read
func (h *Handler) GetLanguages(w http.ResponseWriter, r *http.Request) {
	filter := model.FilterLanguage{}
	query := r.URL.Query()
	if query.Has("name") {
		name := query.Get("name")
		filter.Name = &name
	}

	languages, err := h.Languages.GetLanguages(&filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting languages", err)
		return
	}

	err = json.NewEncoder(w).Encode(languages)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while encoding languages", err)
		return
	}

}

func (h *Handler) GetLanguageByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	language, err := h.Languages.GetLanguageById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting language by Id", err)
		return
	}
	err = json.NewEncoder(w).Encode(language)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while encoding language", err)
		return
	}
}

// Create
func (h *Handler) CreateLanguages(w http.ResponseWriter, r *http.Request) {
	newlanguage := model.Language{}
	err := json.NewDecoder(r.Body).Decode(&newlanguage)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while decoding language", err)
		return
	}
	err = h.Languages.CreateLanguage(&newlanguage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while creating language", err)
		return
	}
}

// Update
func (h *Handler) UpdateLanguage(w http.ResponseWriter, r *http.Request) {
	language := model.Language{}
	vars := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&language)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while decoding language", err)
		return
	}

	language.Id = vars["id"]
	err = h.Languages.UpdateLanguage(&language)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while updating language", err)
		return
	}
}

// Delete
func (h *Handler) DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := h.Languages.DeleteLanguage(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while deleting language", err)
		return
	}
}
