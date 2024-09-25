// todolist/todos.go
package todolist

import (
	"fmt"
	"strconv"

	"github.com/monstercameron/GoSubGo/database"
	"github.com/monstercameron/GoSubGo/events"
)

// Todo represents a single todo item.
type Todo struct {
	ID    int64
	Title string
}

// SubscribeAll subscribes todo list handlers to the event bus.
func SubscribeAll(eb *events.EventBus, db *database.DB) {
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
		return RenderToDOM(html, "#todo-list")
	})
}

// RenderInitialPage performs the initial render using data from the database.
func RenderInitialPage(db *database.DB, selector string) error {
	todos, err := GetAllTodos(db)
	if err != nil {
		return fmt.Errorf("failed to get todos: %w", err)
	}

	html := Render(todos)

	return RenderToDOM(html, selector)
}

// GetAllTodos retrieves all todos from the database.
func GetAllTodos(db *database.DB) ([]Todo, error) {
	rows, err := db.Query("SELECT id, title FROM todos")
	if err != nil {
		return nil, fmt.Errorf("failed to query todos: %w", err)
	}

	var todos []Todo
	for _, row := range rows {
		id, err := strconv.ParseInt(fmt.Sprint(row["id"]), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse todo ID: %w", err)
		}
		todo := Todo{
			ID:    id,
			Title: fmt.Sprint(row["title"]),
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

// RenderToDOM renders the given HTML to the specified DOM element.
func RenderToDOM(html, selector string) error {
	// This function should be implemented in a separate utils package
	// that handles DOM manipulation in the WebAssembly environment.
	// For now, we'll leave it as a placeholder.
	return nil
}