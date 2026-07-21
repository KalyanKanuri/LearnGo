package twopointers

func MoveZeroes[T comparable](arr []T) []T {
	writer := 0
	var zero T

	for reader := range arr {
		if arr[reader] == zero || reader == writer {
			continue
		}

		arr[writer], arr[reader] = arr[reader], arr[writer]
		writer++
	}
	return arr
}
