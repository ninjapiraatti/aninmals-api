package router

import (
	"aninmals-api/pkg/handler"

	"github.com/gorilla/mux"
)

func NewRouter(aninmalHandler handler.AninmalHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/aninmals", aninmalHandler.GetAninmals).Methods("GET")

	return r
}
