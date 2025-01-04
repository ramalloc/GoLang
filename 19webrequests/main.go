package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://api.github.com/users/ramalloc"

func main() {
	fmt.Println("Web Requests in GoLang...")

	response, error := http.Get(url)
	checkNillErr(error)

	fmt.Printf("Type of Response :- %T\n", response)
	// Type of Response :- *http.Response
	// --> Above we are not getting copy of response we are getting actual response as pointer or reference
	defer response.Body.Close()

	// Reading response
	dataBytes, byteError := io.ReadAll(response.Body)
	checkNillErr(byteError)
	fmt.Println("Response data in string :- ", string(dataBytes))

}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
