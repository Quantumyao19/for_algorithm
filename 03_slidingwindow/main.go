package main

import (
	"fmt"
	"slices"
)

func main() {
	var count1, count2 [5]int
	fmt.Println(count1 == count2)
}

// Day 1: Longest Substring Without Repeating Characters
func LengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	left := 0
	res := 0
	for right := 0; right < len(s); right++ {
		if index, ok := m[s[right]]; ok && index >= left {
			left = index + 1
		}

		m[s[right]] = right
		res = max(res, right-left+1)

	}
	return res
}

// Day 2: Permutation in String
func PermutationInString(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	left := 0
	s1byte := []byte(s1)
	slices.Sort(s1byte)
	for right := len(s1) - 1; right < len(s2); right++ {
		s2byte := []byte(s2[left : right+1])
		slices.Sort(s2byte)
		if string(s1byte) == string(s2byte) {
			return true
		}
		left++
	}
	return false
}

func PermutationInString2(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var count1, count2 [26]int

	for i := range s1 {
		count1[s1[i]-'a']++
		count2[s2[i]-'a']++
	}

	if count1 == count2 {
		return true
	}

	for right := len(s1); right < len(s2); right++ {
		count2[s2[right]-'a']++
		count2[s2[right-len(s1)]-'a']--

		if count1 == count2 {
			return true
		}
	}
	return false
}

// Day 3: Longest Repeating Character Replacement
func LongeestRepeatingCharacterReplcament(s string, k int) int {
	left := 0
	var count [26]int
	maxFreq := 0
	result := 0
	for right := 0; right < len(s); right++ {
		index := s[right] - 'A'
		count[index]++

		if count[index] > maxFreq {
			maxFreq = count[index]
		}

		windowSize := right - left + 1
		for windowSize-maxFreq > k {
			count[s[left]-'A']--
			left++

			windowSize = right - left + 1
		}
		if windowSize > result {
			result = windowSize
		}
	}
	return result
}

// Day 4: Minimum Window Substring
func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	window := make(map[byte]int)
	left := 0
	valid := 0
	result := ""

	for right := 0; right < len(s); right++ {
		c := s[right]

		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if result == "" || right-left+1 < len(result) {
				result = s[left : right+1]
			}

			leftChar := s[left]
			left++

			if _, ok := need[leftChar]; ok {
				if window[leftChar] == need[leftChar] {
					valid--
				}
				window[leftChar]--
			}
		}
	}
	return result
}
