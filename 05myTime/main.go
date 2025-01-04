package main

import (
	"fmt"
	"time"
)

func main() {
	// --->
	fmt.Println("We will learn about time in GoLang...")

	presentTime := time.Now()
	fmt.Println("Present Full Time :- ", presentTime)

	fmt.Println("Present Date :- ", presentTime.Format("01-02-2006"))
	fmt.Println("Present Date and Day :- ", presentTime.Format("01-02-2006 Monday"))
	fmt.Println("Present Date, Time and Day :- ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	// ---> Creation of Time 
	createdDate := time.Date(2020, time.August, 15, 23, 45, 36, 001, time.UTC)
	fmt.Println("Created Date :- ", createdDate)
	fmt.Println("Formated Created Date :- ", createdDate.Format("01-02-2006 Monday"))
}
