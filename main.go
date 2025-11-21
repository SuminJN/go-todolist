package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	todos := loadTodos()

	// 명령어가 없으면 사용법 출력
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	// 첫 번째 인자가 명령어
	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("사용법: go run . add <할일>")
			return
		}
		title := strings.Join(os.Args[2:], " ")
		todos = addTodo(todos, title)
		saveTodos(todos)
	case "list":
		listTodos(todos)
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("사용법: go run . done <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID는 숫자여야 합니다")
			return
		}
		todos = markDone(todos, id)
		saveTodos(todos)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("사용법: go run . delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID는 숫자여야 합니다")
			return
		}
		todos = deleteTodo(todos, id)
		saveTodos(todos)
	default:
		fmt.Printf("알 수 없는 명령어: %s\n", command)
		printHelp()
	}
}
