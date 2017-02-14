package controllers

import "net/http"
import "github.com/labstack/echo"
import "github.com/deyvisonguilherme/viajar_brasil/models"

// GetRegioes retorna uma lista de regioes do país
func GetRegioes(c echo.Context) error {
	regioes, err := models.ListRegioes()

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, regioes)
}
