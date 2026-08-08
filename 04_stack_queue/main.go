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
type MinStack struct {
	Stack    []int
	MinStack []int
}

func Constructor() MinStack {
	return MinStack{
		Stack:    []int{},
		MinStack: []int{},
	}
}

func (m *MinStack) Push(val int) {
	m.Stack = append(m.Stack, val)
	if len(m.MinStack) == 0 || val <= m.GetMin() {
		m.MinStack = append(m.MinStack, val)
	}
}

func (m *MinStack) Pop() {
	val := m.Stack[len(m.Stack)-1]
	m.Stack = m.Stack[:len(m.Stack)-1]

	if val == m.GetMin() {
		m.MinStack = m.MinStack[:len(m.MinStack)-1]
	}
}

func (m *MinStack) Top() int {
	return m.Stack[len(m.Stack)-1]
}

func (m *MinStack) GetMin() int {
	return m.MinStack[len(m.MinStack)-1]
}

// Day 3: Daily Temperatures
func DailyTemperatures(tem []int) []int {
	var stack []int
	answer := make([]int, len(tem))

	for i := 0; i < len(tem); i++ {
		for len(stack) > 0 && tem[i] > tem[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[prev] = i - prev
		}

		stack = append(stack, i)
	}
	return answer
}

// Day 4: Largest Rectangle in Histogram
func LargestRectangleInHistogram(heights []int) int {
	maxArea := 0
	stack := []int{}

	for i := 0; i <= len(heights); i++ {
		var current int
		if i == len(heights) {
			current = 0
		} else {
			current = heights[i]
		}

		for len(stack) > 0 && current < heights[stack[len(stack)-1]] {
			heightIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			height := heights[heightIndex]

			width := i

			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			area := height * width

			if area > maxArea {
				maxArea = area
			}
		}
		if i < len(heights) {
			stack = append(stack, i)
		}
	}
	return maxArea
}
