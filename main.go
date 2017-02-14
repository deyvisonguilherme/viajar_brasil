package main

import (
	"log"
	"os"

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
	e.GET("/", controllers.Saudacao)
	e.GET("/regioes", controllers.GetRegioes)
	e.GET("/categorias", controllers.GetCategorias)
	e.GET("/locais/:regiao/:categoria", controllers.GetLocais)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
