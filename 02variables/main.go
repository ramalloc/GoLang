package main

import "fmt"

// --> We can declare constants for variable which we do not want to change
// --> If We Initialize/declare the Variables with Capital means first capital letter then the variable will be public.
const LoginToken = "kjhe78seiu98"

func data()  {
	var name string = "Raj Kundrua"
	fmt.Println(name)
	fmt.Println("Login Token - ", LoginToken)
}

func main()  {
	var username string = "Roshan Kumar"
	fmt.Println("Username - ", username)
	fmt.Printf("Type Of Username :  %T \n", username)
	// Username -  Roshan Kumar
	// Type Of Username :  string 

	var isUser bool = true
	fmt.Println("isUser - ", isUser)
	fmt.Printf("Type Of isUser : %T \n", isUser)
	// 	isUser -  true
	// Type Of isUser : bool 

	var smallInt uint8 = 255
	fmt.Println("smallInt - ", smallInt)
	fmt.Printf("Type Of smallInt : %T \n", smallInt)
	// 	smallInt -  255
	// Type Of smallInt : uint8 

	var smallFloat float32 = 255.6586688796
	fmt.Println("smallFloat - ", smallFloat)
	fmt.Printf("Type Of smallFloat : %T \n", smallFloat)
	// 	smallFloat -  255.65868
	// Type Of smallFloat : float32 

	var smallFloatExtended float64 = 255.6586688796657
	fmt.Println("smallFloatExtended - ", smallFloatExtended)
	fmt.Printf("Type Of smallFloatExtended : %T \n", smallFloatExtended)
	// 	smallFloatExtended -  255.6586688796657
	// Type Of smallFloatExtended : float64 

	// --- >>  Default Values in variable at the time of declaration
	var defaultStringVar string
	var defaultBoolVar bool
	var defaultIntVar int
	var defaultFloatVar float32
	fmt.Println("defaultStringVar - ", defaultStringVar)
	fmt.Printf("Type Of defaultStringVar : %T \n", defaultStringVar)
	fmt.Println("defaultBoolVar - ", defaultBoolVar)
	fmt.Printf("Type Of defaultBoolVar : %T \n", defaultBoolVar)
	fmt.Println("defaultIntVar - ", defaultIntVar)
	fmt.Printf("Type Of defaultIntVar : %T \n", defaultIntVar)
	fmt.Println("defaultFloatVar - ", defaultFloatVar)
	fmt.Printf("Type Of defaultFloatVar : %T \n", defaultFloatVar)

	// ---> Implicit Type/Way of Declaring variable
	var implicitData = "Roshan Kumar"
	fmt.Println("implicitData - ", implicitData)

	// ---> No Var Style variable Declaration
	// -> We are allowed to use/declare variables like this only in methods not in global context/scope or outside of method.
	num := 200.001
	fmt.Println("Without Var Int - ", num)

	data()
}
