package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mergeIntervalsTestCases = []struct {
	intervals [][]int
	expected  [][]int
}{
	{
		intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
	},
	{
		intervals: [][]int{{-100, 4}, {4, 5}},
		expected:  [][]int{{-100, 5}},
	},
	{
		intervals: [][]int{{25, 30}, {2, 19}, {14, 23}, {4, 8}},
		expected:  [][]int{{2, 23}, {25, 30}},
	},
	{
		intervals: [][]int{{1, 4}, {0, 1}},
		expected:  [][]int{{0, 4}},
	},
	{
		intervals: [][]int{{1, 4}, {2, 3}},
		expected:  [][]int{{1, 4}},
	},
}

var validateIntervalsTestCases = []struct {
	intervals [][]int
	expected  bool
}{
	{
		intervals: [][]int{{1, 4}, {2, 3}, {4, 5}},
		expected:  true,
	},
	{
		intervals: [][]int{{5, 4}},
		expected:  false,
	},
	{
		intervals: [][]int{{4, 4}},
		expected:  true,
	},
	{
		intervals: [][]int{{1, -1}},
		expected:  false,
	},
}

func Test_mergeInterval(t *testing.T) {
	for _, testCase := range mergeIntervalsTestCases {
		got, err := MergeIntervals(testCase.intervals)
		assert.Nil(t, err, "should not return error")
		assert.Equal(t, testCase.expected, got, "should merge intervals")
	}
}
