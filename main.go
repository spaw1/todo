package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type TodoRequest struct {
	Message string `json:"message"`
}

type TodoResponse struct {
	Success bool   `json:"success"`
	Body    string `json:"body"`
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/todo", GetAllAction)
	router.HandleFunc("/todo/{id}", GetByIdAction)
	router.HandleFunc("/todo/{id}/updateStatus", UpdateAction).Methods("PATCH")
	router.HandleFunc("/create", CreateAction).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func CreateAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body TodoRequest
	json.NewDecoder(r.Body).Decode(&body)
	var todoItem = Create(body.Message)
	json.NewEncoder(w).Encode(todoItem)
}

func GetAllAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todos := GetAll()
	for _, todo := range todos {
		sendResponse(w, todo)
	}
}

func GetByIdAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	idI, _ := strconv.Atoi(id)
	todo := GetTodo(idI)
	sendResponse(w, Update(idI, todo))
}

func UpdateAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	idI, _ := strconv.Atoi(id)
	sendResponse(w, Update(idI, todo))
}

func sendResponse(w http.ResponseWriter, todo Todo) {
	json.NewEncoder(w).Encode(todo)
}

func main() {
	handleRequests()
}
