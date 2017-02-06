package test

import "testing"
import "github.com/deyvisonguilherme/viajar_brasil/models"
import "fmt"

func TestGetCategorias(t *testing.T){
	negocio, err := models.GetCategorias()
	
	if err != nil {
        fmt.Printf("%s", err.Error())
    }
	fmt.Printf("%s", negocio)
} 