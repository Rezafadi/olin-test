package router

import (
	"olin-test/app/controllers"

	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo) {
	api := app.Group("/v1")
	{
		api.GET("/test1", controllers.TwoSumOutput)
		api.GET("/test2", controllers.ThreeSumOutput)
		api.GET("/test3", controllers.Test4)
	}
}
