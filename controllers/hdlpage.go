package controllers

import (
	"strconv"

	"github.com/deyvisonguilherme/viajar_brasil/models"
	"github.com/labstack/echo"
)

// Votar : grava a votação do usuário.
func Votar(c echo.Context) error {
	usuario, err := strconv.Atoi(c.Param("usuario"))
	local, err := strconv.Atoi(c.Param("local"))
	votacao, err := strconv.Atoi(c.Param("votacao"))
	models.Votar(usuario, local, votacao)
}

// Comentar : grava comentário do usuário.
func Comentar(c echo.Context) error {

}

// Page : carrega componentes de páginas.
func Page(c echo.Context) error {

}
