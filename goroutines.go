package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
	}
}

// A function that returns value a string to channel after one second
func fun1(channel1 chan string) {
	time.Sleep(1 * time.Second)
	channel1 <- "Hi 🙂"
}

// A function that returns value a string to channel after two second
func fun2(channel2 chan string) {
	time.Sleep(2 * time.Second)
	channel2 <- "hello 🤣"
}

// Main function 🥰
func main() {

	// Normal display of goroutine
	go display("hi") // this one needs a little time to be called since it is a go routine so a goroutine has to be created so later function is run before it as its o/p is stored in buffer
	display("bye")   // this little no time to be called as it is normal function

	time.Sleep(1 * time.Second) // we need a little time to sleep because we have to prine what is in buffer as both above function were called at almost same time so one got printed while other is waiting to be printed. If below line is executed at exactly same time print buffer will be cleared and thus above function called by goroutine will not print its result as buffer was overwritten..
	display("bye")

	// making channel through which information can be transferred
	ch1 := make(chan string)
	ch2 := make(chan string)

	// both goroutines started in seperate thread provided by go
	go fun1(ch1)
	go fun2(ch2)

	// select returns(runs) whichever gives outpus faster
	select {
	case option1 := <-ch1:
		fmt.Println(option1) //this one gives output faster here as this one sleeps for about one second only so gives output through its channel faster and thus it this case runs...
	case option2 := <-ch2:
		fmt.Println(option2)
		// default:
		// fmt.Println("hello multiverse🌎🌍🌏")
		// if no case ready then default case is called
	}
	ch3 := make(chan string)

	go fun1(ch3)
	// since both channel return at same time so any at random is called.
	select {
	case a := <-ch3:
		fmt.Println(a, "1")
	case b := <-ch3:
		fmt.Println(b, "2")
	}

	// closing all channels so that no data shall be transfered from here on.
	close(ch1)
	close(ch2)
	close(ch3)
}
