package router

import (
	"olin-test/app/repository"

	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo) {
	api := app.Group("/v1")
	{
		api.GET("/test1", repository.TwoSumOutput)
		api.GET("/test2", repository.ThreeSumOutput)
		api.GET("/test3", repository.Test4)
	}
}
