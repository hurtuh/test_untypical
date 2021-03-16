package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"test_untypical/internal/test_untypical/service"
	"test_untypical/internal/test_untypical/types"
	"test_untypical/pkg/infrastruct"
)

type Handlers struct {
	srv service.ServiceI
}

func NewHandlers(srv service.ServiceI) *Handlers {
	return &Handlers{srv: srv}
}

func (h *Handlers) Get(w http.ResponseWriter, r *http.Request) {

	key := r.FormValue("key")
	value, err := h.srv.Get(key)
	if err != nil {
		apiErrorEncode(w, err)
		return
	}

	apiResponseEncoder(w, types.KeyValue{Value: value})
}

func (h *Handlers) Upsert(w http.ResponseWriter, r *http.Request) {

	keyValue := types.KeyValue{}
	if err := json.NewDecoder(r.Body).Decode(&keyValue); err != nil {
		apiErrorEncode(w, infrastruct.ErrorBadRequest)
		return
	}

	err := h.srv.Upsert(&keyValue)
	if err != nil {
		apiErrorEncode(w, err)
		return
	}
}

func (h *Handlers) Delete(w http.ResponseWriter, r *http.Request) {

	key := r.FormValue("key")
	if err := h.srv.Delete(key); err != nil {
		apiErrorEncode(w, err)
		return
	}
}

func (h *Handlers) List(w http.ResponseWriter, _ *http.Request) {

	list, err := h.srv.List()
	if err != nil {
		apiErrorEncode(w, err)
		return
	}

	apiResponseEncoder(w, list)
}

//Пишет в респонс код ошибки и ошибку в формате json
func apiErrorEncode(w http.ResponseWriter, err error) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if customError, ok := err.(*infrastruct.CustomErrors); ok {
		w.WriteHeader(customError.Code)
	}

	result := struct {
		Err string `json:"error"`
	}{
		Err: err.Error(),
	}

	if err = json.NewEncoder(w).Encode(result); err != nil {
		log.Println(err)
		return
	}
}

//Пишет в респонс в формате json
func apiResponseEncoder(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		return
	}
}
