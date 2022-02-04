package task5

func countingSort(arr []int32) []int32 {
	tmp := make([]int32, 100, 100)
	for i := 0; i < len(arr); i++ {
		tmp[arr[i]]++
	}
	return tmp
}
