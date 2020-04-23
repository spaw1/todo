package main

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type TodoRequest struct {
	Message string `json:"desc"`
}

type TodoResponse struct {
	Success bool   `json:"success"`
	Body    string `json:"body"`
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/todo", GetAllAction).Methods("GET")
	router.HandleFunc("/todo/{id}", GetByIdAction).Methods("GET")
	router.HandleFunc("/todo/{id}/updateStatus", UpdateAction).Methods("PATCH")
	router.HandleFunc("/todo/create", CreateAction).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "PATCH", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

func CreateAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body TodoRequest
	json.NewDecoder(r.Body).Decode(&body)
	var todoItem = Create(body.Message)
	json.NewEncoder(w).Encode(todoItem)
	w.WriteHeader(http.StatusCreated)
}

func GetAllAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	todos := GetAll()
	json.NewEncoder(w).Encode(todos)
}

func GetByIdAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	idI, _ := strconv.Atoi(id)
	todo := GetTodo(idI)
	sendResponse(w, todo)
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
