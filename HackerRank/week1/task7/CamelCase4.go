package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func S(str string) string {
	res := ""
	for i, r := range str {
		if unicode.IsUpper(r) {
			if i == 0 {
				res += strings.ToLower(string(r))
			} else {
				res += " " + strings.ToLower(string(r))
			}
		} else {
			res += string(r)
		}
	}
	if strings.Contains(res, "()") {
		res = strings.Replace(res, "()", "", 1)
	}
	return res
}
func C(str string, num int) string {
	res := ""
	flag := 0
	if num == 2 {
		flag = 1
	}
	for _, val := range str {
		if val == ' ' {
			flag = 1
		}
		if flag == 1 && val != ' ' {
			flag = 0
			res += strings.ToUpper(string(val))
		} else if flag == 0 && val != ' ' {
			res += string(val)
		}
	}
	if num == 3 {
		res += "()"
	}

	return res
}

func get(str string) (string, int) {
	num := 0
	switch {
	case str[0] == 'S':
		num = 0
	case str[0] == 'C':
		switch {
		case str[2] == 'V':
			num = 1
		case str[2] == 'C':
			num = 2
		case str[2] == 'M':
			num = 3
		}

	}
	return str[4:], num
}
func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		str := sc.Text()
		num := 0
		fmt.Scanf("%s\n", &str)
		str, num = get(str)
		if num == 0 {
			fmt.Println(S(str))
		} else {
			fmt.Println(C(str, num))
		}
	}
}


