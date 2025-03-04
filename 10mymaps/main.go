package main

import "fmt"

func main() {
	fmt.Println("Maps in Go Lang...")
	// --> We can use make to create maps as well, It is key value pair data structure

	// ---> MAPS
	languages := make(map[string]string)
	languages["GO"] = "GoLang"
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println("Languages - ", languages)
	// Languages -  map[GO:GoLang JS:Javascript PY:Python RB:Ruby]

	fmt.Println("JS stands for : ", languages["JS"])
	// JS stands for :  Javascript

	// Delete in MAPS, we can use this delete in slices as well
	delete(languages, "PY")
	fmt.Println("Languages after PY delet - ", languages)
	// Languages after PY delet -  map[GO:GoLang JS:Javascript RB:Ruby]


	// --> LOOPS in MAPS
	for key, value := range languages{
		fmt.Printf("For key %v the value is : %v\n", key, value)
	}
	// For key RB the value is : Ruby
	// For key GO the value is : GoLang
	// For key JS the value is : Javascript


	// Practice
	prizes := make(map[int]string)

	prizes[0] = "Chal Nikal"
	prizes[1] = "First"
	prizes[2] = "Second"
	prizes[3] = "Third"

	fmt.Println("Printing Map - ", prizes)
	fmt.Println("Looping Map - ")
	for key, value := range prizes{
		fmt.Printf("For %v number the prize is :- %v\n", key, value)
	}

}
