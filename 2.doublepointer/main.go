package main

import "strings"

func main() {}

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
