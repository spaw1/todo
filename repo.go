package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
)

type Todo struct {
	Id      int    `json:"id"`
	Message string `json:"desc"`
	Status  string `json:"status"`
}

type DbConfig struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

func connect() (db *sql.DB) {
	var config = getConfig()
	dbDriver := "mysql"
	dbName := "online"
	db, err := sql.Open(dbDriver, config.Database.User+":"+config.Database.Password+"@tcp("+config.Database.Host+":"+config.Database.Port+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func getConfig() DbConfig {
	f, _ := os.OpenFile("./config.yaml", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// Todo: Error Handling, logging
	defer f.Close()
	var dbConfig DbConfig
	decoder := yaml.NewDecoder(f)
	decoder.Decode(&dbConfig)
	return dbConfig
}

func CreateTodoEntity(message string) int {
	db := connect()
	stmt, err := db.Exec("INSERT INTO TODO (message) VALUES (?)", message)
	if err != nil {
		panic(err.Error())
	}
	id, err := stmt.LastInsertId()
	defer db.Close()
	return int(id)
}

func GetAllEntities() []Todo {
	db := connect()
	stmt, err := db.Query("SELECT id,message,status FROM TODO WHERE status <> 'deleted'")
	if err != nil {
		panic(err.Error())
	}
	var Todos []Todo
	for stmt.Next() {
		var todo Todo
		err = stmt.Scan(&todo.Id, &todo.Message, &todo.Status)
		if err != nil {
			panic(err.Error())
		}
		Todos = append(Todos, todo)
	}
	defer stmt.Close()
	defer db.Close()
	return Todos
}

func GetTodoEntity(id int) Todo {
	db := connect()
	row := db.QueryRow("SELECT id,message,status FROM TODO WHERE id = ? AND status <> 'deleted'", int64(id))
	var todo Todo
	row.Scan(&todo.Id, &todo.Message, &todo.Status)
	defer db.Close()
	return todo
}

func MarkAsCompleteEntity(id string) bool {
	db := connect()
	idI, _ := strconv.Atoi(id)
	stmt, _ := db.Exec("UPDATE TODO SET status = 'completed' WHERE id = ?", idI)
	rows, _ := stmt.RowsAffected()
	defer db.Close()
	if rows > 0 {
		return true
	}
	return false
}

func DeleteEntity(id string) bool {
	db := connect()
	idI, _ := strconv.Atoi(id)
	stmt, _ := db.Exec("UPDATE TODO SET status = 'deleted' WHERE id = ?", idI)
	rows, _ := stmt.RowsAffected()
	defer db.Close()
	if rows > 0 {
		return true
	}
	return false
}

func UpdateEntity(id int, todo Todo) {
	db := connect()
	db.Exec("UPDATE TODO SET message = ?, status = ? WHERE id = ?", todo.Message, todo.Status, id)
	defer db.Close()
}
