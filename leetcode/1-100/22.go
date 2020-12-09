package main

import "fmt"

// https://leetcode-cn.com/problems/generate-parentheses/solution/zui-jian-dan-yi-dong-de-dong-tai-gui-hua-bu-lun-da/
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	res := make([][]string, 0)
	res = append(res, []string{""})
	res = append(res, []string{"()"})
	for i := 2; i <= n; i++ {
		l := make([]string, 0)
		for j := 0; j < i; j++ {
			nowList1 := res[j]
			nowList2 := res[i-1-j]
			for _, k1 := range nowList1 {
				for _, k2 := range nowList2 {
					l = append(l, "("+k1+")"+k2)
				}
			}
		}
		res = append(res, l)
	}
	return res[n]
}

func main() {
	fmt.Println(generateParenthesis(3))
}
