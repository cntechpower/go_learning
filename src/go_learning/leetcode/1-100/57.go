package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	res := make([][]int, 0)
	newStart, newEnd := newInterval[0], newInterval[1]
	idx, n := 0, len(intervals)
	for idx < n && newStart > intervals[idx][0] {
		res = append(res, intervals[idx])
		idx++
	}
	if len(res) == 0 || res[len(res)-1][1] < newStart {
		res = append(res, newInterval)
	} else {
		res[len(res)-1][1] = max(res[len(res)-1][1], newEnd)
	}

	for idx < n {
		interval := intervals[idx]
		start, end := interval[0], interval[1]
		idx++
		if res[len(res)-1][1] < start {
			res = append(res, interval)
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], end)
		}
	}
	return res
}

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
}
