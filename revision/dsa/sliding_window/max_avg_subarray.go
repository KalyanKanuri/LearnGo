package slidingwindow

func MaxAvgSubArray(nums []int, m int) float64 {
	left := 0
	right := m
	maxSum := 0
	currentSum := 0

	for i := range m {
		currentSum += nums[i]
	}
	maxSum = currentSum

	for left < right {
		currentSum = currentSum - nums[left] + nums[right]

		if currentSum > maxSum {
			maxSum = currentSum
		}
		left++
		right--
	}
	maxAvg := float64(maxSum) / float64(m)
	return maxAvg
}
