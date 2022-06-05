package mergeconcurrently

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



func Test_mergeIntervalConcurrently(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}, {0, 1}, {2, 3}, {4, 5}, {5, 6}, {7, 8}, {9, 10}, {11, 12}, {13, 14}, {15, 16}, {17, 18}, {20, 22}, {10, 14}}
	expected := [][]int{{0, 6}, {7, 14}, {15, 18}, {20, 22}}
	
	got := MergeIntervalsConcurrently(intervals)
	assert.Equal(t, expected, got, "should merge intervals")
}
