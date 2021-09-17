package main

import (
	"fmt"
	"github.com/labstack/echo"
)

const ServerAddress string = "localhost:8080"

func main() {
	e := echo.New()
	e.Static("/", "server/static")
	if err := e.Start(ServerAddress); err != nil {
		fmt.Println(err)
	}
}
