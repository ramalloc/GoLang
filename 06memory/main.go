package main

func main()  {
	// ----> Two Ways to allocate Memory in goLang
	// 1. new()
		// -> Allocate memory but not initialised and we can use that further as it gives memory address
		// -> It returns zeroed storage in which initially we cannot put any data.

	// 2. make()
		// -> Allocate memory and initialise it as well and we can use that further as it gives memory address
		// -> It returns non-zeroed storage in which initially we can put any data.

	// -> Garbage Collection happens automatically in GoLang. Where anything will go out of scope will be nil.
}