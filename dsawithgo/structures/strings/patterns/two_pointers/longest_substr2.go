package twopointersstr

func LongestSubStr2(s string) int {
	left := 0
	maxLen := 0
	windowLen := 0
	lastSeen := make(map[rune]int)

	r := []rune(s)

	for idx, char := range r {
		if prevIdx, ok := lastSeen[char]; ok {
			left = max(left, prevIdx+1)
		}
		lastSeen[char] = idx
		windowLen = idx - left + 1
		if windowLen > maxLen {
			maxLen = windowLen
		}
	}
	return maxLen
}
