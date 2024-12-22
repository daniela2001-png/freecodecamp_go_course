package main

import (
	"fmt"
	"time"

	"github.com/daniela2001-png/freecodecamp_go_course/concurrency"
	"github.com/daniela2001-png/freecodecamp_go_course/conditions"
	"github.com/daniela2001-png/freecodecamp_go_course/functions"
	"github.com/daniela2001-png/freecodecamp_go_course/pointers"
	"github.com/daniela2001-png/freecodecamp_go_course/slices"
	"github.com/daniela2001-png/freecodecamp_go_course/structs"
)

func main() {
	// way to declare multiple variables in go
	milage, company := 85450, "Tesla"

	// constants type does not support short declaration syntax ":="
	// we can concatenate constant values
	const name = "daniela"
	const surname = "morales"
	const fullname = name + surname

	fmt.Printf("My fullname is : %s", fullname)

	// formatting strings in go
	// fmt.Printf - Prints a formatted string to standard output
	// fmt.Sprintf - Returns the formatted string

	// Interpolate a string
	fmt.Printf("I am %s years old\n", "13")

	// Interpolate an integer
	fmt.Printf("I am %d years old\n", 13)

	// Interpolate a decimal
	fmt.Printf("I am %f years old\n", 10.534)

	// The ".2" rounds the number to 2 decimal places
	fmt.Printf("I am %.2f years old\n", 10.559)

	// Using Sprintf:
	const firstName = "Saul"
	const openRate = 30.5
	msg := fmt.Sprintf("Hi %s your open rate is %.10f percent", firstName, openRate)

	fmt.Println(msg)

	fmt.Printf("The milages is equal to: %d and company name is: %s\n", milage, company)

	// call Conditions function from conditions package
	conditions.Conditions()

	// call InitialStatement function from conditions package
	conditions.InitialStatement()

	// Using callbacks in Go:
	// Call MyCallBack:
	result := functions.MyCallBack(functions.AddV2, 200, 400)
	fmt.Printf("The result from MyCallBack is: %d", result)

	// --- STRUCTS ---
	// The fields of a struct can be accessed using the dot "." operator

	myCar := structs.Car{} // all fields are will be initialized with default/zero values
	myCar.Model = "dummy_model"

	fmt.Println(myCar)

	msgToSend := structs.MessageToSend{}
	msgToSend.Message = "my message"
	msgToSend.Sender = structs.User{
		Name:   "name_user_sender",
		Number: 123,
	}
	msgToSend.Recipient = structs.User{
		Name:   "name_user_recipient",
		Number: 0,
	}
	if structs.CanSendMessage(msgToSend) {
		fmt.Println("there was error sending msg")
	} else {
		fmt.Println("we can send correctly the message")
	}
	var costsSlice = []slices.Cost{
		{0, 4.0},
		{1, 2.1},
		{1, 3.1},
		{5, 2.5},
	}
	final := slices.GetCostsByDay(costsSlice)
	fmt.Println(final, len(final))

	matrix := slices.CreateMatrix(5, 10)
	fmt.Println(matrix, len(matrix))

	// pointers example:
	msgPtr := "shoot dang something else heck"
	messagePtr := &msgPtr
	pointers.RemoveProfanity(messagePtr)
	fmt.Println(msgPtr) // ****** ****** something else ******

	// -- Concurrency --

	// first solve problem
	concurrency.SendEmailConcurrently()

	// 2d problem solved using channels:
	emails := [3]concurrency.Email{
		{
			Body: "body one",
			Date: time.Now().UTC(),
		},
		{
			Body: "body two",
			Date: time.Now().UTC(),
		},
		{
			Body: "body three",
			Date: time.Date(2019, 12, 31, 9, 0, 0, 0, time.UTC),
		},
	}
	thereExistsOlds := concurrency.CheckEmailAge(emails)
	fmt.Printf("There are  old people ? : %v", thereExistsOlds) // [false, false, true]

	// 3th problem using tokens channels:
	numsDB := 5
	tokensChan, numbActiveDBs := concurrency.GetDBsChannel(numsDB)
	concurrency.WaitForDBs(numsDB, tokensChan)
	fmt.Printf("The total number of online or active db's are equal to: %d\n", *numbActiveDBs)

	// 4th problem using buffered channels:
	batchEmail := []string{
		"Hi there What is up",
		"Salve !",
		"Hallo !",
	}
	concurrency.ManageEmailsWithAQueue(batchEmail)

	// problem #5 using the validation when a channel is closed:
	concurrency.ManageReportsConcurrently()

	// fibonacci using concurrency:
	upTo := 8
	serie := concurrency.ConcurrentFib(upTo)
	fmt.Println(serie)

	// pingpong concurrency:
	concurrency.PingPongConcurrency(5)

}
