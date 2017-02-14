package controllers

import "net/http"
import "github.com/labstack/echo"
import "github.com/deyvisonguilherme/viajar_brasil/models"

// GetCategorias retorna uma lista de categorias de locais
func GetCategorias(c echo.Context) error {
	categorias, err := models.ListCategoria()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, categorias)
}
