package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition in GoLang...")
	var score = []int{0}

	// Multiple Waitgroups
	wg := &sync.WaitGroup{}
	// Multiple Mutex
	mu := &sync.Mutex{}

	// wg.Add("No. of go routines")
	wg.Add(3)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		fmt.Println("1st Go Routine")
		mu.Lock()
		score = append(score, 1)
		mu.Unlock()
	}(wg, mu)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		fmt.Println("2nd Go Routine")
		mu.Lock()
		score = append(score, 2)
		mu.Unlock()
	}(wg, mu)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		fmt.Println("3rd Go Routine")
		defer wg.Done()
		mu.Lock()
		score = append(score, 3)
		mu.Unlock()
	}(wg, mu)

	wg.Wait()
	fmt.Println("score - ", score)

	// We saw above when we do not use mutex the score which is a memory location and shared between 3 go routines.
	// If there is 1st go routine writing in the memory then for same memory location 2nd or 3rd go routine try to perform
	// write operation in the memory location which leads to race condition for a memory write opeartion between go routines.
	// Therefore we use mutex to lock the resource when a go routine uses a resource/memory location.

}
