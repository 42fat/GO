package task8

func marsExploration(s string) int32 {
	var a int32 = 0
	n := len(s) / 3
	i := 0
	for i < n {
		if s[i*3+0] != 'S' {
			a++
		}
		if s[i*3+1] != 'O' {
			a++
		}
		if s[i*3+2] != 'S' {
			a++
		}

		i++
	}
	return a
}
