package concurrency

// --- Concurrency ---
/*

	Concurrency  is the ability to executes multiple tasks at the same time
	that allow us to reduce the time execution of our programs.

	Concurrency is as simple as using the "go" keyword when calling
	a function:

	>>> go doSomething()

	In the example above, "doSomething()" will be executed concurrently

	The "go" keyword is used to create a new Go Routine.
*/

// Assignment:
/*

	At Mailio we send a lot of network requests. Each email we send must go out over the internet. To serve our millions of customers, we need a single Go program to be capable of sending thousands of emails at once.

	Edit the sendEmail() function to execute its anonymous function concurrently so that the "received" message prints after the "sent" message.
*/

import (
	"fmt"
	"sync"
	"time"
)

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

func SendEmailConcurrently() {
	test("Hello there Kaladin!")
	test("Hi there Shallan!")
	test("Hey there Dalinar!")
}

// --- Channels ---
/*

	Channels are a typed, thread-safe queue.
	Channels allow different goroutines to communicate with each other.

	** Create a channel:
		Like maps and slices, channels must be created before use. They also use the same make keyword:

		ch := make(chan int)
	** Send data to a channel:

		ch <- 69

		The <- operator is called the channel operator. Data flows in the direction of the arrow. This operation will block until another goroutine is ready to receive the value.

	** Receive data from a channel:

		v := <-ch

		This reads and removes a value from the channel and saves it into the variable v. This operation will block until there is a value in the channel to be read.

	** Blocking and deadlocks:

		A deadlock is when a group of goroutines are all blocking so none of them can continue.
		This is a common bug that you need to watch out for in concurrent programming.
*/

// Assignment:
/*
	Run the program.
	You'll see that it deadlocks and never exits.
	The sendIsOld function is trying to send on a channel, but no other goroutines are running that can accept the value from the channel.

	Fix the deadlock by spawning a goroutine to send the "is old" values.
*/
type Email struct {
	Body string
	Date time.Time
}

func CheckEmailAge(emails [3]Email) [3]bool {
	// Creating a channel of type bool
	isOldChan := make(chan bool)

	// sendIsOld sends values to channel isOldChan
	go sendIsOld(isOldChan, emails)

	// Receiving values from channel called isOldChan
	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan

	// return info channel
	return isOld
}

// sendIsOld is a function that ONLY sends info to the isOldChan channel
func sendIsOld(isOldChan chan<- bool, emails [3]Email) {
	for _, e := range emails {
		if e.Date.Before(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)) {
			// sending value as true to our isOldChan
			isOldChan <- true
			continue
		}
		// sending value as false to our isOldChan
		isOldChan <- false
	}
}

// --- Channels (TOKENS) ---
/*

Empty structs are often used as a unary value.
Sometimes, we don't care what is passed through a channel.
We care when and if it is passed.

We can block and wait until something is sent on a channel using the following syntax

	<-ch

This will block until it pops a single item off the channel, then continue, discarding the item.

*/

// Assignment:
/*

Our Mailio server isn't able to boot up until it receives the signal that its databases are all online, and it learns about them being online by waiting for tokens (empty structs) on a channel.

Run the code. It never exits! The channel passed to waitForDBs stays blocked, because it's only popping the first value off the channel.

Fix the waitForDBs function. It should pause execution until it receives a token for every database from the dbChan channel. Each time waitForDBs reads a token, the getDBsChannel goroutine will print a message to the console for you. The succinctly named numDBs input is the total number of databases. Look at the test code to see how these functions are used so you can understand the control flow.

*/

func WaitForDBs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		// for every db connection alive, we receive that event into our token or empt struct using dbChan channel
		<-dbChan
	}
}

func GetDBsChannel(numDBs int) (chan struct{}, *int) {
	ch := make(chan struct{})
	count := 0
	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()
	return ch, &count
}

// ---- Buffered Channels ----
/*

Channels can optionally be buffered.

Creating a channel with a buffer
You can provide a buffer length as the second argument to make() to create a buffered channel:

	ch := make(chan int, 100)

A buffer allows the channel to hold a fixed number of values before sending blocks.
This means sending on a buffered channel only blocks when the buffer is full, and receiving blocks only when the buffer is empty

*/

