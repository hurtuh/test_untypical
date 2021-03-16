package server

import (
	"log"
	"net/http"
	"test_untypical/internal/test_untypical/server/handlers"
)

func StartServer(h *handlers.Handlers, port string) {
	router := NewRouter(h)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("err with ListAndServe: ", err)
	}
}
