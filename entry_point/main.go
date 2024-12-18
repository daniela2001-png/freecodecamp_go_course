package main

import (
	"fmt"
	"freecodecamp_go_course/conditions"
	"freecodecamp_go_course/functions"
	"freecodecamp_go_course/structs"
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
}
