package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)


func WriteJSON (w http.ResponseWriter, status int, v any) error{
	w.WriteHeader(status)
	w.Header().Set("Content-Type","application/json")
	return json.NewEncoder(w).Encode(v) 
}
 func makeHTTPHandleFunc(f apiFunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r*http.Request){
		if err := f(w,r); err != nil{
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
 }
type apiFunc func (http.ResponseWriter, *http.Request) error  
	

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{listenAddr: listenAddr}
}

func (s *APIServer) Run(){
	router := mux.NewRouter()

	router.HandleFunc("/account",makeHTTPHandleFunc(s.HandleAccount) )
}

func (s *APIServer) HandleAccount(w http.ResponseWriter , r *http.Request) error {
	return nil

}
func (s *APIServer) HandleDeleteAccount(w http.ResponseWriter , r *http.Request) error {
	return nil

}
func (s *APIServer) HandleCreateAccount(w http.ResponseWriter , r *http.Request) error {
	return nil

}

