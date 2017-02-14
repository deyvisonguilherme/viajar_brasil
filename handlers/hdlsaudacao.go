package handlers

import "net/http"
import "github.com/labstack/echo"

// Saudacao retorna uma Saudação para o cliente.
func Saudacao(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World!\n")
}
