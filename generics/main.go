package generics

import (
	"errors"
	"fmt"
	"time"
)

//
/*
Imagine some code that splits a slice into 2 equal parts.
The code that splits does not care abpute the avlues stored in the slice

	**

	Unfortunately in Go we would need to write it multiple times for
	EACH TYPE !

	Which is a very un-DRY thing to do.
	**
*/

// For example:

// split into 2 a int slice
func splitIntSlice(s []int) ([]int, []int) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

// split into 2 a string slice
func splitStrSlice(s []string) ([]string, []string) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

// --- Type Parameters ---
/*

Generics allow us to use variables to refer to specific types.
This is an amazing feature because it allow us to write
abstract functions that drastically reduce code duplication.

*/

// For example, instead of write two different funtions, we can use generics as follows:
/*
	T is the name of the type parameter, that it must match the "any" constraint
	which means it can be anything.
*/
func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

// How  call it the previous function ?

// firstIntSlice, secodIntSlice := splitAnySlice([]int{1,2,3,4,5}) --> [1, 2], [3, 4, 5]

// --- Tip: Zero value of a type ---
/*
	Creating a variable that's the zero value of a type is easy:
		>> var myZeroInt int

	It's the same with generics, we just have a variable that represents the type:
		>> var myZero T
*/

// 1) Assignment:

/*

At Mailio we store all the emails for a campaign in memory as a slice.
We store payments for a single user in the same way.

Complete the getLast() function.
It should be a generic function that returns the last element from a slice, no matter the types stored in the slice.
If the slice is empty, it should return the zero value of the type.

*/

// solution
func getLast[T any](s []T) T {
	var lastElement T
	if len(s) == 0 {
		return lastElement
	}
	lastElement = s[len(s)-1]
	return lastElement
}

// --- Why Generics ? ---
/*
	- Generics reduce repetitive code
	- Generics are often used in libraries and packages.
*/

// --- Constraints ---
/*

	Sometimes you need the logic in your generic function to know something about the types it operates on.
	The example we used in the first exercise didn't need to know anything about the types in the slice, so we used the built-in any constraint:

	>>> func splitAnySlice[T any](s []T) ([]T, []T) {
    		mid := len(s)/2
    		return s[:mid], s[mid:]
		}

	Constraints are just interfaces that allow us to write generics that only operate within the constraints of a given interface type.
	In the example above, the any constraint is the same as the empty interface because it means the type in question can be anything.
*/

// Creating a custom constraint:
type errorManager interface {
	Error() string
}

func extractErrorsValues[T errorManager](sliceOfErrors []T) []string {
	msgs := make([]string, 0, len(sliceOfErrors)) // Init with zero values but with  a capacity, to prevent trash values at the begining of the slice
	for _, value := range sliceOfErrors {
		msgs = append(msgs, value.Error())
	}
	return msgs
}

// Calling extractErrorsValues:
// Create something that implements errorManager interface
type CustomError struct {
	msg string
}

func (ce *CustomError) Error() string {
	return ce.msg
}

func Solution() {
	sliceOfErrors := []*CustomError{
		{
			msg: "msg 1",
		},
		{
			msg: "msg 2",
		},
		{
			msg: "msg 3",
		},
	}
	msgValues := extractErrorsValues(sliceOfErrors)
	fmt.Println(msgValues) // ["msg 1", "msg 2", "msg 3"]
}

// 2) Assignment:
/*

	We have different kinds of "line items" that we charge our customer's credit cards for.
	Line items can be things like "subscriptions" or "one-time payments" for email usage.

	Complete the chargeForLineItem function. First, it should check if the user has a balance with enough funds to be able to pay for the cost of the newItem.
	If they don't, then return an "insufficient funds" error and zero values for the other return values.

	If they do have enough funds:

	- Add the line item to the user's history by appending the newItem to the slice of oldItems. This new slice is your first return value.
	- Calculate the user's new balance by subtracting the cost of the new item from their balance. This is your second return value.
*/
func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	newBalance := newItem.GetCost() - balance
	if newBalance < 0.0 {
		var oldItems []T
		return oldItems, 0.0, errors.New("insufficient funds")
	}
	oldItems = append(oldItems, newItem)
	return oldItems, newBalance, nil
}

// don't edit below this line

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}

// --- Interface type list ---
/*

	Is a "new" way to write interfaces

	We can now simply list a bunch of types to get a new interface/constraint

	When we should use interface type list  ?

	- When you know exactly which types satisfy your interface !
*/

type NumbersWith64Bits interface {
	// Syntax : the union notation A|B means “type A or type B”, and the ~T notation stands for “all types that have the underlying type T”
	~float64 | ~int // This will allow us exclusive/only to manage these types into our constraint inside of our generic function :D
}

