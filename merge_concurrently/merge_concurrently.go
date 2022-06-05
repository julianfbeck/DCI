package mergeconcurrently

import (
	"log"
	"runtime"
	"sort"
	"sync"
)

// Chunks an array of intervals into n = runtime.NumCPU() chunks.
func chuckSlice(intervals [][]int) [][][]int {
	var dividedIntervals [][][]int

	//available CPUs
	cpus := runtime.NumCPU()
	chunkSize := (len(intervals) + cpus - 1) / cpus

	for i := 0; i < len(intervals); i += chunkSize {
		end := i + chunkSize

		if end > len(intervals) {
			end = len(intervals)
		}
		dividedIntervals = append(dividedIntervals, intervals[i:end])
	}
	log.Printf("chunked %d intervals into %d chunks", len(intervals), len(dividedIntervals))
	return dividedIntervals
}

func MergeIntervalsConcurrently(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	chunkedIntervals := chuckSlice(intervals)
	resultsChan := make(chan [][]int, len(chunkedIntervals))

	var wg sync.WaitGroup
	// start mergeIntervals jobs
	for _, chunk := range chunkedIntervals {
		wg.Add(1)
		go func(chunk [][]int) {
			defer wg.Done()
			resultsChan <- mergeIntervals(chunk)
		}(chunk)
	}
	// wait for all jobs to finish
	wg.Wait()
	close(resultsChan)

	//collect results
	jobResults := [][]int{}
	for result := range resultsChan {
		jobResults = append(jobResults, result...)
	}
	log.Printf("merged %d intervals into %d intervals using %d workers", len(intervals), len(jobResults), len(chunkedIntervals))

	//merge job results
	return mergeIntervals(jobResults)

}

// Merges arrays of intervals.
func mergeIntervals(localInterval [][]int) [][]int {
	//return empty array if intervals is empty
	if len(localInterval) == 0 {
		return [][]int{}
	}

	// sort intervals by first value
	sort.Slice(localInterval, func(i, j int) bool {
		return localInterval[i][0] < localInterval[j][0]
	})
	// start with first interval in result, because we are sorting by first value
	// the first interval in the result will contain the smallest first value
	localResults := [][]int{localInterval[0]}

	for i := 1; i < len(localInterval); i++ {
		//check if the current interval overlaps with the last interval in the result
		if localInterval[i][0] <= localResults[len(localResults)-1][1] {
			//if so, update the last interval in the result to include the end of the current interval
			//Mark: Data Race Read
			localResults[len(localResults)-1][1] = maxOf(localResults[len(localResults)-1][1], localInterval[i][1])
		} else {
			//Mark Data Race previous write
			//if not, add the current interval to the result
			localResults = append(localResults, localInterval[i])
		}
	}

	return localResults
}

func maxOf(a, b int) int {
	if a > b {
		return a
	}
	return b
}
