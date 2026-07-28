package threepointers

import "fmt"

func Merge2SortedArrs(nums1, nums2 []int, m, n int) {
	ptr1 := m - 1
	ptr2 := n - 1
	writer := m + n - 1

	for ptr1 >= 0 && ptr2 >= 0 {
		if nums1[ptr1] < nums2[ptr2] {
			nums1[writer] = nums2[ptr2]
			ptr2--
		} else {
			nums1[writer] = nums1[ptr1]
			ptr1--
		}
		writer--
	}
	fmt.Println(nums1)
}
