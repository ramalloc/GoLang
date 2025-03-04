package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Concurrency
// -> The ability to deal with multiple tasks simultaneously.
// -> Doesn't necessarily mean they run at the exact same moment but are managed so they appear to make progress.

// Parallelism
// -> Tasks are executed literally at the same time.
// -> Requires multiple CPU cores or threads.
// -> Concurrency is achieved through goroutines and the Go scheduler.
// -> Parallelism occurs when goroutines run simultaneously across multiple CPU cores.

var wg sync.WaitGroup
var mu sync.Mutex   // Mutex for synchronization


// Shared variable
type Counter struct {
	mu    sync.Mutex
	value int
}

var signal = []string{"test"}

func (c *Counter) Increment() {
	c.mu.Lock()  // Lock the mutex
	defer c.mu.Unlock() // Critical section: safe access to shared variable
	c.value++ // Unlock the mutex
}
func main() {
	// to use go routine we use -> "go" before any operation
	// go greeter("Hello")
	// greeter("World")

	// -> As we use go above we observed that go routine is not waiting the thread to return or complete the task
	// 	  it executes the next block to solve this problem below we have used time.sleep

	websiteList := []string{
		"https://google.com",
		"https://fb.com",
		"https://instagram.com",
	}

	for _, endpoint := range websiteList {
		go getStatusCode(endpoint)
		wg.Add(1)
	}
	// -> Using the above loop the api's are taking time to return data, we can make it faster by rolling different threads and
	//     fire goroutine there. But the problem is still there of thread returning.
	// -> So for the above issue we use sync method waitGroup which is advance version of time.sleep. Basically it contains
	// 		one manager, one indicators and one waiter, waitGroup..Add(1) - it manages the threads to complete using other
	// 		indicators, and waitGroup.Done() shows task is done and waitGroup.wait() - it waits to close all tasks/thread then
	// 		it finishes the main() or executes next block.
	wg.Wait()
	fmt.Println("Signal", signal)

	// In concurrent programming, multiple goroutines (or threads) might access and modify shared data simultaneously. This can
	// lead to race conditions, where the result depends on the timing or interleaving of executions, causing unpredictable behavior.

	// A mutex (short for mutual exclusion) is a synchronization primitive used to prevent race conditions and ensure safe
	//  access to shared resources in concurrent programming. It is provided by most programming languages, including Go.

	counter := &Counter{}

	for i := 0; i < 1000; i++ {
		go counter.Increment()
	}
	fmt.Println("Counter:", counter)
	// As we can see that each time counter value is different wuth time.sleep and using waitGroups it is 0 showing unpredictable
	// behaviour

	// therefore we use mutex to lock the storage of the shared resource for one threa until it completes it's task

}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error while hitting endpoint !")
		return
	} else {
		defer res.Body.Close()
		signal = append(signal, endpoint)
		fmt.Printf("%d status code by hitting %s endpoint\n", res.StatusCode, endpoint)
	}
}
