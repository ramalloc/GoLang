package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Welcome to Conversion...")
	fmt.Println("Please Rate our service between 1 to 5..")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks For rating :- ", input)

	// --> Conversion Of Types
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil{
		fmt.Println("Error While Entering Rating :- ", err)
	}else{
		if numRating > 5 || numRating < 1{
			fmt.Println("Enter Rating Under 1 to 5 !")
		}else{	
			fmt.Println("Rating With Addition 1 :- ", numRating + 1)
		}
	}

}