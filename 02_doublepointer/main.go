package main

import (
	"sort"
	"strings"
)

func main() {}

// Day 4: 3Sum
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		start := nums[i]
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := start + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{start, nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
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
	return res
}

// Day 3: Container With Most Water
func ContainerWithMostWater(nums []int) int {
	right := len(nums) - 1
	left := 0
	maxArea := 0
	for left < right {
		var area int
		area = (right - left) * min(nums[left], nums[right])
		if area > maxArea {
			maxArea = area
		}

		if nums[left] < nums[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

// Day 2: Two Sum ||
func TwoSum(nums []int, target int) []int {
	right := len(nums) - 1
	left := 0
	for left < right {
		var sum = nums[left] + nums[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

// Day 1: Valid Palindrome
func ValidPalindrome(str string) bool {
	left, right := 0, len(str)-1

	for left < right {
		for left < right && !isAlphaNum(str[left]) {
			left++
		}

		for left < right && !isAlphaNum(str[right]) {
			right--
		}

		if !strings.EqualFold(string(str[left]), string(str[right])) {
			return false
		}

		left++
		right--
	}
	return true
}

func isAlphaNum(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}
