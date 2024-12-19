package interfaces

import (
	"fmt"
	"math"
)

// --- INTERFACES ---
/*

	An interface is a collection of method signatures
	A type "implements" an interface if it has all of the methods
	of the given interface defined on it.

	Interfaces are implemented implicitily.
	That means that we do not need to specify that a type
	implements an interface, because Go does it automatically.
*/
type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perimeter() float64 {
	return 2*r.height + 2*r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// In the previous example rect and circle implements the shape interface

// Also we can pass interfaces as parameters into methods for example:
type message interface {
	getMessage() string
}

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type mailMessage struct {
	recipient string
	sender    string
}

func (mailmsg mailMessage) getMessage() string {
	return fmt.Sprintf("Sending mail msg from %s to %s", mailmsg.sender, mailmsg.recipient)
}

type SMSMessage struct {
	phoneNumber int
}

func (smsmsg SMSMessage) getMessage() string {
	return fmt.Sprintf("Sending SMS msg to phone number %d", smsmsg.phoneNumber)
}

// calling sendMessage() method:
func interfaceEntryPoint() {
	mailMsg := mailMessage{
		recipient: "pepe",
		sender:    "maria",
	}
	smsMsg := SMSMessage{
		phoneNumber: 123456,
	}
	sendMessage(smsMsg)
	sendMessage(mailMsg)
}

// --- Interface Implementation ---
type employee interface {
	getName() string
	getSalary() int
}

// contractor employees
type contractor struct {
	name          string
	hourlyPay     int
	hourlyPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hourlyPerYear
}

// full-time employees
type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getName() string {
	return ft.name
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

// --- multiple interfaces ---
/*
	A type can implemnt multiples interfaces in Go
	For example, the empty interface, interface {}, is always
	implemented by every type because it has no requirements.
*/

type Study interface {
	totalReadBooks() int
	getDegree() string
}

type Work interface {
	getCareer() string
	getYearsOfExperience() int
}

// a Person type implements Work and Study interfaces !
type Person struct {
	daysWorking int
	degree      string
	career      string
}

func (p Person) getDegree() string {
	return p.degree
}
func (p Person) totalReadBooks(booksReadedByYear, totalYears int) (totalBooks int) {
	totalBooks = booksReadedByYear * totalYears
	return
}

func (p Person) getYearsOfExperience() int {
	return (p.daysWorking * 1) / 365
}

func (p Person) getCareer() string {
	return p.career
}

// --- Name your interface arguments ---
/*
	Add names to our interface arguments allows us to make more clear
	respect to what we are receiving and what we are returning.

*/

// this is a bad practice, what does it means (string, string)  at upload method?
type Uploader interface {
	Upload(string, string) bool
}

// Instead we can name our interface arguments, as follows:
type NamedUploader interface {
	Upload(remoteFilePath string, localFilePath string) (isSuccess bool)
}

// --- Type Assertions in Go ---
/*
	Working with interfaces, every once-in-awhile you will need
	to access to the underlying type of an interface.

	**You can cast an interface to its underlying type using
	"type assertion"**

*/
// for example:
func assertionTypes() {
	myCircle := circle{
		radius: 5,
	}
	var myInstanceInterface shape = myCircle

	c, ok := myInstanceInterface.(circle) // should return c = circle instance and ok = true
	fmt.Printf("The type of c is: %T, is a circle ? : %v", c, ok)

	r, ok := myInstanceInterface.(rect) // should return r = nil and ok = false, because myInstanceInterface it is not a concrete rectangle type
	fmt.Printf("The type of r is: %T, is a circle ? : %v", r, ok)
}

// example #2:
/*

	Implement the getExpenseReport function
	if the expense is an email, then it should return an email info `+ cost
	else if expense == sms return the sms info + cost
	else return en "" && 0.0 for the cost
*/
type expense interface {
	cost() float64
}

type invalid struct{}

func (i invalid) cost() float64 {
	return 0.0
}

type email struct {
	toAddress    string
	body         string
	isSubscribed bool
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

type sms struct {
	toPhoneNumber string
	body          string
	isSubscribed  bool
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func getExpenseReport(e expense) (expenseTypeInfo string, cost float64) {
	// assertions types
	if em, ok := e.(email); ok {
		return em.toAddress, em.cost()
	}
	if s, ok := e.(sms); ok {
		return s.toPhoneNumber, s.cost()
	}
	return
}

// --- TYPE SWITCHES ---
/*

	A type switch makes it easier to do several type assertions
	in a series.
*/

// for example, following the previous logic in getExpenseReport
func getExpenseReportWithSwitch(e expense) (expenseTypeInfo string, cost float64) {
	// assertions types
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return
	}
}

// --- CLEAN INTERFACES ---
/*

	There exists some rules to keep clean interfaces in our code:

	1) Keep interfaces small: That means that interfaces should
	define the minimal behavior necessary to represent an idea or concept.

		- Interfaces should have as FEW methods as possible

	2) Interfaces should have NO knowledge of satisfying types: An interface
	should define what is necessary for other types to classify as a member
	of that interface.

		- In a brief words, interfaces does not know the concrete types that implements
		but concrete types can know in a implicity way their interfaces!

	3) Interfaces are NOT classes:
		- Interfaces do not have constructors
		- Interfaces define function signatures, but NO underlying behavior.
*/

// As an example of point number 1 (Keep interfaces small)

// This interface has only one behavior that is read a file for example
// DO
type Reader interface {
	Read(p []byte) (n int, err error)
}

// This interface has only one behavior that is write a file for example
type Writer interface {
	Write(p []byte) (n int, err error)
}

// An against example will be an interface with many responsibilities:
// DON'T
type FileProcessor interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
	Delete() (err error)
	Close() (err error)
	Rename(newName string) (err error)
}

// As an  againts example of point number 2 (Interfaces should have NO knowledge of satisfying types)

// This interface is linking their concrete types to its definition, this is WRONG
type ComputerManager interface {
	turnOn() bool
	turnOff() bool
	getOS() string
	isMACOS() bool   // this is a anti-pattern, why all the computer manager should validate if a computer has MAC OS or does not ?
	isLinuxOS() bool // this is also another anti-pattern, why all the computer manager should validate if a computer has Linux OS or does not ?
}

// A good example of point number 2 is using type assertions:  (Interfaces should have NO knowledge of satisfying types)
type ComputerManagerV2 interface {
	turnOn() bool
	turnOff() bool
	getOS() string
}

type LinuxComputer struct {
	isOn         bool
	isOff        bool
	architecture string
	nameOS       string
	yearsOfUse   int
}

func (c LinuxComputer) turnOn() bool {
	return !c.isOn
}

func (c LinuxComputer) turnOff() bool {
	return !c.isOff
}

func (c LinuxComputer) getOS() string {
	return c.nameOS
}

type MacComputer struct {
	LinuxComputer
}

func (c MacComputer) turnOn() bool {
	return !c.isOn
}

func (c MacComputer) turnOff() bool {
	return !c.isOff
}

func (c MacComputer) getOS() string {
	return c.nameOS
}

// manage type assertions is a way to solve if i want to know the implemented interface of a concrete type
func manageTypesOfOS(c ComputerManagerV2) (msg string) {
	switch v := c.(type) {
	case MacComputer:
		return fmt.Sprintf("we have a %s computer", v.getOS())
	case LinuxComputer:
		return fmt.Sprintf("we have a %s computer", v.getOS())
	default:
		return
	}
}
