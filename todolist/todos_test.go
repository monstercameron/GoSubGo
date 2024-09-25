// tests/todolist_test.go
package todolist_test

import (
	"database/sql"
	"myapp/database"
	"myapp/todolist"
	"testing"
)

func TestGetAllTodos(t *testing.T) {
	db, err := database.NewDatabase()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Insert test data
	_, err = db.Exec("INSERT INTO todos (title) VALUES (?)", "Test Todo 1")
	if err != nil {
		t.Fatalf("Failed to insert test todo: %v", err)
	}
	_, err = db.Exec("INSERT INTO todos (title) VALUES (?)", "Test Todo 2")
	if err != nil {
		t.Fatalf("Failed to insert test todo: %v", err)
	}

	todos, err := todolist.GetAllTodos(db)
	if err != nil {
		t.Fatalf("GetAllTodos failed: %v", err)
	}

	if len(todos) != 2 {
		t.Fatalf("Expected 2 todos, got %d", len(todos))
	}

	expectedTitles := []string{"Test Todo 1", "Test Todo 2"}
	for i, todo := range todos {
		if todo.Title != expectedTitles[i] {
			t.Errorf("Expected todo title '%s', got '%s'", expectedTitles[i], todo.Title)
		}
	}
}
