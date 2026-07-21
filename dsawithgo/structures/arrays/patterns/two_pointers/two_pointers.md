# Two Pointers

Two Pointers is a pattern where we initialize two variables to traverse through indices in an array either in the same direction or opposite direction as per the problem

```txt
--------------------
|1 | 2 | 3 | 4 | 5 |
--------------------
```

-- pseudocode

left = 0,
right = len(arr) - 1

left, right = right, left
                4,    1
left++
right--

```txt
--------------------
|5 | 2 | 3 | 4 | 1 |
--------------------
```
