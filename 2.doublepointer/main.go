package main

import (
	"strings"
)

func main() {}

// Day 3: Container With Most Water
func ContainerWithMostWater(nums []int) int {
	right := len(nums) - 1
	left := 0
	maxArea := 0
	var area int
	for left < right {
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
