package task7

import (
	"strings"
)

func pangrams(s string) string {
	i := 97
	for {
		if strings.Contains(s, string(i)) || strings.Contains(s, strings.ToUpper(string(i))) {

		} else {
			return "not pangram\n"
		}
		if i == 122 {
			break
		}
		i++
	}
	return "pangram\n"

}
