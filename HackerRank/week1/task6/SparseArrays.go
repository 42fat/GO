package main

import (
	"fmt"
)

/*
 * Complete the 'matchingStrings' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY strings
 *  2. STRING_ARRAY queries
 */

func matchingStrings(str []string, queries []string) []int32 {
	res := make([]int32, len(queries), len(queries))
	for i, val := range queries {
		for _, value := range str {
			if value == val {
				res[i]++
			}
		}
	}
	return res
}

func main() {
	str := []string{"aba", "baba", "aba", "xzxb"}
	str2 := []string{"aba", "xzxb", "ab"}
	fmt.Println(matchingStrings(str, str2))
}
