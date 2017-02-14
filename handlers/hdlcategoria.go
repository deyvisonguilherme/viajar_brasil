package handlers

import "net/http"
import "github.com/labstack/echo"

// Categorias retorna uma lista de categorias de locais
func Categorias(c echo.Context) error {
	return c.String(http.StatusOK, "Categorias")
}
