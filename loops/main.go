package loops

import "fmt"

// --- Loops ---
/*

	Loops in Go is written in standard C-like syntax:

	for INITIAL; CONDITION; AFTER {
		do something
	}

	Where:

	- INITIAL is run once at the beginning of the loop
	and can create variables within the scope of the loop.

	- CONDITION is checked before each iteration. If the condition
	does not pass then the loop breaks.

	- AFTER is run after each iteration.
*/

// For example:

func loopsEntryPoint() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func bulkSend(numMessages int) (totalCost float64) {
	for i := 0; i <= numMessages; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}
	return totalCost
}

// --- Omitting conditions from a for loop ---
/*

	Loops can omit sections of a for loop. For example, the
	CONDITION can be omitted which causes the loop to run forever.

	for INITIAL;; AFTER {
		do something forever
	}
*/

// For example:
// maxMessages is gonna calculate the maximum number of messages that can be sent
func maxMessages(thresh float64) (totalMessages int) {
	var currentCost float64
	for i := 0; ; i++ {
		currentCost += 1.0 + (0.01 * float64(i))
		if currentCost > thresh {
			return i
		}
	}
}

// --- There is no while loop in go ---
/*

	In Go a while loop is just a for loop that ONLY has a CONDITION:

	for CONDITION {
		do some stuff while CONDITION is true
	}
*/

// For example:
func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) (maxMessagesToSend int) {
	actualCostInPennies := 1.0
	//  for loop as a while loop
	for actualCostInPennies <= float64(maxCostInPennies) {
		maxMessagesToSend++
		actualCostInPennies *= costMultiplier
	}
	return
}
