package main

import (
	"fmt"
	"olin-test/app/router"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	router.Init(app)

	app.Start(":8080")

	fmt.Println("Hello, World!")
}
