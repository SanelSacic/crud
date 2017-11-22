package main

import "fmt"

func main() {
	signalAck()
}

// signalAck shows how to signal an event and wait for an
// acknowledgment it is done,
func signalAck() {

	ch := make(chan string)
	go func() {
		fmt.Println(<-ch)
		ch <- "ok done"
	}()

	ch <- "do this"
	fmt.Println(<-ch)
}
