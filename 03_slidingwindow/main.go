package main

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
