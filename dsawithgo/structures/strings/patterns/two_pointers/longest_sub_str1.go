package twopointersstr

func LongestSubString1(str string) int {
	left := 0
	windowLength := 0
	maxLength := 0
	arrState := make(map[rune]int)
	runeSlc := []rune(str)

	for right, char := range runeSlc {
		arrState[char]++
		for arrState[char] > 1 {
			arrState[runeSlc[left]]--
			left++
		}
		windowLength = right - left + 1
		if windowLength > maxLength {
			maxLength = windowLength
		}
	}
	return maxLength
}