// Assignment:
/*

	We want to be able to send emails in batches.
	A writing goroutine will write an entire batch of email messages to a buffered channel, and later, once the channel is full, a reading goroutine will read all of the messages from the channel and send them out to our clients.

	Complete the addEmailsToQueue function.
	It should create a buffered channel with a buffer large enough to store all of the emails it's given. It should then write the emails to the channel in order, and finally return the channel.

*/
func AddEmailsToQueue(emails []string) chan string {
	buffChanEmails := make(chan string, len(emails)) // this is gonna be the queue with a fixed len
	for _, emailMsg := range emails {
		// send value into buffChanEmails buffer channel
		buffChanEmails <- emailMsg
	}
	return buffChanEmails

}

// SendEmail only reads the info from the buffer channel
func SendEmail(emails []string, emailMsg <-chan string) {
	for i := 0; i < len(emails); i++ {
		fmt.Printf("Sending email msg: %s\n", <-emailMsg)
	}

}

// An implementation of a queue using channels without  go routines
func ManageEmailsWithAQueue(emails []string) {
	buffChanEmails := AddEmailsToQueue(emails)
	SendEmail(emails, buffChanEmails)

}

// --- CLOSE CHANNELS ---
/*

	Channels can be explicitly closed by a **sender**:

		ch := make(chan int)
		// do some stuff with the channel
		close(ch)


	** Checking if a channel is closed: **

	Similar to the ok value when accessing data in a map, receivers can check the ok value when receiving from a channel to test if a channel was closed.

		v, ok := <-ch

		// ok is false if the channel is empty and closed.


	** DO NOT send on channel CLOSED !
	Sending on a closed channel will cause a panic.
	A panic on the main goroutine will cause the entire program to crash, and a panic in any other goroutine will cause that goroutine to crash.

	Closing isn't necessary.
	There's nothing wrong with leaving channels open, they'll still be **garbage collected if they're unused**.
	You should close channels to indicate explicitly to a receiver that nothing else is going to come across.

*/

// Assignment:
/*
	At Mailio we're all about keeping track of what our systems are up to with great logging and telemetry.

	The sendReports function sends out a batch of reports to our clients and reports back how many were sent across a channel.
	It closes the channel when it's done.

	Complete the countReports function. It should:

	- Use an infinite for loop to read from the channel:
	- If the channel is closed, break out of the loop
	- Otherwise, keep a running total of the number of reports sent
	- Return the total number of reports sent
*/

// countReports works as a receiver
func countReports(numSentCh <-chan int) int {
	numReports := 0
	for {
		reportsSend, ok := <-numSentCh
		if !ok { // if the channel is closed, we will break out of the loop
			break
		}
		numReports += reportsSend
	}
	return numReports
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}

func ManageReportsConcurrently() {
	numBatches := 10
	ch := make(chan int)
	go sendReports(numBatches, ch)
	totalReports := countReports(ch)
	fmt.Printf("--- The total of reports to send are: %d ---\n", totalReports)
}

// --- RANGE OVER A CHANNEL ---
/*
	Similar to slices and maps, channels can be ranged over.

	for item := range ch {
    	// item is the next value received from the channel
	}
	This example will receive values over the channel (blocking at each iteration if nothing new is there) and will exit only when the channel is closed.
*/

// Assignment:
/*
	It's that time again, Mailio is hiring and we've been assigned to do the interview.
	The Fibonacci sequence is Mailio's interview problem of choice.
	We've been tasked with building a small toy program we can use in the interview.

	Complete the concurrentFib function. It should:

		- Create a new channel of ints
		- Call fibonacci concurrently
		- Use a range loop to read from the channel and append the values to a slice
		- Return the slice
*/
func ConcurrentFib(n int) (serie []int) {
	chanInt := make(chan int)
	go fibonacci(n, chanInt)
	for value := range chanInt {
		serie = append(serie, value)
	}
	return serie
}

// don't touch below this line

// fibonacci will be the sender
func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x // send the actual value from fibonacci series.
		// "x" will have the actual value and "y" the next value
		x, y = y, x+y
	}
	close(ch)
}

// --- SELECT ---
/*

	Sometimes we have a single goroutine listening to multiple channels and want to process data in the order it comes through each channel.

	A select statement is used to listen to multiple channels at the same time. It is similar to a switch statement but for channels.

		select {
		case i, ok := <- chInts: // ok will be true if the channel is open for can readed it correctly !
			fmt.Println(i)
		case s, ok := <- chStrings:
			fmt.Println(s)
		}

	The first channel with a value ready to be received will fire and its body will execute.
	If multiple channels are ready at the same time one is chosen randomly.
	The ok variable in the example above refers to whether or not the channel has been closed by the sender yet.
*/

// Assignment:

/*
	Complete the logMessages function.
	Use an infinite for loop and a select statement to log the emails and sms messages as they come in order across the two channels.
	Add a condition to return from the function when one of the two channels closes, whichever is first.
	Use the logSms and logEmail functions to log the messages.
*/

