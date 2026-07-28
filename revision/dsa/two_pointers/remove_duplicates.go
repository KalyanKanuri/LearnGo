package twopointers

func RemoveDuplicates(nums []int) []int {
	writer := 1

	for reader := range nums {
		if reader == 0 {
			continue
		}

		if nums[reader] != nums[reader-1] {
			nums[writer] = nums[reader]
			writer++
		}
	}
	return nums[:writer]
}