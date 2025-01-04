package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Just like other languages we can handle only text files in GoLang and for other file type we have to use libraries.
	fmt.Println("We are handling files here...")
	content := "This will be the content for text file."

	// File creation
	file, error := os.Create("./myTextFile.txt")

	if error != nil {
		panic(error)
		// panic stops the execution of program and show the error which occurs.
	}

	defer file.Close()

	// File Writing
	// on using io's write string method this will return length of the content in the file.
	legth, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("File created successfully with length ....", legth)
	}

	readfile("./myTextFile.txt")

}

func readfile(fileName string) {
	// To read files we have os.ReadFile
	// When we read a file we do not read the file in string format, it reads file/data in bytes format. So ReadFile return bytes.
	dataBytes, err := os.ReadFile(fileName)

	// if err != nil {
	// 	panic(err)
	// }

	checkNillErr(err)

	fmt.Println("The my text file is  :- ")
	// To convert these bytes into string/text we can do like below using string mathod
	fmt.Println(string(dataBytes))

}

// Common Error Check
func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
