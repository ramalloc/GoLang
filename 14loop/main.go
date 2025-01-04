package main

import "fmt"

func main() {
	fmt.Println("Loops in GoLang")
	// days := make([]string, 0)
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	// days = append(days, "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday")

	fmt.Println("Days in Week : ", days)

	// For Loop
	for i:=0; i < len(days); i++ {
		fmt.Printf("Day at poisition %d is %s\n",i , days[i])
	}


	// For Range Loop
	for i := range days{
		// It gives index in as i
		fmt.Printf("Day at poisition %d is %s\n",i , days[i])
	}

	// For Key Value Loop
	for index, day := range days{
		// day will return data at that index
		fmt.Printf("Day at poisition %v is %v\n",index , day)
	}


	// While Loop here used as for

	rougueValue := 1

	for rougueValue <= 10 {
		if rougueValue == 2 {
			rougueValue++
			continue
		}
		if rougueValue == 8 {
			goto ramalloc
		}
		if rougueValue == 9 {
			rougueValue++
			break
		}
		fmt.Println("value is : ", rougueValue)
		rougueValue++
	}

	// LABEL 
	ramalloc: 
		fmt.Println("Lable is a pieces of code here we are using for goto...")
}
