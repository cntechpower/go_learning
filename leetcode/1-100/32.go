package main

import "fmt"

func longestValidParentheses1(s string) int {
	dp := make([]int, len(s))
	max := 0
	maxInt := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				left := 0
				if i >= 2 {
					left = dp[i-2]
				}
				dp[i] = left + 2
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				left := 0
				if i-dp[i-1] >= 2 {
					left = dp[i-dp[i-1]-2]
				}
				dp[i] = dp[i-1] + left + 2
			}
			max = maxInt(max, dp[i])
		}
	}
	return max
}

func longestValidParentheses2(s string) int {
	stack := make([]int, 0)
	max := 0
	maxInt := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				max = maxInt(max, i-stack[len(stack)-1])
			}
		}
	}
	return max
}

func longestValidParentheses(s string) int {
	left, right, max := 0, 0, 0
	maxInt := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			max = maxInt(max, 2*right)
		} else if right >= left {
			left = 0
			right = 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			max = maxInt(max, 2*left)
		} else if left >= right {
			left = 0
			right = 0
		}
	}
	return max
}

func main() {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
}
