package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in for GO Routines...")
	// Channels are the way by which the go routines can communicate with each other

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}
	// myCh <- 5
	// fmt.Println(myCh)
	wg.Add(2)
	// GO ROUTINE GETTING VALUE FROM CHANNEL
	go func(myCh <-chan int, wg *sync.WaitGroup) {
		// fmt.Println(<- myCh)

		// close(myCh) -> no allowed in receiving channel

		// We do not pass any value into channel it will show by by default 0 and this can create issue as we don't know is
		// this value passed from any routine or not. So we can consume channel like this which return boolean and value
		val, isOpen := <-myCh
		fmt.Println("Channel Value - ", val)
		fmt.Println("Is Channel Opened - ", isOpen)

		// fmt.Println(<- myCh)
		defer wg.Done()
	}(myCh, wg)

	// To prevent inifite loop if channel is closing before consuming we define go routines according to to assign or consume
	// GO ROUTINE PUTTING VALUE INTO CHANNEL
	go func(myCh chan<- int, wg *sync.WaitGroup) {
		myCh <- 0

		// myCh <- 5
		// myCh <- 6
		// We can close the channel as well
		defer close(myCh)
		defer wg.Done()
	}(myCh, wg)

	wg.Wait()
}
