package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumNumber function below.
func minimumNumber(n int32, password string) int32 {
	// Return the minimum number of characters to make the password strong
	numbers := "0123456789"
	lower_case := "abcdefghijklmnopqrstuvwxyz"
	upper_case := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special_characters := "!@#$%^&*()-+"

	hasDigit := false
	hasLower := false
	hasUpper := false
	hasSpecial := false
	for _, r := range password {
		char := string(r)
		if !hasDigit {
			hasDigit = strings.Contains(numbers, char)
		}
		if !hasLower {
			hasLower = strings.Contains(lower_case, char)
		}
		if !hasUpper {
			hasUpper = strings.Contains(upper_case, char)
		}
		if !hasSpecial {
			hasSpecial = strings.Contains(special_characters, char)
		}
	}

	numberOfNeededChars := int32(0)
	if !hasDigit {
		numberOfNeededChars += 1
	}
	if !hasLower {
		numberOfNeededChars += 1
	}
	if !hasUpper {
		numberOfNeededChars += 1
	}
	if !hasSpecial {
		numberOfNeededChars += 1
	}

	passwordLength := int32(len(password))
	if passwordLength+numberOfNeededChars < 6 {
		return 6 - passwordLength
	}

	return numberOfNeededChars
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	password := readLine(reader)

	answer := minimumNumber(n, password)
	fmt.Println(answer)
	fmt.Fprintf(writer, "%d\n", answer)

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
