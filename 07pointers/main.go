package main

import "fmt"

func main() {
	// ---> Pointers - It is a reference to a direct memory location.
	// Pointers solve a problem - When we declare any variable it stores in memory and as we use that variable in multiple
	// functions passing the copies of that variable in multiple things as it can cause some irregularities in the program.
	// So therefore we use pointers, pointers point to the memory address so we use or pass pointers in functions or wants
	// to pass actual value so that we can remove that irregularities in program.

	fmt.Println("Let's learn pointers...")
	// var ptr *int
	// fmt.Println("Value of empty *ptr :- ", ptr)
	// this will return <nil>

	myNum := 435
	var ptr = &myNum
	fmt.Println("ptr value without referance :- ", ptr)
	// ptr value without referance :-  0xc0000120b0

	fmt.Println("ptr value with referance :- ", *ptr)
	// ptr value with referance :-  435

	*ptr = *ptr * 2
	fmt.Println("Added actual value thorugh pointer :- ", myNum)
	// Added actual value thorugh pointer :-  870

}
