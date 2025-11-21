package main

import "fmt"

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

// 사용법 출력
func printHelp() {
	fmt.Println("\n=== TODO 리스트 사용법 ===")
	fmt.Println("go run . list          - 할일 목록 보기")
	fmt.Println("go run . add <할일>     - 새로운 할일 추가")
	fmt.Println("go run . done <id>     - 할일 완료 표시")
	fmt.Println("go run . delete <id>   - 할일 삭제")
	fmt.Println()
}
