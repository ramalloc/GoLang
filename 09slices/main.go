package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Lear about slices..")
	// Slices --> Under the hood slices are array with extra features
	// We have to initialize the slice if we are using the below syntax.
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
	// In slices the starting index in range is always non-inclusive and index starts with 1 not 0.
	fruitList = append(fruitList[1:]) // [Tomato Peach Banana Mango] -> Apple removed
	fmt.Println("Slice from first index and end is after all - ", fruitList)
	fruitList = append(fruitList[:3]) // [Apple Tomato Peach] -> Banana and Mango removed as it is after 3 index
	fmt.Println("Slice from null index and end is 3 - ", fruitList)

	fruitList = append(fruitList[1:3]) //  [Tomato Peach] -> Apple, Tomato and Mango removed
	fmt.Println("Slice from first index and end is 3 - ", fruitList)

	// Slices using make() memory allocation
	// make takes three arguments (type, length, size) - > The type can be only slice, map or chan.
	highScores := make([]int, 4) // This make allocates an slice of length 4
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
	course = append(course[:index], course[index+1:]...) // append([1, 2], [2+1 = 3 => 4, 5])
	fmt.Println("Removeed 2nd index value :- ", course)
	// Removeed 2nd index value :-  [reactjs javascript tailor none] -> 3rd element removed
}
