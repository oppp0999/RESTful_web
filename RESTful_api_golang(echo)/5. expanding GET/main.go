package main

import (
	"net/http"
	"os"
)

func main(){
	port := os.Getenv("MY_APP_PORT")
	if port == ""{
		port = "8080"
	}
	e:=echo.New()
	products := []map[int]string{{1: "m"}, {2:"t"},{3:"l"}}

	e.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, "well, hello")
	})
	e.GET("/prodicts/id", func(c echo.Context) error{
		var product map[int]string
		
	})
}