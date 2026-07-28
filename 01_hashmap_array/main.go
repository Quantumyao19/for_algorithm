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

// Day 4: Group Anagrams
func GroupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}

		groups[count] = append(groups[count], s)
	}

	result := make([][]string, 0)
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}
func GroupAnagrams2(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		sbyte := []byte(str)
		slices.Sort(sbyte)

		key := string(sbyte)
		m[key] = append(m[key], str)
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}

	return result
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
