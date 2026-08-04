package main

import "slices"

func main() {}

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

	var count1 [26]int
	var count2 [26]int

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
