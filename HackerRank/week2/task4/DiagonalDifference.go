package task4

import "math"

func diagonalDifference(arr [][]int32) int32 {
	suml := int32(0)
	sumr := int32(0)
	n := len(arr)
	for i := 0; i < n; i++ {
		suml += arr[i][i]
		sumr += arr[i][n-i-1]
	}
	return int32(math.Abs(float64(suml - sumr)))
}
