package task3

func flippingBits(n int64) int64 {
	index := 0
	for index < 32 {
		n = n ^ (1 << index)
		index++
	}
	return n

}
