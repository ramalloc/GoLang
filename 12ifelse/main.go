package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Control Flow - If/Else")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your age : ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input : ", err)
	} else{
		age, err2 := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
		if err2 != nil {
			fmt.Println("Error While Taking input :- ", err2)
		} else {
			if age < 16 {
				fmt.Printf("Your are not allowed in the party as your age : %d under 18 !\n", age)
			} else if age == 18 {
				fmt.Printf("You are not allowed to drink as your age is %d\n", age)
			} else if drinkWithEighteen := age; drinkWithEighteen > 24 {
				fmt.Printf("You can drink good stuffs as you are %d \n", drinkWithEighteen)
			} else {
				fmt.Println("Have Fun buddy...")
			}
		}

	}

}
