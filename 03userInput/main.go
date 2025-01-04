package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome := "Taking Input from User..."
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	// Comma ok || Comma err Syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for your input rating  - ", input)
	fmt.Printf("Type of your input rating  - %T \n", input)

	// -> Ok or Err
	message := "Enter Value for ok or err..."
	fmt.Println(message)
	ok, err := reader.ReadString('\n')
	fmt.Println("Ok rating  - ", ok)
	fmt.Println("Error in rating  - ", err)
	
}