func logMessages(chEmails, chSms chan string) {
	for {
		select {
		case i, ok := <-chEmails:
			if !ok {
				break
			}
			logEmail(i)
		case j, ok := <-chSms:
			if !ok {
				break
			}
			logSms(j)
		}
		break
	}
}

// don't touch below this line

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("Email:", email)
}

// --- Channels Review ---
/*

	* A declared but uninitialized channel is nil just like a slice
		var s []int       // s is nil
		var c chan string // c is nil

		var s = make([]int, 5) // s is initialized and not nil
		var c = make(chan int) // c is initialized and not nil

	* A send to a nil channel blocks forever
		var c chan string        // c is nil
		c <- "let's get started" // blocks

	* A receive from a nil channel blocks forever
		var c chan string // c is nil
		fmt.Println(<-c)  // blocks

	* A send to a closed channel panics
		var c = make(chan int, 100)
		close(c)
		c <- 1 // panic: send on closed channel

	* A receive from a closed channel returns the zero value immediately
		var c = make(chan int, 100)
		close(c)
		fmt.Println(<-c) // 0
*/

// Assignment:
/*
	Run the code as-is.
	You should see that it doesn't do anything interesting: no ping or pong messages are printed.

	Fix the bug in the pingPong function.

	Remember: if a program exits before its goroutines have completed, those goroutines will be killed silently.
	Which of the function calls probably shouldn't run in the background as a goroutine?
*/

func pingPong(numPings int) {
	pings := make(chan struct{})
	pongs := make(chan struct{})
	go pinger(pings, pongs, numPings)
	go ponger(pings, pongs)

	i := 0
	for range pongs {
		fmt.Println("got pong", i)
		i++
	}
	fmt.Println("pongs done")
}

// don't touch below this line

func pinger(pings, pongs chan struct{}, numPings int) {
	sleepTime := 50 * time.Millisecond
	for i := 0; i < numPings; i++ {
		fmt.Printf("sending ping %v\n", i)
		pings <- struct{}{}
		time.Sleep(sleepTime)
		sleepTime *= 2
	}
	close(pings)
}

func ponger(pings, pongs chan struct{}) {
	i := 0
	for range pings {
		fmt.Printf("got ping %v, sending pong %v\n", i, i)
		pongs <- struct{}{}
		i++
	}
	fmt.Println("pings done")
	close(pongs)
}

func PingPongConcurrency(numPings int) {
	fmt.Println("Starting game...")
	pingPong(numPings)
	fmt.Println("===== Game over =====")
}

// --- Mutexes ---
/*

	Mutexes in Go (prevents race conditions)
	- Mutexes allow us to lock access to data.
	This ensures that we can control which goroutines can access certain data at which time.

	Go's standard library provides a built-in implementation of a mutex with the sync.Mutex type and its two methods:

		- .Lock()
		- .Unlock()
	We can protect a block of code by surrounding it with a call to Lock and Unlock as shown on the protected() method below.

	It's good practice to structure the protected code within a function so that defer can be used to ensure that we never forget to unlock the mutex.

		func protected(){
			mu.Lock()
			defer mu.Unlock()
			// the rest of the function is protected
			// any other calls to `mu.Lock()` will block
		}

	Mutexes are powerful. Like most powerful things, they can also cause many bugs if used carelessly.

	** Maps are not thread-safe **
	- Maps are not safe for concurrent use! If you have multiple goroutines accessing the same map, and at least one of them is writing to the map, you must lock your maps with a mutex.

*/

// Assignment:

/*
We send emails across many different goroutines at Mailio.
To keep track of how many we've sent to a given email address, we use an in-memory map.

Our safeCounter struct is unsafe! Update the inc() and val() methods so that they utilize the safeCounter's mutex to ensure that the map is not accessed by multiple goroutines at the same time.
*/
type safeCounter struct {
	counts map[string]int
	mu     *sync.Mutex
}

func (sc safeCounter) inc(key string) {
	// lock until the map is updated
	sc.mu.Lock()
	// unlock function with the purpose that another go routine takes it
	defer sc.mu.Unlock()
	// update value from map by key
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	// lock until the map is readed
	sc.mu.Lock()
	// unlock function with the purpose that another go routine takes it
	defer sc.mu.Unlock()
	// gets map value by key
	return sc.slowVal(key)
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func (sc safeCounter) slowVal(key string) int {
	time.Sleep(time.Microsecond)
	return sc.counts[key]
}
