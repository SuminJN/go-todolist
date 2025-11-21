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
	if err != nil {
		fmt.Printf("JSON 파싱 오류: %v\n", err)
		os.Exit(1)
	}

	return todos
}

// JSON 파일에 할일 목록 저장
func saveTodos(todos []Todo) {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		fmt.Printf("JSON 생성 오류: %v\n", err)
		return
	}

	err = os.WriteFile(todoFile, data, 0644)
	if err != nil {
		fmt.Printf("파일 저장 오류: %v\n", err)
		return
	}
}
