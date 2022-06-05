# Coding Task 2 ![](https://github.com/julianfbeck/DCI/actions/workflows/test.yml/badge.svg)

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