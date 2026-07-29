package main

import (
	"bytes"
	"fmt"
	"slices"
	"sort"
)

func main() {
	f := ValidAnagram2("asdfga", "gdfsaa")
	fmt.Println(f)
}

// Day 5: Top K Frequent Elements
func TopKFrequentElements(nums []int, k int) []int {
	var res []int
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	type Pair struct {
		num  int
		freq int
	}

	pairs := []Pair{}
	for num, freq := range m {
		pairs = append(pairs, Pair{
			num:  num,
			freq: freq,
		})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})

	for i := 0; i < k; i++ {
		res = append(res, pairs[i].num)
	}

	return res
}

// Day 4: Group Anagrams
func GroupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int
		for _, s := range str {
			count[s-'a']++
		}
		m[count] = append(m[count], str)
	}

	result := make([][]string, 0)
	for _, v := range m {
		result = append(result, v)
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
