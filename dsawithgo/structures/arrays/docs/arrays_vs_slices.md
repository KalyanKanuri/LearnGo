# Arrays Vs Slices

Arrays are actual data structures that holds the data in the memory, whereas slices are actually a data structure which has an underlying array, a pointer to the underlying array, len(), cap() it actually contains 3 values in the memory, len() of the array, cap() of the array, pointer to the array

In arrays we cannot append a new value or extend the array, whereas in slices we can append new values internally when the array is full go decides to create a new array, copy all the existing values to the new array and append the value and update the array pointer in the slice

```go
arr := [5]int{1, 2, 3, 4, 5}
slc := []int{1, 2, 3, 4, 5}
// this is a simple difference in code but underneath there's much more work being done
```

## make()

make() function in go is actually used to create slice, it takes 3 arguments []T - type of the slice, len() length of slice, cap() capacity of the slice

```go
slc := make([]int, 5, 5)

// here []int is the type of the slice, 5 is the length of the slice, 5 is the capacity of the slice
```

## copy()

usually we use below syntax to copy arrays

```go
slc1 := make([]int, 0, 5)
slc1 = append(slc1, 10)
slc1 = append(slc1, 20)
slc2 := slc1
```

what does this actually do it will just copy the slice components of the slice1 to slice2 it will not create the copy of the underlying array, hence both slc2 and slc1 points to the same underlying array with the pointer when the slc1 is updated it will be reflected in slc2 as well

hence if we want to create a new copy of a slice we can use copy() function provided by go

```go
src := []int{1, 2, 3, 4, 5}
dst := make([]int, len(src))
n := copy(dst, src)

// n is the number of elements copied into dst from src
```
