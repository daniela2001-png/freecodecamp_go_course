package structs

import "fmt"

type Wheel struct {
	Radius   int
	Material string
}
type Car struct {
	Model      string
	Height     string
	Widht      string
	FrontWheel Wheel // structs supports more than primary data types
	BackWheel  Wheel
}

type MessageToSend struct {
	Message   string
	Sender    User
	Recipient User
}

type User struct {
	Name   string
	Number int
}

func CanSendMessage(mToSend MessageToSend) (isTrue bool) {
	if mToSend.Recipient.Name != "" {
		return true
	}
	if mToSend.Sender.Name != "" {
		return true
	}
	if mToSend.Recipient.Number != 0 {
		return true
	}
	if mToSend.Sender.Number != 0 {
		return true
	}
	return
}

// --- ANONYMOUS STRUCTS ---
/*
They are just like normal sturct,
but its defined without a name and therefore can not
be referenced elsewhere in the code.

Main Reason for can create a anonymous struct:

	- If you have no reason to create more than one instance
		of the struct.
*/

func AnonymousStructs() {
	myCar := struct {
		Make  string
		Model string
		// we can nest anonymous structs into anothers anonymous structs
		Wheel struct {
			Radius   int
			Material string
		}
	}{
		Make:  "tesla",
		Model: "model 3",
	}
	fmt.Println(myCar)
	// You can even nest anonymous structs as fields within other structs:

}

// --- EMBEDDED STRUCTS ---
/*
Embedded structs are a way to compose structs by including
the fields of one struct in another struct.

This is similar to inheritance in object-oriented programming
languages, but with some important differences.

In Go, there is no inheritance, but you can achieve similar
functionality by using embedded structs.

*/

type Animal struct {
	name   string
	colour string
}

type Dog struct {
	Animal    // Here we do not need to access to animal struct first for can get the field of animal's struct, we can do it directly, because its embedded not nested !
	barkSound string
}

func embeddedStructs() {
	animal := Animal{
		name:   "Doggy",
		colour: "brown",
	}
	dog := Dog{
		Animal:    animal,
		barkSound: "wooof",
	}
	fmt.Printf("The embedded dog has a %s colour and the barsk sounds like this: %s", dog.colour, dog.barkSound)
}

type DogNested struct {
	barkSound string
	animal    Animal // Here we need to access to animal struct first for can get the field of animal's struct, because it is nested !
}

func nestedStructs() {
	animal := Animal{
		name:   "Doggy",
		colour: "brown",
	}
	dogv2 := DogNested{
		barkSound: "wooaf",
		animal:    animal,
	}

	fmt.Printf("The nested dog has a %s colour and the barsk sounds like this: %s", dogv2.animal.colour, dogv2.barkSound)
}

// ---- STRUCT METHODS IN GO ----
/*

	Methods are just functions that have a receiver,
	where a receiver  is a special parameter that goes
	before the name of the function.
*/
type authenticationInfo struct {
	username string
	password string
}

func (auth authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", auth.username, auth.password)
}
