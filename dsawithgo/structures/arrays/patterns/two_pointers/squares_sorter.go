package twopointers

func SortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	writer := len(nums) - 1
	sqrs := make([]int, len(nums))

	for left <= right {
		leftSqr := nums[left] * nums[left]
		rightSqr := nums[right] * nums[right]

		if leftSqr > rightSqr {
			sqrs[writer] = leftSqr
			left++
		} else {
			sqrs[writer] = rightSqr
			right--
		}
		writer--
	}
	return sqrs
}
