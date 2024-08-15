package todoList

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"sort"
	"strconv"
)

var rd *render.Render

type Todo struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Completed bool   `json:"completed,omitempty"`
}

type Success struct {
	Success bool `json:"success"`
}

type Todos []Todo

func (t Todos) Len() int {
	return len(t)
}

func (t Todos) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Todos) Less(i, j int) bool {
	return t[i].ID > t[j].ID
}

var todoMap map[int]Todo
var lastID int = 0

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Todos, 0)
	for _, todo := range todoMap {
		list = append(list, todo)
	}
	sort.Sort(list)
	rd.JSON(w, http.StatusOK, list)
}
func postTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastID++
	todo.ID = lastID
	todoMap[lastID] = todo
	rd.JSON(w, http.StatusCreated, todo)
}
func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(w, http.StatusOK, Success{true})
		return
	}
	rd.JSON(w, http.StatusNotFound, Success{false})
}
func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if todo, ok := todoMap[id]; ok {
		todo.Name = newTodo.Name
		todo.Completed = newTodo.Completed
		rd.JSON(w, http.StatusOK, Success{true})
		return
	}
	rd.JSON(w, http.StatusBadRequest, Success{false})

}
func makeWebHandler() http.Handler {
	todoMap = make(map[int]Todo)
	mux := mux.NewRouter()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/todos", getTodoListHandler).Methods(http.MethodGet)
	mux.HandleFunc("/todos", postTodoHandler).Methods(http.MethodPost)
	mux.HandleFunc("/todos/{id:[0-9]+}", removeTodoHandler).Methods(http.MethodDelete)
	mux.HandleFunc("/todos/{id:[0-9]+}", updateTodoHandler).Methods(http.MethodPut)
	return mux
}

func Init() {
	rd = render.New()
	m := makeWebHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", n)
	if err != nil {
		panic(err)
	}
}
