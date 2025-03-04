package main

import "fmt"

func main() {
	greet()
	fmt.Println("Starting Functions in GoLang...")

	result := adder(5, 55)
	fmt.Println("Result is : ", result)

	result1 := proAdder(1, 2, 3, 4, 5)
	fmt.Println("Result 1 - ", result1)

	result2, myMessage, name, tantra := proAdderWithString(1, 2, 3, 4, 5)
	fmt.Println("Result 2 - ", result2, myMessage, name, tantra[1])
}

// func func_name(variable type_of_variable) type_of_return_data {}
func adder(val1 int, val2 int) int {
	return val1 + val2
}

// We can add multiple value without defining multiple variables
func proAdder(values ...int) int {
	// the values we are getting here in arguments is a slice
	fmt.Println("Values - ", values)
	total := 0
	for _, value := range values{
		total += value
	}
	return total
}

func proAdderWithString(values ...int) (int, string, []string, map[int]string) {
	// the values we are getting here in arguments is a slice
	total := 0
	for _, value := range values{
		total += value
	}
	return total, "Roshan Kumar", []string{"Shri Ram", "Chandra"}, map[int]string{1: "Hari Om"}
}

func greet()  {
	fmt.Println("Namastey GoLang World")
}
