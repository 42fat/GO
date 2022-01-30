package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */
func timeConversion(s string) string {
	result := " "
	if strings.Contains(s, "PM") {
		result = strings.TrimRight(s, "PM")
		tmp, _ := strconv.Atoi(result[:2])
		if tmp != 12 {
			tmp += 12
		}
		result = strconv.Itoa(tmp) + result[2:]
	}
	if strings.Contains(s, "AM") {
		result = strings.TrimRight(s, "AM")
		tmp, _ := strconv.Atoi(result[:2])
		var res string
		if tmp == 12 {
			res = "00"
		} else {
			if tmp < 10 {
				res = "0" + strconv.Itoa(tmp)
			} else {
				res = strconv.Itoa(tmp)
			}
		}
		result = res + result[2:]
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
