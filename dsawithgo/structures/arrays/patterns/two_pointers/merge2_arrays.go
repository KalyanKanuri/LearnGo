package twopointers

func MergeSortedArrays(nums1, nums2 []int) []int {
	ptr1 := 0
	ptr2 := 0
	merged := make([]int, 0, len(nums1)+len(nums2))

	for ptr1 < len(nums1) && ptr2 < len(nums2) {
		if nums1[ptr1] <= nums2[ptr2] {
			merged = append(merged, nums1[ptr1])
			ptr1++
		} else {
			merged = append(merged, nums2[ptr2])
			ptr2++
		}
	}
	merged = append(merged, nums1[ptr1:]...)
	merged = append(merged, nums2[ptr2:]...)
	return merged
}
