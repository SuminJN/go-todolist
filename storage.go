package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const todoFile = "todos.json"

// JSON 파일에서 할일 목록 읽기
func loadTodos() []Todo {
	data, err := os.ReadFile(todoFile)
	if err != nil {
		// 파일이 없다면 빈 배열로 시작
		if os.IsNotExist(err) {
			return []Todo{}
		}
		fmt.Printf("파일 읽기 오류: %v\n", err)
		os.Exit(1)
	}

	var todos []Todo
	err = json.Unmarshal(data, &todos)

	return todos
}

// JSON 파일에 할일 목록 저장
func saveTodos(todos []Todo) {
	data, _ := json.MarshalIndent(todos, "", "  ")

	_ = os.WriteFile(todoFile, data, 0644)
}
