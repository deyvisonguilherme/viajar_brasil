package controllers

import "net/http"
import "github.com/labstack/echo"
import "github.com/deyvisonguilherme/viajar_brasil/models"

// Pesquisar : lista um conjunto de categorias e regiões para executar uma consulta.
func Pesquisar(c echo.Context) error {
	pesquisa, err := models.Pesquisa()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, pesquisa)
}
