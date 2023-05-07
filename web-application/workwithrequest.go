package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Course struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Instructor string  `json:"instructor"`
}

var courseList []Course

func init() {
	courseJSON := `[
		{
			"id" : 1,
			"name" : "Python",
			"price" : 2000,
			"instructor" : "AIT"
		},
		{
			"id" : 2,
			"name" : "C",
			"price" : 1000,
			"instructor" : "AIT"
		},
		{
			"id" : 3,
			"name" : "GoLang",
			"price" : 5000,
			"instructor" : "AIT"
		}
	]`
	err := json.Unmarshal([]byte(courseJSON), &courseList)
	if err != nil {
		log.Fatal(err)
	}
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	courseJSON, err := json.Marshal(courseList)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJSON)
	case http.MethodPost:
		var newCourse Course
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		courseList = append(courseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func main() {
	http.HandleFunc("/course", courseHandler)
	http.ListenAndServe(":5000", nil)
}
