package main

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}
	e := echo.New()
	v := validator.New()
	products := []map[int]string{{1: "mo"}, {2: "t"}, {3: "lap"}}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})
	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, products)
	})

	e.POST("/products", func(c echo.Context) error {
		type body struct {
			Name string `json:"product_name" validate:"required, min=4"`
		}
		var reqBody body
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		if err := v.Struct(&reqBody); err != nil {
			return err
		}

		product := map[int]string{
			len(products) + 1: reqBody.Name,
		}
		return c.JSON(http.StatusOK, product)
	})

	e.PUT("/products/:id", func(c echo.Context) error {
		var product map[int]string
		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		for _, p := range products {
			for k := range p {
				if pID == k {
					product = p
				}
			}
		}

		if product == nil {
			return c.JSON(http.StatusNotFound, "product not found")
		}
		return c.JSON(http.StatusNotFound, product)
	})

	e.DELETE("/products/:id", func(c echo.Context) error {
		var product map[int]string
		var index int
		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		for i, p := range products {
			for k := range p {
				if pID == k {
					product = p
					index = i
				}
			}
		}

		if product == nil {
			return c.JSON(http.StatusNotFound, "product not found")
		}

		splice := func(s []map[int]string, index int) []map[int]string {
			//단축된 슬라이스 또는 접합된 슬라이스를 얻어 반환하고자하는 값을 가져온다.
			return append(s[index:], s[:index+1]...)
		}
		products = splice(products, index)
		//인덱스를 제공하고 이 특정 제품에서 내 제품을 가져옴. 즉 인덱스가 제거됨.

		return c.JSON(http.StatusNotFound, product)
	})

}
