# Time Complexity Analysis

Time complexity Analysis means analyzing how the runtime increases as the input(n) increases, there are 7 order of complexities

## o(n) - constant time

Even if the input increases the runtime doesn't increase.
e.g., accessing an element with an index in array

```go
value := arr1[0]
```

## o(log n) - logarithmic time

The main concept of o(log n) is halving the search space with each iteration or each visit of the operation
e.g., suppose we are trying to access a letter on page 900 in a book of 1000 pages, with linear search we will visit each page but with logarithmic approach (binary search) we half the search pages into halves that is 500th page then we look if the required search term is before 500th page set or after 500th page set then again go halving the data set like 250, 125, 62.5 etc.,

```go
func doBinarySearch(arr []int, target int) int {
    low, high := 0, len(arr-1)

    for low <= high {
        mid = low + (high - low) / 2
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            low = mid+1
        } else if arr[mid] > target {
            high = mid-1
        }
    }
    return -1
}
```

## o(n) - linear time

The runtime increases as the data increases linearly
e.g., as we have taken an example of finding a text in a book, if the book contains 100 pages we have to go through 100 pages for finding the text

```go
func findText(arr []string, searchString string) int {
    for _, val := range arr {
        if val == searchString {
            return arr[val]
        }
    }
    return -1
}
```

## o(n log n) - log linear time

This is the combination of logarithmic as well as linear time pattern, that is repeatedly perform linear operations on repeatedly halved sets[logarithmic]
e.g., lets say we are organizing a music festival we have to distribute wristbands to each attendee, with linear approach we have to distribute wristband to each person in a single line if the crowd increases the wait time also increases, now with log linear approach we divide the crowd in half and assign a manager to each half and these managers divides the crowd into half again and assign supervisors in this way if we have 100 attendees they are divided into 10 persons each batch, now for each batch we will distribute wristbands linearly

```go
func distributeWristBands(batch []int) {
    for _, person := range batch {
        fmt.Println("Assigning wristband to person", person)
    }
}

func divideAttendeesToBatches(attendees []int) {

    distributeWristBands(attendees) // this is just a raw demo to show that we process o(n)
    // in real world this is a bug as each function call distributes the band to same person recursively

    half := len(attendees/2)
    leftHalf := attendees[:half]
    rightHalf := attendees[half:]

    divideAttendeesToBatch(leftHalf)
    divideAttendeesToBatch(rightHalf)
}

```

## o(n^2) - quadratic time

execution time of the operation grows proportionally to the input size
e.g., we can easily consider loop over loop for this example when the input increases the loop doubles

```go
for i := 0; i <= len(arr); i++ {
    for j := 1; j <= i; j++ {
        fmt.Println(i, j)
    }
}
```

## o(2^n) - exponential time complexity

the execution time doubles with every single addition to the input data

```go
// this is simply a function creates two more function calls each time this function is called
func febanoci_series(n int) int {
    if n < 1 {
        return n
    }
    febanoci_series(n-1) + febanoci_series(n-2)
}
```

## o(n!) - factorial time

the execution time increases n factorial times with with every single addition to the input data

```go
func factorial(n int) int {
    var result int = 1
    for i=1; i<n; i++ {
        result = result * i
    }
    return result
}
```
