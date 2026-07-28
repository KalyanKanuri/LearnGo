package twopointers

func MoveZeroes(nums []int) []int {
	writer := 0

	for reader := range nums {
		if nums[reader] == 0 {
			continue
		}
		nums[writer], nums[reader] = nums[reader], nums[writer]
		writer++
	}

	return nums
}