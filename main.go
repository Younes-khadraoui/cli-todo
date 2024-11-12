package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Younes-khadraoui/hmm/handler"
)

func main() {
	showTodos()
	for {
		listenToUser()
	}
}

func showTodos() {
	todos, err := handler.GetTodos()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if len(todos) == 0 {
		fmt.Println("No Todos yet -_-")
	} else {
		for _, todo := range todos {
			if todo.Isdone.Int64 == 1 {
				fmt.Printf("[X]")
			} else {
				fmt.Printf("[]")
			}
			fmt.Printf(" %d.", todo.ID)
			fmt.Printf(" %s", todo.Content)
			fmt.Println()
		}
	}
}

func listenToUser() {
	fmt.Printf("\n \n> ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		checkInput(input)
		fmt.Printf("\n \n> ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func checkInput(input string) {
	if strings.HasPrefix(input, "new ") {
		todo := string(input[4:])
		handler.NewTodo(todo)
	} else if strings.HasPrefix(input, "delete ") {
		id, err := strconv.Atoi(input[7:])
		if err != nil {
			log.Println("Couldnt convert string to integer")
			panic(err)
		}
		handler.DeleteTodo(id)
	} else if strings.HasPrefix(input, "done ") {
		handler.DoneTodo()
	} else if strings.HasPrefix(input, "edit ") {
		handler.EditTodo()
	} else if input == "help" {
		fmt.Printf("\n \nHELP: \nnew {todo} to add new todo \ndelete {id} to delete \ndone {id} to mark done ")
	} else if input == "quit" || input == "exit" || input == "q" {
		fmt.Println("\nbye ...")
		os.Exit(1)
	} else if input == "show" || input == "todo" || input == "todos" {
		showTodos()
	} else {
		fmt.Println("wrong input")
	}
}
