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

type Course struct {
	CourseId    string  `json:"course_id"`
	CoursName   string  `json:"course_name"`
	CoursePrise int     `json:"course_price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"full_name"`
	Website  string `json:"website"`
}

// DB
var courses []Course

func (c *Course) iSEmpty() bool {
	return c.CoursName == ""
}

func main() {
	fmt.Println("Ash-Api")
	r := mux.NewRouter()
	//seeding
	courses = append(courses, Course{CourseId: "1", CoursName: "Go", CoursePrise: 100, Author: &Author{FullName: "Ash", Website: "ash.com"}})
	courses = append(courses, Course{CourseId: "2", CoursName: "Python", CoursePrise: 200, Author: &Author{FullName: "Ash", Website: "ash.com"}})
	courses = append(courses, Course{CourseId: "3", CoursName: "Java", CoursePrise: 300, Author: &Author{FullName: "Ash", Website: "ash.com"}})

	//routing
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", geytAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleatOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Welcome to the 24API running on localhost"}`))
}

func geytAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")
	// grrab Id from the request
	params := mux.Vars(r)
	fmt.Println("Params", params)

	// looping through courses and find the course with id and return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	// waht if: bofy is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.iSEmpty() {
		json.NewEncoder(w).Encode("No data in the request body")
		return
	}

	//genarate unique id
	//append the course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Course")
	w.Header().Set("Content-Type", "application/json")

	//garb id from req
	param := mux.Vars(r)

	//loop, id, remove, add

	for index, course := range courses {
		if course.CourseId == param["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// to-do send a response when id not found
	json.NewEncoder(w).Encode("No course found with given id")
}

func deleatOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from req
	param := mux.Vars(r)

	//loop, id, remove
	for index, course := range courses {
		if course.CourseId == param["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			break
		}
	}
	// to-do send a response when id not found
	json.NewEncoder(w).Encode("No course found with given id")
	return
}
