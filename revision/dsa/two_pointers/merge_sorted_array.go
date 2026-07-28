package twopointers

func MergeSortedArrays(nums1 []int, nums2 []int) []int {
	ptr1 := 0
	ptr2 := 0
	mergedArr := make([]int, 0, len(nums1)+len(nums2))

	for ptr1 < len(nums1) && ptr2 < len(nums2) {
		if nums1[ptr1] < nums2[ptr2] {
			mergedArr = append(mergedArr, nums1[ptr1])
			ptr1++
		} else {
			mergedArr = append(mergedArr, nums2[ptr2])
			ptr2++
		}
	}
	mergedArr = append(mergedArr, nums1[ptr1:]...)
	mergedArr = append(mergedArr, nums2[ptr2:]...)
	return mergedArr
}
