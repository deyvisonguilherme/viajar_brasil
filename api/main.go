package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.POST("/caderno/add/:params", parametros)
	router.GET

	log.Fatal(http.ListenAndServe(":8080", router))
}
