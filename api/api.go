package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lbragadev/engineering-assessment/store"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

// Decorates an apiFunc to http.HandlerFunc.
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
	store      store.Storage
}

func NewAPIServer(listenAddr string, store store.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/foodtrucks", makeHTTPHandleFunc(s.handleGetFoodTrucks))
	log.Println("JSON API server running on port : ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleGetFoodTrucks(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		foodTrucks, err := s.store.GetFoodTrucks()
		if err != nil {
			return err
		}

		return WriteJSON(w, http.StatusOK, foodTrucks)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}
