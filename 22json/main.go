package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string            `json:"course_name"`
	Price    string               `json:"price"`
	Platform string            `json:"platform"`
	Password string            `json:"-"`
	Tags      []string          `json:"tags,omitempty"`
	Category map[string]string `json:"catoegory"`
	Meta     [3]string         `json:"meta"`
}

func main() {
	fmt.Println("JSON Creation in Go...")
	// EncodingJson()
	DecodeJson()
}

func EncodingJson() {
	webCourse := []course{
		{
			"GoLang",
			"2999",
			"Youtube",
			"ramalloc",
			[]string{"web-dev", "GoLang", "Backend", "Server"},
			map[string]string{
				"GO": "GoLang",
				"JS": "JavaScript",
			},
			[3]string{"Shallow", "Dark Web", "Internet"},
		},
		{
			"Javascript",
			"4999",
			"Youtube",
			"ramalloc",
			[]string{"web-dev", "GoLang", "Backend", "Server"},
			map[string]string{
				"GO": "GoLang",
				"JS": "JavaScript",
				"ND": "Node",
			},
			[3]string{"Shallow", "Dark Web", "Internet"},
		},
		{
			"RubyOnRails",
			"9999",
			"Youtube",
			"ramalloc",
			nil,
			map[string]string{
				"GO": "GoLang",
				"JS": "JavaScript",
				"RB": "Ruby",
			},
			[3]string{"", "", ""}, // Replace nil with valid array
		},
	}

	// Packaging of Data as JSON Data using json.Marshaler, Marshaler is the implementation of json.
	finalJson, jsonConvErr := json.Marshal(webCourse)
	checkNillErr(jsonConvErr, "Error in JSON Conversion !")
	fmt.Printf("JSON of Web Courses :- %s\n", finalJson)

	// In Marshal there are some issues like - reading is not easy, some key-value issues (nil -> null)
	// So therefore we use MarshalIndent

	// json.MarshalIndent(v {}interface, prefix, indent) -> prefix before each line of data, indent means indentatin type
	//  like space, tab or new line etc.
	// indentedFinalJson, indenJsonErr := json.MarshalIndent(webCourse, "-->", "\t")
	// 	-->             "Meta": [
	// -->                     "Shallow",
	// -->                     "Dark Web",
	// -->                     "Internet"
	// -->             ]
	indentedFinalJson, indenJsonErr := json.MarshalIndent(webCourse, "", "\t")
	checkNillErr(indenJsonErr, "Error in JSON Conversion through MarshalIndent !")
	fmt.Printf("JSON of Web Courses through Marshal :- %s\n", indentedFinalJson)

	// We convert the struct instance into json, but the keys are in upper case and password is not hidden or encrypted
	// So we define `json: small_case_variiable_name` correspondent the key

	// type course struct {
	// 	Name     string `json:"course_name"`
	// 	Price    int `json:"price"`
	// 	Platform string `json:"platform"`
	// 	Password string `json:"-"`
	// 	Tag      []string `json:tag`
	// 	Category map[string]string `json:"catoegory"`
	// 	Meta     [3]string `json:"meta"`
	// }

	// Password string `json : "-"` -> "-" This will not show the password in json response
	// Tag []string `json : tag,omitempty` -> "omitempty" means if there is null at that value then do not show

}


func DecodeJson (){
	// Array of json datas
	jsonWebFrame := []byte(`
		[
			{
				"course_name":"GoLang",
				"price":2999,
				"platform":"Youtube",
				"password":"ramalloc",
				"tags":[
						"web-dev",
						"GoLang",
						"Backend",
						"Server"
				],
				"catoegory":{
						"GO": "GoLang",
						"JS": "JavaScript"
				},
				"meta":[
						"Shallow",
						"Dark Web",
						"Internet"
				]
			}
		]
	`)

	// Array Instance of struct
	var backCourse []course

	// Before performing any operation we need to verify/check that the data is in correct format for that we have json.Valid()
	checkValidJson := json.Valid(jsonWebFrame)
	if(checkValidJson){
		fmt.Println("JSON Is Valid")
		// To consume json we have unmarshal which takes data and struct to verify/validate json
		// We pass address or pointer of structure here to verify
		json.Unmarshal(jsonWebFrame, &backCourse)
		// to print interfaces we use #
		fmt.Printf("%#v\n", backCourse)
	} else{
		fmt.Println("Wrong JSON Format !")
	}


	// There are some cases where we want to pass data as key : value pair
	// As we do not know in which format we are getting values so we are using interface {}
	var myOnlineData []map[string]interface{}
	json.Unmarshal(jsonWebFrame, &myOnlineData)
	fmt.Printf("myOnlineData \n- %#v\n", myOnlineData)

	for i, obj := range myOnlineData {
		fmt.Printf("Index - %d\n", i) // Index of the slice element
		for k, v := range obj {
			fmt.Printf("Key - %v, Value - %v and type is : %T\n", k, v, v)
		}
		fmt.Println()
	}
	
}

func checkNillErr(err error, errFrom string) {
	if err != nil {
		println(errFrom)
		panic(err)
	}
}
