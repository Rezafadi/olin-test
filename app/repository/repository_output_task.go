package repository

import (
	"sort"
)

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if index, found := numMap[complement]; found {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return []int{}
}

// func TwoSumOutput(c echo.Context) error {
// 	var response reqres.OutputTask1

// 	nums := []int{1, 2, 3, 4, 5}
// 	target := 9
// 	data := TwoSum(nums, target)
// 	response.Output = data
// 	response.Number = nums
// 	response.Target = target

// 	return c.JSON(200, map[string]interface{}{
// 		"success": 200,
// 		"data":    response,
// 		"message": "success",
// 	})
// }

func ThreeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // Skip duplicate elements
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++ // Skip duplicate elements
				}
				for left < right && nums[right] == nums[right-1] {
					right-- // Skip duplicate elements
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

// func ThreeSumOutput(c echo.Context) error {
// 	var response reqres.OutputTask2

// 	nums1 := []int{-1, 0, 1, 2, -1, -4}
// 	data := threeSum(nums1)
// 	response.Output = data
// 	response.Number = nums1

// 	return c.JSON(200, map[string]interface{}{
// 		"success": 200,
// 		"data":    response,
// 		"message": "success",
// 	})
// }

func FindSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	wordLength := len(words[0])
	numWords := len(words)
	totalLength := wordLength * numWords

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	result := []int{}

	for i := 0; i <= len(s)-totalLength; i++ {
		seen := make(map[string]int)
		for j := 0; j < numWords; j++ {
			wordIndex := i + j*wordLength
			word := s[wordIndex : wordIndex+wordLength]
			if _, ok := wordCount[word]; ok {
				seen[word]++
				if seen[word] > wordCount[word] {
					break
				}
			} else {
				break
			}
		}
		if CompareMaps(seen, wordCount) {
			result = append(result, i)
		}
	}

	return result
}

func CompareMaps(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for key, val := range a {
		if b[key] != val {
			return false
		}
	}
	return true
}

// func Test4(c echo.Context) error {
// 	var response reqres.OutputTask3

// 	s := "barfoofoobarthefoobarman"
// 	words := []string{"bar", "foo", "the"}

// 	data := findSubstring(s, words)
// 	response.Output = data
// 	response.String = s
// 	response.Words = words

// 	return c.JSON(200, map[string]interface{}{
// 		"success": 200,
// 		"data":    response,
// 		"message": "success",
// 	})
// }
