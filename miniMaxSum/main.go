package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
	min := int64(math.MaxInt64)
	max := int64(0)
	sum := int64(0)

	for i := 0; i < 5; i++ {
		temp := int64(arr[i])
		sum += temp
		if temp < min {
			min = temp
		}
		if temp > max {
			max = temp
		}

	}
	fmt.Printf("%v %v", sum-max, sum-min)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

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
