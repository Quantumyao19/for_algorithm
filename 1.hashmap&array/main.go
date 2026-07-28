package main

import (
	"bytes"
	"fmt"
	"slices"
)

func main() {
	f := ValidAnagram2("asdfga", "gdfsaa")
	fmt.Println(f)
}

// Day 3: Valid Anagram
func ValidAnagram(s, t string) bool {
	sbyte := []byte(s)
	slices.Sort(sbyte)
	tbyte := []byte(t)
	slices.Sort(tbyte)

	return bytes.Equal(sbyte, tbyte)
}
func ValidAnagram2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		count[s[i]]++
		count[t[i]]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}

// Day 2: Contains Duplicate
func ContainsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = struct{}{}
	}
	return false
}

// Day 1: TwoSum
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{i, j}
		}
		m[num] = i
	}
	return nil
}
