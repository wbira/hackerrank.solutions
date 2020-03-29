package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Complete the camelcase function below.
func camelcase(s string) int32 {
	numberOfWords := int32(1)
	for _, letter := range s {
		if letter >= 60 && letter <= 90 {
			numberOfWords += 1
		}
	}
	return numberOfWords
}

func main() {
	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 1024*1024)

	// s := readLine(reader)
	s := "saveChangesInTheEditor"
	result := camelcase(s)
	fmt.Println(result)
	// fmt.Fprintf(writer, "%d\n", result)

	// writer.Flush()
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
