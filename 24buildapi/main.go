package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// CRUD Operation for a course

type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"-"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"full_name"`
	Website  string `json:"website"`
}

// DB
var courses []Course

// middlewares/helper
func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - GoLang Welcomes You")

	// Seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "React.js", CoursePrice: 299, 
	Author: &Author{Fullname: "Roshan Kumar", Website: "www.google.com"}})

	courses = append(courses, Course{CourseId: "2", CourseName: "GoLang", CoursePrice: 299, 
	Author: &Author{Fullname: "Raghav Kumar", Website: "www.instagram.com"}})

	r := mux.NewRouter()

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Listening on Port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers

// --- Routes ---
// Serve Home Route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Serving/Home Page...</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses...")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course...")
	w.Header().Set("Content-Type", "application/json")

	// Getting id from request using mux variables
	params := mux.Vars(r)

	// Loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found with respect to id !")

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course...")
	w.Header().Set("Content-Type", "application/json")

	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data !")
		return
	}

	defer r.Body.Close()

	// what if there is empty json - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data in json !")
		return
	}

	for _, existingCourse := range courses {
		if existingCourse.CourseName == course.CourseName {
			json.NewEncoder(w).Encode(map[string]string{
				"error":       "Course with the same name already exists",
				"course_name": course.CourseName,
			})
			return
		}
	}

	// Data Creation
	// Generate unique id and convert into string
	// append course in Courses
	source := rand.NewSource(time.Now().UnixNano())
	localRand := rand.New(source)
	course.CourseId = strconv.Itoa(localRand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}


func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Course...")
	w.Header().Set("Content-Type", "application/json")

	// Check if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Please send some data!"})
		return
	}

	defer r.Body.Close()

	// Decode the JSON into a Course struct
	var updatedCourse Course
	err := json.NewDecoder(r.Body).Decode(&updatedCourse)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON data!"})
		return
	}

	// Check if the decoded course is empty
	if updatedCourse.isEmpty() {
		json.NewEncoder(w).Encode(map[string]string{"error": "No data in JSON!"})
		return
	}

	// Get ID from the request parameters
	params := mux.Vars(r)
	courseID := params["id"]

	// Loop through the courses to find and update
	for index, existingCourse := range courses {
		if existingCourse.CourseId == courseID {
			// Replace the course in the slice
			courses[index] = updatedCourse
			courses[index].CourseId = courseID // Ensure the ID remains unchanged

			response := map[string]interface{}{
				"message": "Course updated successfully",
				"course":  courses[index],
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// If no course is found with the given ID
	json.NewEncoder(w).Encode(map[string]string{"error": "No course found with the given ID!"})
}



func deleteOneCourse (w http.ResponseWriter, r *http.Request){
	fmt.Println("Deleting One Course...")	
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil{
		json.NewEncoder(w).Encode("please send some data !")
		return
	}

	defer r.Body.Close()

	var course Course
	_ =json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data in JSON !")
		return
	}

	params := mux.Vars(r)
	courseID := params["id"]

	for index, course := range courses {
		if course.CourseId == courseID {
			deletedCourse := course
			courses = append(courses[:index], courses[index+1:]...)
			response := map[string]interface {}{
				"message" : "Course Deleted Successfully",
				"course" : deletedCourse,
			}

			json.NewEncoder(w).Encode(response)
			return
		}
	}
}
