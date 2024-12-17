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
	}{
		Make:  "tesla",
		Model: "model 3",
	}
	fmt.Println(myCar)
	// You can even nest anonymous structs as fields within other structs:

}
