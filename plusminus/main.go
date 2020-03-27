package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	total := len(arr)
	positive := 0
	negative := 0
	zeros := 0

	for _, value := range arr {
		if value == 0 {
			zeros += 1
		}
		if value > 0 {
			positive += 1
		}
		if value < 0 {
			negative += 1
		}
	}
	fmt.Printf("%.6f\n", float64(positive)/float64(total))
	fmt.Printf("%.6f\n", float64(negative)/float64(total))
	fmt.Printf("%.6f\n", float64(zeros)/float64(total))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
