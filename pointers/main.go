package pointers

import (
	"fmt"
	"strings"
)

/*
A pointer is variable that stores a memory address of another variable

x := 5
z := &x -> point to memory address of x

*z = 6 -> change the value of x from 5 to 6

Reasons to implements pointers in our code:

-1) We can optimize the use of the memory, because by default
all variable are passed by value, that means the we always getting a copy
instead, we can use pointers to pass by reference using directly the value to update or use
and optimize memory.

2) we  can directly change the value of our variable with pointers.
*/

func entryPointpointers() {
	mystring := "hi"
	mystringPtr := &mystring

	fmt.Println(*mystringPtr) // read mystring through the pointer
	*mystringPtr = "there"    // set mystring through the pointer.
}

/*
Complete the removeProfanity function.

It should use the strings.ReplaceAll function to replace all
instances of the following wrods in the input message with asterisks.

- "dang" -> ****
- "shoot" -> ****
- "heck" -> ****

It should mutate the value in the pointer and return nothing
*/
func RemoveProfanity(message *string) {
	if message == nil {
		return
	}
	// update the param value without return it using pointer
	*message = strings.ReplaceAll(*message, "dang", "******")
	*message = strings.ReplaceAll(*message, "shoot", "******")
	*message = strings.ReplaceAll(*message, "heck", "******")
}

// --- Pointer receiver ---
/*
	Methods with pointer receivers can modify the value to which the receiver
	Since methods often need to modify their receiver, pointer receives
	are more common than value receivers.

*/
// Example with pointer receiver:
type hair struct {
	color string
}

func (h *hair) setColor(newColor string) {
	h.color = newColor
}

func changeColor() {
	h := hair{
		color: "brown",
	}
	h.setColor("black")
	fmt.Println(h.color) // h.color black
}

// Example with non-pointer receiver or value receiver:
func (h hair) setColorV2(newColor string) {
	h.color = newColor
}

func changeColorV2() {
	h := hair{
		color: "brown",
	}
	h.setColor("black")
	fmt.Println(h.color) // h.color brown
}
