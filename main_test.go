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

func Test_mergeIntervalForEmptyInterval(t *testing.T) {
	intervals := [][]int{}

	expected := [][]int{}

	got, err := MergeIntervals(intervals)
	assert.Nil(t, err, "should not return error")
	assert.Equal(t, expected, got, "should return empty array")
}

func Test_mergeIntervalForInvalidInterval(t *testing.T) {
	intervals := [][]int{
		{1, 4},
		{2, 3},
		{4, 5},
		{5, 4},
	}

	_, err := MergeIntervals(intervals)
	assert.NotNil(t, err, "should return error")
}

func Test_max(t *testing.T) {
	assert.Equal(t, 5, maxOf(5, 3))
	assert.Equal(t, 5, maxOf(5, 5))
	assert.Equal(t, 3, maxOf(3, 3))
}

func Test_validateIntervals(t *testing.T) {
	for _, testCase := range validateIntervalsTestCases {
		got := validateIntervals(testCase.intervals)
		assert.Equal(t, testCase.expected, got, "should validate intervals")
	}
}

func Benchmark_mergeIntervals(b *testing.B) {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	for i := 0; i < b.N; i++ {
		MergeIntervals(intervals)
	}
}
