package main

import "fmt"

func main() {
	fmt.Println("Array in GO...")

	var fruitList[4] string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Orange"

	fmt.Println("FruitList :- ", fruitList)
	fmt.Println("Length of FruitList :- ", len(fruitList))

	var vegList = [3]string{"Potato", "Ginger", "Bootleguard"}
	fmt.Println("Veg List :- ", vegList)


	// Practice
	var nameList[3] string
	nameList[0] = "Roshan"
	nameList[1] = "Kumar"
	nameList[2] = "Gupta"
	fmt.Println("Name List - ", nameList)

	var fullName = [3]string{"Roshan", "Kumar", "Gupta"}
	fmt.Println("Full Name - ", fullName)
}