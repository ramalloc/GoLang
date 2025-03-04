package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
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
	checkNillErr(err, "Get Method Error while connecting !")
	defer res.Body.Close()

	fmt.Println("Status Code - ", res.StatusCode)
	fmt.Println("Content Length Code - ", res.ContentLength)

	dataBytes, byteErr := io.ReadAll(res.Body)
	checkNillErr(byteErr, "Reading Byte Error in connection !")

	fmt.Println("Get Mehtod Body - ", string(dataBytes))

}

func PerformGetRequest() {
	getFullUrl := &url.URL{
		Scheme: "http",
		Host:   "localhost:3000",
		Path:   "get",
	}

	res, err := http.Get(getFullUrl.String())
	checkNillErr(err, "Get Method Error !")
	defer res.Body.Close()

	fmt.Println("Status Code - ", res.StatusCode)
	fmt.Println("Content Length Code - ", res.ContentLength)

	dataBytes, byteErr := io.ReadAll(res.Body)
	checkNillErr(byteErr, "Reading Byte error in Get !")

	fmt.Println("Get Mehtod Body - ", string(dataBytes))

}

func PerformPostRequest() {
	var builder strings.Builder
	builder.WriteString("http://localhost:3000/")
	builder.WriteString("post")

	postUrl := builder.String()

	data := make(map[string]interface{})
	// We use interface{} is a special type known as the empty interface, and it is used to represent any type of value

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	data["name"] = strings.TrimSpace(name)

	fmt.Print("Enter age: ")
	age, _ := reader.ReadString('\n')
	ageInt, err := strconv.Atoi(strings.TrimSpace(age))
	checkNillErr(err, "string conversion error to int !")
	data["age"] = ageInt

	// Convert the data map to JSON
	postData, marshalErr := json.Marshal(data)
	checkNillErr(marshalErr, "Json consversion error !")
	fmt.Println("Post Data - ", postData)

	// Converting Data bytes of json into json string
	fmt.Println("Post Data conversion thorugh byte.NewBuffer byte Buffer - ", bytes.NewBuffer(postData))
	fmt.Println("Post Data thorugh string - ", string(postData))

	res, err := http.Post(postUrl, "application/json", bytes.NewBuffer(postData))
	checkNillErr(err, "Post Method Error")

	defer res.Body.Close()

	fmt.Println("Status Code - ", res.StatusCode)
	fmt.Println("Content Length Code - ", res.ContentLength)

	dataBytes, byteErr := io.ReadAll(res.Body)
	checkNillErr(byteErr, "Error in Post Databyte !")

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
	checkNillErr(err, "Error in Post Method!")

	defer res.Body.Close()

	fmt.Println("Status Code - ", res.StatusCode)
	fmt.Println("Content Length Code - ", res.ContentLength)

	dataBytes, byteErr := io.ReadAll(res.Body)
	checkNillErr(byteErr, "Error in Reading Byte")

	if res.StatusCode != 200 {
		fmt.Println("Error - ", string(dataBytes))
	} else {
		fmt.Println("Post Mehtod Body Reponse Using String Dummy JSON Data is - ", string(dataBytes))
	}

}

func PostFormData() {
	var builder strings.Builder
	builder.WriteString(myUrl)
	builder.WriteString("postform")

	postFormUrl := builder.String()

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter First Name :- ")
	firstName, _ := reader.ReadString('\n')
	fmt.Printf("Enter Last Name :- ")
	lastName, _ := reader.ReadString('\n')
	fmt.Printf("Enter Age Name :- ")
	age, _ := reader.ReadString('\n')
	fmt.Printf("Enter Interest Name :- ")
	interest, _ := reader.ReadString('\n')



	// URL Values or RawQuery
	data := url.Values{}
	data.Add("first_name", strings.TrimSpace(firstName))
	data.Add("last_name", strings.TrimSpace(lastName))
	data.Add("age", strings.TrimSpace(age))
	data.Add("interest", strings.TrimSpace(interest))

	res, err := http.PostForm(postFormUrl, data)
	checkNillErr(err, "Error while post Form")

	defer res.Body.Close()

	fmt.Println("Status Code - ", res.StatusCode)
	fmt.Println("Content Length Code - ", res.ContentLength)

	dataBytes, _ := io.ReadAll(res.Body)
	var dataBuilder strings.Builder
	byteCount, _ := dataBuilder.Write(dataBytes)
	fmt.Println("byteCount from builder reponse - ", byteCount)
	fmt.Println("Response from Post Form Method :- ", dataBuilder.String())

}

func checkNillErr(err error, errFrom string) {
	if err != nil {
		println(errFrom)
		panic(err)
	}
}
