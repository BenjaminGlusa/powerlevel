package handler

import (
	"github.com/BenjaminGlusa/powerlevel/pkg/adapter"
	"github.com/gorilla/mux"

)

func NewRouter(db adapter.DatabaseAdapter) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/measure", MakeAddMeasurementHandler(db)).Methods("POST")
	return r
}
