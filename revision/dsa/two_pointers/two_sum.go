package twopointers

func TwoSum(nums []int, target int) []int {
	left := 0
	right := 1

	for right < len(nums) {
		if nums[left]+nums[right] == target {
			return []int{nums[left], nums[right]}
		} else {
			left++
		}
		right++
	}
	return []int{}
}
