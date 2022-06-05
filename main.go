package main

import (
	"errors"
	"fmt"
	"sort"
)

//Merges overlapping intervals in to one interval
//An interval is a slice of length 2, where the first element is the start value and the second element is the end value
//Returns an error if the input is not valid
func MergeIntervals(intervals [][]int) ([][]int, error) {

	//return empty array if intervals is empty
	if len(intervals) == 0 {
		return [][]int{}, nil
	}
	//check if the provided intervals array is valid
	if !validateIntervals(intervals) {
		return nil, errors.New("invalid intervals")
	}
	// sort intervals by first value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// start with first interval in result, because we are sorting by first value
	// the first interval in the result will contain the smallest first value
	results := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		//check if the current interval overlaps with the last interval in the result
		if intervals[i][0] <= results[len(results)-1][1] {
			//if so, update the last interval in the result to include the end of the current interval
			results[len(results)-1][1] = maxOf(results[len(results)-1][1], intervals[i][1])
		} else {
			//if not, add the current interval to the result
			results = append(results, intervals[i])
		}
	}

	return results, nil
}

//validates that the end value of the interval is greater or equal than the start value
func validateIntervals(intervals [][]int) bool {
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] > intervals[i][1] {
			return false
		}
	}
	return true
}

func maxOf(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := [][]int{{1, 4}, {2, 3}, {4, 5}}
	merged, err := MergeIntervals(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(merged)
}
