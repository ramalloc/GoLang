package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// 1. The mux will show in go.mod with require and indirect comment as that package is not in use if it will in use then
	// indirect will be removed.
	// 2. When we download any package, it will create a file called go.sum, Which has verify any changes in the github hash
	// with respect ot that repo.
	// All the library which we download that will go in cache file of go mod not in working directory and whenever we make
	// new request for that package go fetch that package from local copy stored in cache not from internet.
	// go mod why github.com/gorilla/mux -> This show the dependencies on that package or where it is being used.
	// go mod graph - This command will show the dependency graph of modules.
	// go mod edit -go 1.16, go mod edit -module 1.16 -> It is used to edit the go.mod file through terminal.
	// go mod vendor -> As go need any file it fetch from cache but when we run vendor command then all things came here.
	// 					Now it will fetch all things from vendor
	// go run -mod=vendor main.go -> Now the programme will get the dependencies from vendor not from cache or internet

	Greeter()

	// Route and Handler in Go using MUX
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	// Serving the api
	// We Log package to handle fatal/error if any error comes in it
	log.Fatal(http.ListenAndServe(":4000", r))

	// The indirect is still showing in go.mod as we use "go mod tidy" command in cmd the package will be used as direct.
	// and tidy will remove all unused pacakges and bundle everything.
	// We use "go mod veirfy" to verify the modules in mod thorugh go.sum and take those hashes and will check.
	//  "go list" -> This will show the list of files of directory in go that have been used.
	//  "go list all" -> Show all file list used and un-used both 
	//  "go list -m all" -> will show only used files
	//  "go list -m -versions github/gorilla/mux" -> will show versions of mux
}

func Greeter() {
	fmt.Println("Good Morning...")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	// w http.ResponseWriter is used write response
	// r *http.Request is used get request details
	w.Write([]byte("<h1>Welocme to goL	ang Backend Route...</h1>"))
}