// Using NumbersWith64Bits interface as a generic constraint:
func divideNumbers[T NumbersWith64Bits](x, y T) T {
	return x / y
}

func SolutionInterfaceTypeList() {
	// Calling divideNumbers using float64:
	divideNumbers(2.3, 4.78)

	// Calling divideNumbers using int:
	divideNumbers(30, 60)

	// trying calling divideNumbers using an unsigned integer we're gonna get an error
	// divideNumbers(uint(30), uint(60)) -> uint does not satisfy NumbersWith64Bits (uint missing in ~float64 | ~int)

}

// --- Parametric Constraints ---
/*
	Your interface definitions, which can later be used as constraints, can accept type parameters as well.
*/

// The store interface represents a store that sells products.
// It takes a type parameter P that represents the type of products the store sells.
type store[P product] interface {
	Sell(P)
}

type product interface {
	Price() float64
	Name() string
}

type book struct {
	title  string
	author string
	price  float64
}

func (b book) Price() float64 {
	return b.price
}

func (b book) Name() string {
	return fmt.Sprintf("%s by %s", b.title, b.author)
}

type toy struct {
	name  string
	price float64
}

func (t toy) Price() float64 {
	return t.price
}

func (t toy) Name() string {
	return t.name
}

// The bookStore struct represents a store that sells books.
type bookStore struct {
	booksSold []book
}

// Sell adds a book to the bookStore's inventory.
func (bs *bookStore) Sell(b book) {
	bs.booksSold = append(bs.booksSold, b)
}

// The toyStore struct represents a store that sells toys.
type toyStore struct {
	toysSold []toy
}

// Sell adds a toy to the toyStore's inventory.
func (ts *toyStore) Sell(t toy) {
	ts.toysSold = append(ts.toysSold, t)
}

// sellProducts takes a store and a slice of products and sells
// each product one by one.
func sellProducts[P product](s store[P], products []P) {
	for _, p := range products {
		s.Sell(p)
	}
}

func main() {
	bs := bookStore{
		booksSold: []book{},
	}

	// By passing in "book" as a type parameter, we can use the sellProducts function to sell books in a bookStore
	sellProducts[book](&bs, []book{
		{
			title:  "The Hobbit",
			author: "J.R.R. Tolkien",
			price:  10.0,
		},
		{
			title:  "The Lord of the Rings",
			author: "J.R.R. Tolkien",
			price:  20.0,
		},
	})
	fmt.Println(bs.booksSold)

	// We can then do the same for toys
	ts := toyStore{
		toysSold: []toy{},
	}
	sellProducts[toy](&ts, []toy{
		{
			name:  "Lego",
			price: 10.0,
		},
		{
			name:  "Barbie",
			price: 20.0,
		},
	})
	fmt.Println(ts.toysSold)
}

// Assignment:
/*

	The chief architect at Mailio has decided she wants to implement billing with generics.
	Specifically, she wants us to create a new biller interface. A biller is an interface that can be used to charge a customer, and it can also report its name.

	There are two kinds of billers:

		userBiller (cheaper)
		orgBiller (more expensive)

	A customer is either a user or an org.
	A user will be billed with a userBiller and an org with an orgBiller.

	Create the new biller interface. It should have 2 methods:

		Charge
		Name
	The good news is that the architect already wrote the userBiller and orgBiller types for us that fulfill this new biller interface.
	Use the definitions of those types and their methods to figure out how to write the biller interface definition.
*/
type biller[C customer] interface {
	Charge(C) bill
	Name() string
}

// don't edit below this line

type userBiller struct {
	Plan string
}

func (ub userBiller) Charge(u user) bill {
	amount := 50.0
	if ub.Plan == "pro" {
		amount = 100.0
	}
	return bill{
		Customer: u,
		Amount:   amount,
	}
}

func (sb userBiller) Name() string {
	return fmt.Sprintf("%s user biller", sb.Plan)
}

type orgBiller struct {
	Plan string
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller", ob.Plan)
}

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.0
	if ob.Plan == "pro" {
		amount = 3000.0
	}
	return bill{
		Customer: o,
		Amount:   amount,
	}
}

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount   float64
}

type user struct {
	UserEmail string
}

func (u user) GetBillingEmail() string {
	return u.UserEmail
}

type org struct {
	Admin user
	Name  string
}

func (o org) GetBillingEmail() string {
	return o.Admin.GetBillingEmail()
}

// --- Naming  Generic Types ---
/*
	Remember "T" is just a variable name, We could have named the type
	parameter anything. "T" happens to be a fairly common convention
	for a type variable :D
*/
