# Case Scenarios in DSA

## Best Case Scenario

Best Case Scenario is nothing but minimum time taken for an algorithm

```go
func searchArray(arr []int, n int) int, int {
    for pos, elem := range arr {
        if elem == n {
            return elem, pos
        }
    }
    return -1, n
}

// Here search element is 1, when the search starts it immediately returns as the search element is first element in the array
func main() {
    arr := []int{1, 2, 3, 4, 5}
    searchArray(arr, 1)
}
// Time Complexity -- o(1)
// Space Complexity -- o(1)
```

## average Case Scenario

Average Case Scenario is expected time taken for an algorithm

```go
func searchArray(arr []int, n int) int, int {
    for pos, elem := range arr {
        if elem == n {
            return elem, pos
        }
    }
    return -1, n
}

// Here search element is 3, we'll go through half of the array size which makes the time to n/2
func main() {
    arr := []int{1, 2, 3, 4, 5}
    searchArray(arr, 3)
}
// Time Complexity -- o(n) -- actually we'll get n/2 but for Big O Notation we ignore constants which is /2 in this case hence the time complexity becomes o(n)
// Space Complexity -- o(1)
```

## Worst Case Scenario

Worst Case Scenario is the maximum time taken for an algorithm

```go
func searchArray(arr []int, n int) int, int {
    for pos, elem := range arr {
        if elem == n {
            return elem, pos
        }
    }
    return -1, n
}

// Here search element is 5, we are looping through every element in the array which takes maximum time for the loop to end
func main() {
    arr := []int{1, 2, 3, 4, 5}
    searchArray(arr, 5)
}
// Time Complexity -- o(n)
// Space Complexity -- o(1)
```
