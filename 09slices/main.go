package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Lear about slices..")
	// Slices --> Under the hood slices are array with extra features
	// We have to initialize the slice if we are using the nelow syntax.
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	//slice -> it does have dynamic storage
	fmt.Printf("Fruitlist 1 slice type - %T\n", fruitList) // --> []string

	//array -> it doesn't have dynamic storage
	var fruitList1 = [2]string{"Apple", "Peach"}
	fmt.Printf("Fruitlist 1 Array type - %T\n", fruitList1) // --> [2]string

	// Putting Values in Slice we use append
	fruitList = append(fruitList, "Banana", "Mango")
	fmt.Println("Appended Values : - ", fruitList)
	// Appended Values : -  [Apple Tomato Peach Banana Mango]

	// Slicing the slice
	// The slices range are always non-inclusive
	fruitList = append(fruitList[1:]) // [Tomato Peach Banana Mango] -> Apple removed
	fmt.Println("Slice from first index and end is after all - ", fruitList)
	fruitList = append(fruitList[:3]) // [Tomato Peach Banana] -> Apple and Mango removed
	fmt.Println("Slice from null index and end is 3 - ", fruitList)

	fruitList = append(fruitList[1:3]) // [Peach Banana] -> Apple, Tomato and Mango removed
	fmt.Println("Slice from first index and end is 3 - ", fruitList)

	// Slices using make() memory allocation
	highScores := make([]int, 4)
	highScores[0] = 90
	highScores[1] = 60
	highScores[2] = 99
	highScores[3] = 70
	// the below will throw an error as the size of slice allocated 4 and we are putting 5th data but we can do this using append
	// highScores[4] = 80

	highScores = append(highScores, 88, 69, 76)
	fmt.Println("highScores :- ", highScores)
	// highScores :-  [90 60 99 70 88 69 76]

	sort.Ints(highScores)
	fmt.Println("sorted highScore - ", highScores)
	// sorted highScore -  [60 69 70 76 88 90 99]

	isSorted := sort.IntsAreSorted(highScores)
	fmt.Println("Is High Score sorted :- ", isSorted)
	// Is High Score sorted :-  true


	// ---> How to remove value from slices based on index
	course := []string{"reactjs", "javascript", "swift", "tailor", "none"}
	var index int = 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println("Removeed 2nd index value :- ", course)
	// Removeed 2nd index value :-  [reactjs javascript tailor none]
}
