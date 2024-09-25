// todolist/todos.go
package todolist

import (
	"database/sql"
	"fmt"

	"github.com/monstercameron/GoSubGo/events"
	"github.com/monstercameron/GoSubGo/utils"
)

// Todo represents a single todo item.
type Todo struct {
	ID    int
	Title string
}

// SubscribeAll subscribes todo list handlers to the event bus.
func SubscribeAll(eb *events.EventBus, db *sql.DB) {
	eb.On("submit", "todo-form", func(event events.EventData) error {
		title, ok := event.Params["title"].(string)
		if !ok || title == "" {
			return fmt.Errorf("invalid or missing 'title' parameter")
		}

		// Insert into database
		_, err := db.Exec("INSERT INTO todos (title) VALUES (?)", title)
		if err != nil {
			return fmt.Errorf("failed to insert todo: %w", err)
		}

		// Retrieve current todos
		todos, err := GetAllTodos(db)
		if err != nil {
			return err
		}

		// Render the todo list
		html := Render(todos)
		if err := utils.Render(html, "#todo-list"); err != nil {
			return err
		}

		return nil
	})
}

// GetAllTodos retrieves all todos from the database.
func GetAllTodos(db *sql.DB) ([]Todo, error) {
	rows, err := db.Query("SELECT id, title FROM todos")
	if err != nil {
		return nil, fmt.Errorf("failed to query todos: %w", err)
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title); err != nil {
			return nil, fmt.Errorf("failed to scan todo: %w", err)
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

// Render returns the HTML representation of the todo list.
func Render(todos []Todo) string {
	html := "<ul>"
	for _, todo := range todos {
		html += fmt.Sprintf("<li>%s</li>", todo.Title)
	}
	html += "</ul>"
	return html
}