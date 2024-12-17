package conditions

import "fmt"

// if statements in Go do not use paretheses around conditions.
func Conditions() {
	/*
		Some operators in Go
			- == equal to
			- != not equal to
			- < less than
			- > greater than
			- <= less than or equal to
			- >= greater than or equal to
	*/
	messageLen := 10
	maxMessageLen := 20

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}

}

func getLength(email string) int {
	return len(email)
}

// The initial statement of an if block
/*
	if INITIAL_STATEMENT; CONDITION {
		// evaluate something
	}
*/
func InitialStatement() {
	email := "hello there!"

	// instead of writing:
	// lenght exists aroung all scope into the function
	length := getLength(email)
	if length < 1 {
		fmt.Println("email is invalid")
	}

	// we can do:
	// Removes length from the parent scope, we only need access to it while checking the condition
	if length := getLength(email); length < 1 {
		fmt.Println("email is invalid")
	} else {
		fmt.Println("email is valid")
	}
}
