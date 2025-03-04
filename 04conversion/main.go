package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Conversion...")
	fmt.Println("Please Rate our service between 1 to 5..")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks For rating :- ", input)

	// --> Conversion Of Types
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error While Entering Rating :- ", err)
	} else {
		if numRating > 5 || numRating < 1 {
			fmt.Println("Enter Rating Under 1 to 5 !")
		} else {
			fmt.Println("Rating With Addition 1 :- ", numRating+1)
		}
	}

	// Practice
	accept := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Count :- ")
	num, _ := accept.ReadString('\n')

	fmt.Printf("Type of num :- %T\n", num)

	newNum, _ := strconv.ParseFloat(strings.TrimSpace(num), 64)
	fmt.Printf("Converted to Number thorugh parseFloat :- %f\n", newNum)

	newNumInt, _ := strconv.ParseInt(strings.TrimSpace(num), 10, 64)
	fmt.Printf("Converted to Number thorugh ParseInt :- %d\n", newNumInt)

	newNum1, _ := strconv.Atoi(strings.TrimSpace(num))
	fmt.Printf("Converted to Number thorugh Atoi :- %d\n", newNum1)

	// THe below methods can only inter convert into int/float will not work with strings.
	newNum32 := float32(newNum)
	fmt.Printf("Converted to Number through float32 :- %f\n", newNum32)

	newNumInt64 := int64(newNum)
	fmt.Printf("Converted to Number through int64 :- %d\n", newNumInt64)

}
