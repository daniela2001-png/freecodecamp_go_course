package functions

import "errors"

/*
Functions in Go can take zero or more arguments

# To make Go code easier to read, the variable type comes after the variable name

For example, the following function:
*/
func AddV1(x int, y int) int {
	return x + y
}

// Here Add(x int, y int) int is know as the "function signature".

// When we have multiple arguments with the same type, we can declare functions as follows:
func AddV2(x, y int) int { return x + y }

// --- CallBacks ---
/*
	Also in go, we can implements callbacks (functions that takes as a parameters another functions)
	the callback is gonna be "myFunc"
*/
func MyCallBack(myFunc func(int, int) int, x, y int) int {
	if result := myFunc(x, y); result < x*y {
		return result
	} else {
		return 0
	}
}

// --- Named return values ---
/*
	A return statement without arguments returns the named return values.
	This is know as a "naked" return.
	Naked return statements *** should be used only in short functions ***
*/
func GetCoords() (x, y int) {
	// x and y are initialized with zero values

	return // automatically returns x and y
}

// --- Early returns ---
/*

Go supports the ability to return early from a function.

This is apowerful feature that can clean up code, especially when
used as a **guard clauses**.

That means that instead of use If/Else chains, we just return early
from the function at the end of each conditional block.


*/
func Divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		// error handling using a guard clause (an early return when this condition is true)
		return 0, errors.New("can not divide by zero")
	}
	return dividend / divisor, nil
}
