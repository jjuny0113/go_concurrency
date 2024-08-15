package restfulAPI

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sort"
	"strconv"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

var students map[int]Student
var lastId int

func getStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0)

	for _, student := range students {
		list = append(list, student)
	}

	sort.Sort(list)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}
func getStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func postStudentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post")
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastId++
	student.Id = lastId
	students[lastId] = student
	w.WriteHeader(http.StatusCreated)
}

func deleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	_, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	delete(students, id)
	w.WriteHeader(http.StatusOK)
}

func makeWebHandler() http.Handler {
	fmt.Println("init!!")
	mux := mux.NewRouter()
	mux.HandleFunc("/students", getStudentListHandler).Methods(http.MethodGet)
	mux.HandleFunc("/students/{id:[0-9]+}", getStudentHandler).Methods(http.MethodGet)
	mux.HandleFunc("/students", postStudentHandler).Methods(http.MethodPost)
	mux.HandleFunc("/students/{id:[0-9]+}", deleteStudentHandler).Methods(http.MethodDelete)

	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2
	return mux
}

func Init() {
	http.ListenAndServe(":8080", makeWebHandler())
}
