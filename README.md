# Coding Task 2 ![https://github.com/julianfbeck/DCI/actions/workflows/test.yml](https://github.com/julianfbeck/DCI/actions/workflows/test.yml/badge.svg)

## Question:
Given an array of intervals, merge all overlapping intervals into one interval.

### Assumptions:
- The intervals are considered overlapping if the end of one interval is **equal or greater** than the start of another interval.
- The intervals are not sorted.
- The intervals only contain integers. (i.e. no nil values, however this is currently result in an compile error due to the type system of golang)
- Intervals can also be negative.

### Characteristics:
#### Time Complexity:
- The time complexity of this algorithm is O(nlog(n)) as we need to sort the array. All other operations are less.
#### Space Complexity:
- Space complexity is O(n), because we are storing the merged intervals in a new array and in the worst case scenario there may not be any intervals and we would need to store all elements.
#### Robustness and Scalability:
The Robustness can be improve by: 
- Validating the input and throw an error if the input is invalid.
- Using Custom types to represent the intervals.(i.e. struct Interval { int start; int end; })
- Using a package to represent the intervals.
- Add logging to the algorithm.

The Scalability can be improve by:
- Splitting the array into smaller chunks and merging them (using goroutines).
  - This will reduce the time used to sort the complete array.
- For large memory usage a worker pool can be used.

### Implementation Details:
We sort the array of intervals first. This makes all subsequent operations a lot easier.
We start with the first interval in the result interval. Next we iterate over the rest of the intervals and merge them into the result.
Herby we check the max value of the result or current interval to determine the ending position of the result interval.
If the current and next interval overlap, we merge them into the result. If not we just add the current interval to the result.

#### Example Walkthrough:
- Input: [[1, 3], [8, 10], [2, 6]]
- Step 1: Sort the intervals by start value.
  - Input: [[1, 3], [2, 6], [8, 10]]
  - Results: []
- Step 2: Populate the result array with the first interval.
  - Input: [[1, 3], [2, 6], [8, 10]]
  - Results: [[1, 3]] 
- Step 3: loop through the Intervals index = 1. 
  - Input: [[1, 3], [2, 6], [8, 10]]	
  - Results: [[1, 3]]
  - Index Interval: [2, 6]
  - Check if the start of the current interval [2,6] is smaller or equal than the end in the last element of the result [1,3] array: 2 <= 3: True
    - True: Update the end of the last result element with the max value of the last result value or the current interval end value: max(3, 6) = 6
- Step 3: (continued) index = 2.
  - Input: [[1, 3], [2, 6], [8, 10]]
  - Results: [[1, 6]]
  - Index Interval: [8, 10]
  - Check if the start of the current interval [8,10] is smaller or equal than the end in the last element of the result [1,6] array: 8 <= 6: False
    - False: Append the current interval to the result array: [[1, 6], [8, 10]]


## Execution:

To execute the code run the following command in the terminal:
```
make run
```

Alternatively, docker can be used to execute the code locally directly in a container
```
docker run -it --rm -v $(pwd):/app -w /app golang:1.18-alpine go run ./...
```


## Testing
To test the code run the following command in the terminal:
```
make install
make test
```
This will install the testify package and also generate a coverage report.

Alternatively, docker can be used to execute the tests locally directly in a container
```
docker run -it --rm -v $(pwd):/app -w /app --env CGO_ENABLED=0 --env GOOS=linux golang:1.18-alpine go test -v ./...
```