package tronics

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var e = echo.New()
var v = validator.New()

//Start starts the application
func Start() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}	

	e.GET("/products", getProducts)
	e.GET("products/:id", getProducts)
	e.DELETE("products/:id", deleteProduct)
	e.PUT("/product/:id", updateProduct)
	e.POST("/products", createProduct)

	e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}