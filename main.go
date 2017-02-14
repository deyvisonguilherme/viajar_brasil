package main

import "github.com/labstack/echo"
import "github.com/labstack/echo/middleware"
import "github.com/deyvisonguilherme/viajar_brasil/handlers"

func main() {
	// Echo Instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Router => handler
	e.GET("/", handlers.Saudacao)
	e.GET("/regioes", handlers.Regioes)
	e.GET("/categorias", handlers.Categorias)
	e.GET("/locais", handlers.Locais)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
