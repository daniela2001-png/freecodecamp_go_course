package errors

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// --- Error Handling ---
/*

	Golang allow us to manage individually each type of error
	when we call differents methods into our program.
	Comparing with languages like JS that creates nested try-catch
	for can validate differents types of errors due to the scope that
	this try-catch creates.

	NOTE -> When you return a non-nil error in Go, it's conventional
	to return the "zero" values of all other return values.

	Go programs express errors with error values. An Error is any type
	that implements the simple built-in error interface:
*/
type error interface {
	Error() string
}

func convertToInteger(stringNumber string) int {
	integer, err := strconv.Atoi(stringNumber)
	if err != nil {
		// return zero value for int type
		return 0
	}
	return integer
}

// --- Custom Error Types ---
/*

	Due to an type error is at the end just an interface
	we can make our own implementations, "through custom errors"
	as follows:
*/

// example # 1
var canSendMsg bool

type userError struct {
	reason string
}

func (u userError) Error() string {
	return fmt.Sprintf("there was an error with your account: %s", u.reason)
}

func sendSMS(msg string) error {
	if !canSendMsg {
		// Call custom error
		return userError{
			reason: "failed to send message.",
		}
	}
	return nil
}

type invalidOperationError struct {
	reason string
}

func (io invalidOperationError) Error() string {
	return fmt.Sprintf("invalid operation, the reason was: %s", io.reason)
}

// example #2
func getRootSquare(radicand int) (root float64, err error) {

	if radicand < 0 {
		return 0.0, invalidOperationError{
			reason: "the root square must be a real number.",
		}
	}
	return math.Sqrt(float64(radicand)), nil
}

// --- The errors package ---
/*

	The go standard library provides an "errors" package
	that makes easier deal with errors.

	As the errors.New() method does
*/

// here is not needed to create a struct that implements  th error interface, instead we can create an error directly.
var err error = errors.New("my message error")
