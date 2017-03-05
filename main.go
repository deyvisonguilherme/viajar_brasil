package main

import (
	"log"
	"os"

	"net/http"

	"github.com/deyvisonguilherme/viajar_brasil/controllers"
	"github.com/deyvisonguilherme/viajar_brasil/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	// Conexão com banco de dados
	models.InitDB("postgres://developer:colleto@localhost/viajar_brasil?sslmode=disable")
	log.SetOutput(os.Stdout)
}

func main() {
	// Echo Instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Router => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Olá benvindo a API viaja Brasil!")
	})

	// Index
	e.GET("/locais/:regiao/:categoria", controllers.GetLocais)
	e.GET("/pesquisa", controllers.Pesquisar)

	// Pages
	e.GET("/page", controllers.Page)
	e.GET("/page/votar/:usuario/:local/:votacao", controllers.Votar)
	e.GET("/page/comentar/:usuarios/:local/:comentario", controllers.Comentar)

	// Usuários
	e.GET("/user/logar", controllers.Logar)
	e.GET("/user/cadastro", controllers.Cadastro)
	e.GET("/user/logout", controllers.Logout)
	e.GET("/user/dadosuser", controllers.DadosUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
