package handlers

import "net/http"
import "github.com/labstack/echo"

// Locais retorna os locais buscados pelo usuário apartir de paramêtros especifícados
func Locais(c echo.Context) error {
	return c.String(http.StatusOK, "Locais")
}
