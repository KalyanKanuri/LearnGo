package twopointers

func IsPalindrome(nums []int) bool {
	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left] != nums[right] {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}
