package main

import "fmt"

func main() {
	fmt.Println("We will learn Methods...")

	// As we make functions in classes that's called methods as there is no classes, we have structs.

	userDetail := User{"Roshan", "roshangupta1887@gmail.com", true, 23}
	userDetail.GetStatus()
	// The User is Active...

	fmt.Println("User Old Name", userDetail.Name)
	// User Old Name Roshan
	userDetail.Name = "Ram"
	fmt.Println("User Updated Name", userDetail.Name)
	// User Updated Name Ram

	userDetail.NewMail("ramalloc@123")
	// New mail :-  ramalloc@123

	fmt.Println("Original Email - ", userDetail.Email)
	// Original Email -  roshangupta1887@gmail.com

	// The original Values did not affect because only copies of objects are passing and an instance is creating in funciton.


}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}

// THis is the way of declaring function/method of struct we cannot define it in main we have to define it explicitly
func (u User) GetStatus(){
	if u.Status {
		fmt.Println("The User is Active...")
	} else {
		fmt.Println("The user is not active!")
	}
}

func (u User) NewMail(email string) {
	u.Email = email
	fmt.Println("New mail :- ", u.Email)
}
