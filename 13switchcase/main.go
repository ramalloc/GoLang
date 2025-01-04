package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in GoLang...")
	source := rand.NewSource(time.Now().UnixNano())
	localRand := rand.New(source)
	diceNumber := localRand.Intn(7) + 1
	fmt.Printf("Your dice number is : %d\n", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1, Now you can open")
	case 2:
		fmt.Println("Dice value is 2, Now you can move to 2")
	case 3:
		fmt.Println("Dice value is 3, Now you can move to 3")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4, Now you can move to 4")
		fallthrough
	case 5:
		fmt.Println("Dice value is 5, Now you can move to 5")
	case 6:
		fmt.Println("Dice value is 6, You can roll the dice once again...")
	default: 
		fmt.Println("What was that out of the box!!!")
	}

	// fallthrough causes execution of next Case without rechecking. It implement next case without checking case.
}
