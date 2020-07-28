package main

import (
	"fmt"
	"sort"
)

func mergeWrong(intervals [][]int) [][]int {
	res := make([][]int, 0)
	if len(intervals) == 0 {
		return res
	}
OUTER:
	for _, s1 := range intervals {
		if len(res) == 0 {
			res = append(res, s1)
			continue
		}
		for idx, s2 := range res {
			/*
				s1:   (-----)
				s2: (------------)
				s1 is already contain in s2, skip
			*/
			if s1[0] >= s2[0] && s1[1] <= s2[1] {
				continue OUTER
			}

			/*
				s1: (------------)
				s2:   (-----)
				s2 is contains by s1, replace s2 with s1
			*/
			if s1[0] <= s2[0] && s1[1] >= s2[1] {
				res[idx] = s1
				continue OUTER
			}

			/*
					s1: (------------)
					s2:   (-------------)
				we should expand s2[0] with s1[0]
			*/
			if s1[0] <= s2[0] && s1[1] <= s2[1] && s1[1] >= s2[0] {
				res[idx][0] = s1[0]
				continue OUTER
			}

			/*
					s1:        (------------)
					s2:   (-------------)
				we should expand s2[1] with s1[1]
			*/
			if s1[0] >= s2[0] && s1[1] >= s2[1] && s1[0] <= s2[1] {
				res[idx][1] = s1[1]
				continue OUTER
			}
		}
		/*
			s1:                     (------------)
			s2:   (-------------)
			append s1 to res directly
		*/
		res = append(res, s1)
	}
	return res
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var result [][]int
	for i := 0; i < len(intervals); i++ {
		if len(result) == 0 || intervals[i][0] > result[len(result)-1][1] {
			result = append(result, intervals[i])
		} else if intervals[i][1] > result[len(result)-1][1] {
			result[len(result)-1][1] = intervals[i][1]
		}
	}
	return result
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{1, 4}, {1, 5}}))
}
