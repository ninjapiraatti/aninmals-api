package handler

import (
	"encoding/json"
	"net/http"

	"aninmals-api/pkg/repository"
)

type AninmalHandler struct {
	AninmalRepository repository.AninmalRepository
}

func (h AninmalHandler) GetAninmals(w http.ResponseWriter, r *http.Request) {
	aninmals := h.AninmalRepository.GetAninmals()

	json.NewEncoder(w).Encode(aninmals)
}
