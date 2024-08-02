package controllers

import (
	"olin-test/app/repository"
	"olin-test/app/reqres"

	"github.com/labstack/echo/v4"
)

func TwoSumOutput(c echo.Context) error {
	var response reqres.OutputTask1

	nums := []int{1, 2, 3, 4, 5}
	target := 9
	data := repository.TwoSum(nums, target)
	response.Output = data
	response.Number = nums
	response.Target = target

	return c.JSON(200, map[string]interface{}{
		"success": 200,
		"data":    response,
		"message": "success",
	})
}

func ThreeSumOutput(c echo.Context) error {
	var response reqres.OutputTask2

	nums1 := []int{-1, 0, 1, 2, -1, -4}
	data := repository.ThreeSum(nums1)
	response.Output = data
	response.Number = nums1

	return c.JSON(200, map[string]interface{}{
		"success": 200,
		"data":    response,
		"message": "success",
	})
}

func Test4(c echo.Context) error {
	var response reqres.OutputTask3

	s := "barfoofoobarthefoobarman"
	words := []string{"bar", "foo", "the"}

	data := repository.FindSubstring(s, words)
	response.Output = data
	response.String = s
	response.Words = words

	return c.JSON(200, map[string]interface{}{
		"success": 200,
		"data":    response,
		"message": "success",
	})
}
