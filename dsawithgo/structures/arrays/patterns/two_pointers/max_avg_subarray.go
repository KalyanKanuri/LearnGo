package twopointers

func MaxAvgSubArray(nums []int, size int) int {
	left := 0
	currentSum := 0
	maxSum := 0

	for i := range size {
		currentSum += nums[i]
	}

	for right := range nums {
		currentSum = currentSum - nums[left] + nums[right]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	maxAvg := maxSum / size
	return maxAvg
}
