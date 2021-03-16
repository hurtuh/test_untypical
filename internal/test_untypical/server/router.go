package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"test_untypical/internal/test_untypical/server/handlers"
)

func NewRouter(h *handlers.Handlers) *mux.Router {
	router := mux.NewRouter()

	router.Methods(http.MethodGet).Path("/get").HandlerFunc(h.Get)
	router.Methods(http.MethodGet).Path("/list").HandlerFunc(h.List)
	router.Methods(http.MethodPost).Path("/upsert").HandlerFunc(h.Upsert)
	router.Methods(http.MethodDelete).Path("/delete").HandlerFunc(h.Delete)

	return router
}
