package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"aninmals-api/pkg/repository"

	"github.com/gorilla/mux"
)

type AninmalHandler struct {
	AninmalRepository repository.AninmalRepository
}

func (h AninmalHandler) GetAninmals(w http.ResponseWriter, r *http.Request) {
	aninmals, err := h.AninmalRepository.GetAninmals()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(aninmals)
}

func (h AninmalHandler) GetAninmalById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	aninmal, err := h.AninmalRepository.GetAninmalById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(aninmal)
}
