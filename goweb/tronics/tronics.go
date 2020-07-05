package tronics

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

var e = echo.New()
var v = validator.New()

func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load configuration")
	}
}

func serverMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside middleware")
		c.Request().URL.Path = "/krunal"
		fmt.Printf("%+v\n", c.Request())
		return next(c)
	}
}

// type MiddlewareFunc func(HandlerFunc) HandlerFunc
// type HandlerFunc func(Context) error

//Start starts the application
func Start() {
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/products", getProducts)
	e.GET("products/:id", getProduct)
	e.DELETE("products/:id", deleteProduct)
	e.PUT("/product/:id", updateProduct)
	e.POST("/products", createProduct, middleware.BodyLimit("1K"))

	e.Logger.Print(fmt.Sprintf("Listening on port %s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", cfg.Port)))
}