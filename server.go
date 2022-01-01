package main

import (
	"fmt"
	"mygraphql/config"
	"mygraphql/graphql"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDB()
	route := echo.New()

	h, err := graphql.NewHandler(config.DB)
	if err != nil {
		fmt.Println(err)
	}
	route.POST("/graphql", echo.WrapHandler(h))
	route.Start(":9000")
}

type Response struct {
	ErrorCode int         `json:"error_code" form:"error_code"`
	Message   string      `json:"message" form:"message"`
	Data      interface{} `json:"data"`
}
