package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://api.github.com:3000/users/ramalloc?id=95184750&type=User"
func main() {
	// There are oftem time when we need to get some info from Urls like params, body etc. We have different libraries for that
	// in easier way But We can handle the urls from go also
	fmt.Println("Welcome to URLs")

	// PARSING
	result, error := url.Parse(myUrl)
	checkNillErr(error)

	fmt.Println("Result Scheme:- ", result.Scheme)
	// Result :-  https
	fmt.Println("Result Host:- ", result.Host)
	fmt.Println("Result Path:- ", result.Path)
	fmt.Println("Result RawQuery:- ", result.RawQuery)
	fmt.Println("Result Port Method:- ", result.Port())

	qParams := result.Query()
	fmt.Printf("Type of qParams - %T\n", qParams)
	// Type of qParams - url.Values or (Key, Value)
	fmt.Println("qParams - ", qParams)
	// qParams -  map[id:[95184750] type:[User]]

	for key, value := range qParams{
		fmt.Println(key, value)
	}

	// CREATING URL FROM URL METHOD

	partsOfUrls := &url.URL{
		Scheme: "https",
		Host: "api.github.com",
		Path: "/users/ramalloc",
		// RawQuery: "id=95184750&type=User",
	}

	anotherUrl := partsOfUrls.String()
	fmt.Println("anotherUrl - ", anotherUrl)

}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
