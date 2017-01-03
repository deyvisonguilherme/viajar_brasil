package main

import (
	"fmt"
	"github.com/deyvison/senem"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Feliz Ano Novo!")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "hello, %s!\n", ps.ByName("name"))

}

func GetListCaderno(w http.ResponseWriter, r *http.Request, ps httrouter.Params) {
	var caderno Caderno
	db.First(&caderno)
}

func GetCadernoById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var caderno Caderno
	db.First(&caderno, ps.Cd_Caderno)
}

func AddCaderno(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db.Create(&Caderno{})
}

func UpCaderno() {
	var caderno Caderno
	db.Model(&caderno).Update("params", "value")
}

func DelCaderno(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var caderno Caderno
	db.Delete(&caderno)
}
