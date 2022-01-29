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
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
func miniMaxSum(arr []int32) {
	min := arr[0]
	max := arr[0]
	Max := 0
	Min := 0
	for _, val := range arr {
		switch {
		case min > val:
			min = val
		case max < val:
			max = val
		}
	}
	for _, val := range arr {
		switch {
		case (min == val) && (min != max):
			Min += int(val)
		case (max == val) && (min != max):
			Max += int(val)
		default:
			Max += int(val)
			Min += int(val)
		}
	}
	if max == min {
		Max -= (int)(max)
		Min = Max
	}
	fmt.Println(Min, Max)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
