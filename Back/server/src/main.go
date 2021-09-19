package main

import (
	"github.com/labstack/echo"
	"server/server/src/config"
	"server/server/src/controller"
	"server/server/src/service"

	"fmt"
)

const ServerAddress string = "localhost:8080"

func main() {
	e := echo.New()
	config.InitLogger(e)
	e.Static("/", "server/static")
	e.GET("/statistics", controller.GetStatistics)

	if err := config.InitDB(); err != nil {
		fmt.Println(err)
		return
	}
	service.RunServices()
	if err := e.Start(ServerAddress); err != nil {
		fmt.Println(err)
		return
	}
}
