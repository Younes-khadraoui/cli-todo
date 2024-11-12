package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	dbs "github.com/Younes-khadraoui/hmm/db"
	_ "github.com/mattn/go-sqlite3"
)

func Handler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func GetTodos() ([]dbs.Todo, error) {
	db, err := sql.Open("sqlite3", "todos.db")
	if err != nil {
		log.Println("Couldnt add Todo")
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()
	queries := dbs.New(db)

	todos, err := queries.GetAllTodos(ctx)
	if err != nil {
		log.Println("Couldnt Get All Todos")
		return []dbs.Todo{}, nil
	}

	return todos, nil
}

func NewTodo(todo string) {
	db, err := sql.Open("sqlite3", "todos.db")
	if err != nil {
		log.Println("cant open the database")
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()
	queries := dbs.New(db)

	err = queries.NewTodo(ctx, todo)
	if err != nil {
		log.Println("Couldnt Add Todo")
		return
	}

	fmt.Println("Todo added")
}

func DeleteTodo(id int) {
	db, err := sql.Open("sqlite3", "todos.db")
	if err != nil {
		log.Println("Couldnt add Todo")
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()
	queries := dbs.New(db)

	err = queries.DeleteTodo(ctx, int64(id))
	if err != nil {
		log.Println("Couldnt Delete Todo")
	}

	fmt.Println("Todo Deleted")
}

func DoneTodo() {
}

func EditTodo() {
}
