package main

import "fmt"

// 새로운 할일 추가
func addTodo(todos []Todo, title string) []Todo {
	// 새 ID는 현재 최대 ID + 1
	newID := 1
	for _, todo := range todos {
		if todo.ID >= newID {
			newID = todo.ID + 1
		}
	}

	// 새 할일 생성
	newTodo := Todo{
		ID:        newID,
		Title:     title,
		Completed: false,
	}

	// 목록에 추가
	todos = append(todos, newTodo)
	fmt.Printf("할일이 추가되었습니다 (ID: %d)\n", newID)
	return todos
}

// 할일 목록 출력
func listTodos(todos []Todo) {
	if len(todos) == 0 {
		fmt.Println("할일이 없습니다")
		return
	}

	fmt.Println("\n=== 할일 목록 ===")
	for _, todo := range todos {
		status := "[ ]"
		if todo.Completed {
			status = "[✓]"
		}
		fmt.Printf("%s %d. %s\n", status, todo.ID, todo.Title)
	}
	fmt.Println()
}

// 할일을 완료로 표시
func markDone(todos []Todo, id int) []Todo {
	found := false
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = true
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("ID %d인 할일을 찾을 수 없습니다\n", id)
		return todos
	}

	fmt.Printf("할일 %d를 완료했습니다\n", id)
	return todos
}

// 할일 삭제
func deleteTodo(todos []Todo, id int) []Todo {
	found := false
	newTodos := []Todo{}

	for _, todo := range todos {
		if todo.ID == id {
			found = true
			continue
		}
		newTodos = append(newTodos, todo)
	}

	if !found {
		fmt.Printf("ID %d인 할일을 찾을 수 없습니다\n", id)
		return todos
	}

	fmt.Printf("할일 %d를 삭제했습니다\n", id)
	return newTodos
}

// 사용법 출력
func printHelp() {
	fmt.Println("\n=== TODO 리스트 사용법 ===")
	fmt.Println("go run . list          - 할일 목록 보기")
	fmt.Println("go run . add <할일>     - 새로운 할일 추가")
	fmt.Println("go run . done <id>     - 할일 완료 표시")
	fmt.Println("go run . delete <id>   - 할일 삭제")
	fmt.Println()
}
