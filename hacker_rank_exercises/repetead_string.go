package hackerrankexercises

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
	numberOfA := int64(0)
	counter := 0
	length_str := int64(len(s))
	number_of_substr := (n / length_str)
	for _, value := range s {
		if string(value) == "a" {
			counter += 1
		}
	}
	reminder := 0
	reminder = int(n) - (int(length_str) * int(number_of_substr))
	fmt.Println(reminder)
	strReminderA := s[:reminder]
	counterreminder := 0
	for _, value := range strReminderA {
		if string(value) == "a" {
			counterreminder += 1
		}
	}
	numberOfA = (int64(counter) * int64(number_of_substr)) + int64(counterreminder)
	return numberOfA

}

func repeatedStringEntryPoint() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

	fmt.Fprintf(writer, "%d\n", result)

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
