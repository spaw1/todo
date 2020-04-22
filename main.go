package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type TodoRequest struct {
	Todo string `json:"todo"`
}

type TodoResponse struct {
	Success bool   `json:"success"`
	Body    string `json:"body"`
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/todo", GetAllAction)
	router.HandleFunc("/todo/{id}", GetByIdAction)
	router.HandleFunc("/todo/{id}/markAsComplete", MarkAsCompleteAction)
	router.HandleFunc("/todo/{id}/delete", DeleteAction)
	router.HandleFunc("/todo/create", CreateAction).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func CreateAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body TodoRequest
	json.NewDecoder(r.Body).Decode(&body)
	id := Create(body.Todo)
	fmt.Print(w, id)

}

func GetAllAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todos := GetAll()
	for _, todo := range todos {
		json.NewEncoder(w).Encode(todo)
	}
	fmt.Print(w, todos)
}

func GetByIdAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	todo := GetTodo(id)
	json.NewEncoder(w).Encode(todo)
	fmt.Print(w, todo)
}

func MarkAsCompleteAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	success := MarkAsComplete(id)
	sendResponse(w, success, "")
}

func DeleteAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	success := Delete(id)
	sendResponse(w, success, "")
}

func sendResponse(w http.ResponseWriter, status bool, body string) {
	var response TodoResponse
	response.Success = status
	response.Body = body
	json.NewEncoder(w).Encode(response)
	fmt.Print(w, response)

}
func main() {
	handleRequests()
}
