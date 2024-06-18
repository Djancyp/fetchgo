package gofetch

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	// add assertion library
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func TestFetchGet(t *testing.T) {

	gofetch, err := New(Config{
		BaseUrl: "https://jsonplaceholder.typicode.com",
	})
	if err != nil {
		t.Error(err)
	}
	todosStr, err := gofetch.DoRequest("GET", "/todos", "", nil)
	if err != nil {
		t.Error(err)
	}
	todos := []Todo{}
	bytes := []byte(todosStr)
	err = json.Unmarshal(bytes, &todos)
	if err != nil {
		t.Error(err)
	}
}

func TestFetchPost(t *testing.T) {
	gofetch, err := New(Config{
		BaseUrl: "https://jsonplaceholder.typicode.com",
	})
	if err != nil {
		t.Error(err)
	}
	todo := Todo{
		UserID:    1,
		ID:        1,
		Title:     "delectus aut autem",
		Completed: false,
	}
	todoStr, err := json.Marshal(todo)

	if err != nil {
		t.Error(err)
	}
	todosStr, err := gofetch.DoRequest("POST", "/todos", string(todoStr), nil)
	if err != nil {
		t.Error(err)
	}
	todos := Todo{}
	bytes := []byte(todosStr)
	err = json.Unmarshal(bytes, &todos)
	if err != nil {
		t.Error(err)
	}
}

func TestFetchPut(t *testing.T) {
	gofetch, err := New(Config{
		BaseUrl: "https://jsonplaceholder.typicode.com",
	})
	if err != nil {
		t.Error(err)
	}
	todo := Todo{
		UserID:    1,
		ID:        1,
		Title:     "delectus aut autem",
		Completed: true,
	}
	todoStr, err := json.Marshal(todo)

	if err != nil {
		t.Error(err)
	}
	todosStr, err := gofetch.DoRequest("PUT", "/todos/1", string(todoStr), nil)
	if err != nil {
		t.Error(err)
	}
	todos := Todo{}
	bytes := []byte(todosStr)
	err = json.Unmarshal(bytes, &todos)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, todos.Completed, false)
}

func TestFetchDelete(t *testing.T) {
	gofetch, err := New(Config{
		BaseUrl: "https://jsonplaceholder.typicode.com",
	})
	if err != nil {
		t.Error(err)
	}
	todosStr, err := gofetch.DoRequest("DELETE", "/todos/1", "", nil)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, todosStr, "{}")
}
