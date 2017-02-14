package handlers

import "net/http"
import "github.com/labstack/echo"

// import "github.com/deyvisonguilherme/viajar_brasil/models"

// Regioes retorna uma lista de regioes do país
func Regioes(c echo.Context) error {
	return c.String(http.StatusOK, "Regioes")
}
