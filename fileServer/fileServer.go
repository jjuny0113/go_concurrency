package fileServer

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"aaa", 16, 87}
	data, _ := json.Marshal(student)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(data))
}

func makeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student", studentHandler)
	return mux
}
func Init() {
	http.ListenAndServe(":8080", makeWebHandler())
}
