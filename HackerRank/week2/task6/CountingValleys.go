package task6

func countingValleys(steps int32, path string) int32 {
	h := 0
	d := 0
	flag1 := 0
	for _, val := range path {
		switch {
		case val == 'U':
			h++
		case val == 'D':
			h--
		}
		if h < 0 && flag1 == 0 {
			flag1 = 1
		}
		if h == 0 && flag1 == 1 {
			d++
			flag1 = 0
		}
	}
	return int32(d)
}
