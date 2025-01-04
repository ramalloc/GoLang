package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const myUrl string = "http://localhost:3000/"

func main() {
	fmt.Println("Welcome Web Request Verbs...")
	ConnectToServer()
	PerformGetRequest()
	PerformPostRequest()
	PerformPostRequestUsingString()
	PostFormData()
}

func ConnectToServer() {
	res, err := http.Get(myUrl)
	checkNillErr(err)
	defer res.Body.Close()

	fmt.Println("Status Code - ", res.StatusCode)
	fmt.Println("Content Length Code - ", res.ContentLength)

	dataBytes, byteErr := io.ReadAll(res.Body)
	checkNillErr(byteErr)

	fmt.Println("Get Mehtod Body - ", string(dataBytes))

}

func PerformGetRequest() {
	getFullUrl := &url.URL{
		Scheme: "http",
		Host:   "localhost:3000",
		Path:   "get",
	}

	res, err := http.Get(getFullUrl.String())
	checkNillErr(err)
	defer res.Body.Close()

	fmt.Println("Status Code - ", res.StatusCode)
	fmt.Println("Content Length Code - ", res.ContentLength)

	dataBytes, byteErr := io.ReadAll(res.Body)
	checkNillErr(byteErr)

	fmt.Println("Get Mehtod Body - ", string(dataBytes))

}

func PerformPostRequest() {
	var builder strings.Builder
	builder.WriteString("http://localhost:3000/")
	builder.WriteString("post")

	postUrl := builder.String()

	data := make(map[string]interface{})
	// We use interface{} is a special type known as the empty interface, and it is used to represent any type of value

	var name string
	fmt.Print("Enter name: ")
	fmt.Scan(&name)
	data["name"] = name

	var age int
	fmt.Print("Enter age: ")
	fmt.Scan(&age)
	data["age"] = age

	// Convert the data map to JSON
	postData, marshalErr := json.Marshal(data)
	checkNillErr(marshalErr)
	fmt.Println("Post Data - ", postData)

	res, err := http.Post(postUrl, "application/json", bytes.NewBuffer(postData))
	checkNillErr(err)

	defer res.Body.Close()

	fmt.Println("Status Code - ", res.StatusCode)
	fmt.Println("Content Length Code - ", res.ContentLength)

	dataBytes, byteErr := io.ReadAll(res.Body)
	checkNillErr(byteErr)

	fmt.Println("Post Mehtod Body Reponse is - ", string(dataBytes))

}
func PerformPostRequestUsingString() {
	var builder strings.Builder
	builder.WriteString(myUrl)
	builder.WriteString("post")

	postUrl := builder.String()

	requestBody := strings.NewReader(`
		{
			"couser_name" : "GoLang",
			"price" : "399",
			"author" : "Roshan Kumar"
		}
	`)

	res, err := http.Post(postUrl, "application/json", requestBody)
	checkNillErr(err)

	defer res.Body.Close()

	fmt.Println("Status Code - ", res.StatusCode)
	fmt.Println("Content Length Code - ", res.ContentLength)

	dataBytes, byteErr := io.ReadAll(res.Body)
	checkNillErr(byteErr)

	if res.StatusCode != 200 {
		fmt.Println("Error - ", string(dataBytes))
	} else {
		fmt.Println("Post Mehtod Body Reponse Using String Dummy JSON Data is - ", string(dataBytes))
	}

}

func PostFormData () {
	var builder strings.Builder
	builder.WriteString(myUrl)
	builder.WriteString("postform")

	postFormUrl := builder.String()
	
	// URL Values or RawQuery
	data := url.Values{}
	data.Add("first_name", "Roshan")
	data.Add("last_name", "Kumar")
	data.Add("age", "23")
	data.Add("interest", "GO")

	res, err := http.PostForm(postFormUrl, data)
	checkNillErr(err)

	defer res.Body.Close()

	fmt.Println("Status Code - ", res.StatusCode)
	fmt.Println("Content Length Code - ", res.ContentLength)

	dataBytes, _ := io.ReadAll(res.Body)
	var dataBuilder strings.Builder
	byteCount, _ := dataBuilder.Write(dataBytes)
	fmt.Println("byteCount from builder reponse - ", byteCount)
	fmt.Println("Response from Post Form Method :- ", dataBuilder.String())

}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
