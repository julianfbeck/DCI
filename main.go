package main

import (
	"fmt"
	"sort"
)

func MergeIntervals(intervals [][]int) ([][]int, error) {

	//return empty array if input is empty
	if len(intervals) == 0 {
		return intervals, nil
	}

	//sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	results := [][]int{intervals[0]}

	for i := 1; i < len(intervals)-1; i++ {
		if intervals[i][0] <= results[i][1] {

			results[i][1] = intervals[i][1]
		} else {
			results = append(results, intervals[i])
		}
	}

	return results, nil
}

func main() {
	input := [][]int{{1, 4}, {2, 3}, {4, 5}}
	merged, err := MergeIntervals(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(merged)
}
