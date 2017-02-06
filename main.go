package main

import (
	"github.com/deyvisonguilherme/viajar_brasil/models"
	"net/http"
	"encoding/json"
	"log"
)

func main() {
	models.InitDB("postgres://developer:colleto@localhost/viajar_brasil?sslmode=disable")

	http.HandleFunc("/categorias", categoriasList)
	http.HandleFunc("/regioes", regioesList)
	log.Fatal(http.ListenAndServe(":3000", nil))
}


func categoriasList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	categorias, err := models.GetCategorias()

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	
	js, err := json.Marshal(categorias)
	 
	 if err != nil {
	 	http.Error(w, err.Error(), http.StatusInternalServerError)
	 	return
	 }

	 w.Header().Set("Content-Type", "application/json")
	 w.Write(js)
}

func regioesList(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w, http.StatusText(405), 405)
		return
	}
	
	regioes, err := models.GetRegioes()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	
	js, err := json.Marshal(regioes)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func locaisList(w http.ResponseWrite, r *http.Request){
	
}
