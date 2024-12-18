package hackerrankexercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 * Complete the 'ModuloFibonacciSequence' function below.
 *
 * The function accepts following parameters:
 *  1. chan bool requestChan
 *  2. chan int resultChan
 */

// https://www.golangprograms.com/read-and-write-fibonacci-series-to-channel-in-golang.html
func ModuloFibonacciSequence(requestChan chan bool, resultChan chan int) {
	x, y := 1, 1
	mod := 1000000000
	// loop infinite
	for {
		if <-requestChan {
			// break condition
			time.Sleep(3 * time.Millisecond)
			x, y = y, (x+y)%mod
			resultChan <- x
		}
	}
}

func ModuloFibonacciSequencePointEntry() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	skipTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	skip := int32(skipTemp)

	totalTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	total := int32(totalTemp)

	resultChan := make(chan int)
	requestChan := make(chan bool)
	go ModuloFibonacciSequence(requestChan, resultChan)
	for i := int32(0); i < skip+total; i++ {
		start := time.Now().UnixNano()
		requestChan <- true
		new := <-resultChan
		if i < skip {
			continue
		}
		end := time.Now().UnixNano()
		timeDiff := (end - start) / 1000000
		if timeDiff < 3 {
			fmt.Println("Rate is too high")
			break
		}
		fmt.Println(new)
	}
}
