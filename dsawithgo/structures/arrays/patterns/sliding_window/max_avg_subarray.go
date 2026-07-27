package sdwInt

func MaxAvgSubArray(nums []int, size int) float64 {
	left := 0
	right := size
	currentSum := 0
	maxSum := 0

	for i := range size {
		currentSum += nums[i]
	}
	maxSum = currentSum

	for right < len(nums) {
		currentSum = currentSum - nums[left] + nums[right]
		if currentSum > maxSum {
			maxSum = currentSum
		}
		left++
		right++
	}

	maxAvg := float64(maxSum) / float64(size)
	return maxAvg
}
