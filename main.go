package main

import "fmt"

func MergeIntervals(intervals [][]int) ([][]int, error) {
	if len(intervals) == 0 {
		return intervals, nil
	}
	return intervals, nil
}

func main() {
	input := [][]int{{1, 4}, {2, 3}, {4, 5}}
	merged, err := MergeIntervals(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(merged)
}
