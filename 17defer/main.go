package main

import "fmt"

func main() {
	// When we use defer before any code or execution, program cuts that code and place it just before of last curly braces.
	// As we use defer before any code that code will execute in the last of execution.

	// If there are multiple defer codes/progrmas then the LIFO happen last defer code executed till first in this LIFO order.
	defer fmt.Println("World")
	defer fmt.Println("First")
	defer fmt.Println("The")
	fmt.Println("Hello")

	// Hello
	// The
	// First
	// World

	myDefer()
	//	5
	//  4
	//  3
	//  2
	//  1
}

func myDefer() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
