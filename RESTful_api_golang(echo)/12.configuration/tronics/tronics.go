package tronics

import (
	"fmt"
	"os"
)

var e = echo.New()
var v = validator.New()

// start the application
func Start() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}

	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.DELETE("/products/:id", deleteProduct)
	e.PUT("/products/:id", updateProduct)
	e.POST("/products", createProduct)

	e.Logger.Print(fmt.Sprintf("Listening on port %s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
