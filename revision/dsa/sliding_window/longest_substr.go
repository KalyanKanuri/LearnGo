package slidingwindow

func LongestSubString(str string) int {
	left := 0
	lastSeen := make(map[rune]int)
	maxLen := 0
	windowLen := 0
	r := []rune(str)

	for idx, ch := range r {
		if prevIdx, ok := lastSeen[ch]; ok {
			left = max(left, prevIdx+1)
		}
		lastSeen[ch] = idx
		windowLen = idx - left + 1
		if windowLen > maxLen {
			maxLen = windowLen
		}
	}
	return maxLen
}
