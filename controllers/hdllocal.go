package controllers

import "net/http"
import "github.com/labstack/echo"
import "github.com/deyvisonguilherme/viajar_brasil/models"
import "strconv"

// GetLocais retorna os locais buscados pelo usuário apartir de paramêtros especifícados
func GetLocais(c echo.Context) error {
	regiao, err := strconv.Atoi(c.Param("regiao"))
	categoria, err := strconv.Atoi(c.Param("categoria"))

	locais, err := models.ListLocais(regiao, categoria)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, locais)
}
