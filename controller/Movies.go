package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	_"fmt"
)

type Movie struct{
	Router *mux.Router `inject:""`
}

func (m *Movie) Get(w http.ResponseWriter , r *http.Request){
	w.Write([]byte("The list of the movies"))
}

func (m *Movie) Register(){
	m.Router.HandleFunc("/movies",m.Get).Methods("GET")
}