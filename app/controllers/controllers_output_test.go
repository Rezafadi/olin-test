package controllers

import (
	"olin-test/app/repository"
	"olin-test/app/reqres"

	"github.com/labstack/echo/v4"
)

func TwoSum(c echo.Context) error {
	var response reqres.OutputTask1

	nums := []int{1, 2, 3, 4, 5}
	target := 9
	data := repository.TwoSum(nums, target)
	response.Output = data

	return c.JSON(200, map[string]interface{}{
		"success": 200,
		"data":    response,
		"message": "success",
	})
}
