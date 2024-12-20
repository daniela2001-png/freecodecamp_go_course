package slices

import "fmt"

// --- Array ---
/*

	Arrays in golang  have fixed size and will contains
	values of the SAME TYPE, for example:
	[3] int -> will be an array with three integers in total
*/

// --- Slice ---
/*
	A slice is a dinamically sized flexible
	view into an array, slices are built on top of arrays

	- ** Slices wrap arrays ** -> slices references arrays

	[]int -> an slice of type int

*/

// --- make function slices ---
/*

	We can create direclty a slice wihtout an array using
	the "make" function:

	func make([]T, len, cap) []T :

	- []T -> The type values that we're gonna store

	- len -> The initial  length of the slice

	- The cap -> capacity is kind of the total space that we have to grow
	the slice until we need to allocate a new address memory

	Note -> If we do not define the capacity the default value that is gonna take will be the same as the len

		mySlice := make([]int, 5, 10)

*/
// Example using make:

func getMessageCosts(messages []string) []float64 {
	msgCostSlice := make([]float64, len(messages))
	for _, msg := range messages {
		msgCostSlice = append(msgCostSlice, float64(len(msg))*0.01)
	}
	return msgCostSlice
}

// --- Variadic Functions ---
/*
	Are those that can take an arbitrary number of final arguments

	using the "..." syntax in the function signature.
*/

// For example:

func sum(nums ...int) int {
	// nums is just a slice
	var num int
	for i := 0; i < len(nums); i++ {
		num += nums[i]
	}
	return num
}

// --- Spread Operator ---
/*
	The spread operator allow us to pass a slice into a variadic function.
	The spread opeartor consists of three dots folloing the slice
	in the function call.
*/

func slicesEntryPoint() {
	// ------------------ Array ---------------------
	// For example, we can initialized an array wiht zero values, as follows:
	var myInts [10]int

	// Or we can directly initialized:
	toSix := [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println(myInts, toSix)

	// --------------- Slice ------------------------

	// completeCopySlice := array[:]
	mySlice := toSix[1:4] // [start_index: end_index -1] -> [2, 3, 4]
	fmt.Println(mySlice)

	// calling sum variadic function:
	total := sum(1, 2, 3)
	fmt.Println(total)

	// SPREAD OPERATOR:
	numsToCalculate := []int{1, 2, 3, 4, 5}
	_ = sum(numsToCalculate...) // is way to pass a slice directly into variadic funtion (spread operator)

}

// Example:
type Cost struct {
	Day   int
	Value float64
}

func GetCostsByDay(costs []Cost) []float64 {
	var max int
	for i := 0; i <= len(costs)-1; i++ {
		if costs[i].Day > max {
			max = costs[i].Day
		}
	}
	final := make([]float64, max+1)
	for y := 0; y < len(costs); y++ {
		final[costs[y].Day] += costs[y].Value
	}
	return final
}

// --- Slice of slices ---
// slices can hold other slices, creating matrix or a 2D slice.
// rows := [][]int{}
func CreateMatrix(rows, cols int) [][]int {
	var matrix [][]int
	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}
