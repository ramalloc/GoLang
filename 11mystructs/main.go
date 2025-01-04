package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang as there is not class...")
	// -> There is not inheritance in goLang, super or parent/child


	// Utitlizing Struct
	roshan := User{"Roshan", "roshangupta1887@gmail.com", true, 23}
	ram := User{"Ram", "ram1911@gmail.com", false, 27}
	fmt.Println("User - ", roshan)
	fmt.Println("User - ", ram)
	// User -  {Roshan roshangupta1887@gmail.com true 23}
	fmt.Printf("roshan user details are :- %v \n", roshan)
	// roshan user details are :- {Roshan roshangupta1887@gmail.com true 23} 


	fmt.Printf("roshan user details are :- %+v \n", roshan)
	fmt.Printf("ram user details are :- %+v \n", ram)
	// roshan user details are :- {Name:Roshan Email:roshangupta1887@gmail.com Status:true Age:23} 

	fmt.Printf("Name is %v and Email is %v\n", roshan.Name, roshan.Email)
	fmt.Printf("Name is %v and Email is %v\n", ram.Name, ram.Email)
	// Name is Roshan and Email is roshangupta1887@gmail.com
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
