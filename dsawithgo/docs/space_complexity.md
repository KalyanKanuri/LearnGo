# Space Complexity Analysis

Space Complexity means analyzing how much extra space an algorithm consumes as the input size increases

## o(1) - constant space

the algorithm consumes constant space whatever the input is

```go
func sum(arr []int) int {
    var total int = 0
    for _, i := range arr {
        total += i
    }
    return total
}
```

## o(n) - linear space

the space increases as the input increases linearly

```go
func square(arr []int) []int {
    result := make([]int, 0)
    for i := range arr {
        result[i] = arr[i] * arr[i]
    }
    return result
}
```

## o(n^2) - quadratic space

the space squares as the input increases

```go
func matrix(n int) [][]int {
    grid := make([][]int, n)
    for i := range grid {
        grid[i] = make([]int, n)
    }
    return grid
}

```

here let's say input is n=3, then the array becomes
0 0 0
0 0 0
0 0 0
the memory occupied is for 9 cells i.e. 3^2 which is square of the input
