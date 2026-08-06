package main

func main() {}

// Day 1: Valid Parentheses
func ValidParentheses(s string) bool {
	pairs := map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
	}

	stack := []rune{}

	for _, ch := range s {
		if ch == '{' || ch == '[' || ch == '(' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if top != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// Day 2: Min Stack
func MinStack() {}
