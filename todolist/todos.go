// todolist/todos.go
package todolist

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/monstercameron/GoSubGo/database"
	"github.com/monstercameron/GoSubGo/events"
	"github.com/monstercameron/GoSubGo/utils"
)

// Todo represents a single todo item.
type Todo struct {
	ID        int64
	Title     string
	Completed bool
}

// SubscribeAll subscribes todo list handlers to the event bus.
func SubscribeAll(eb *events.EventBus, db *database.DB) {
	// Handler for adding a new todo
	eb.On("submit", "todo-form", func(event events.EventData) error {
		title, ok := event.Params["title"].(string)
		if !ok || title == "" {
			return fmt.Errorf("invalid or missing 'title' parameter")
		}

		// Insert into database
		_, err := db.Exec("INSERT INTO todos (title, completed) VALUES (?, ?)", title, 0)
		if err != nil {
			return fmt.Errorf("failed to insert todo: %w", err)
		}

		return RenderAndUpdate(db)
	})

	// Handler for toggling todo completion
	eb.On("change", "", func(event events.EventData) error {
		if strings.HasPrefix(event.ElementID, "todo-checkbox-") {
			idStr := strings.TrimPrefix(event.ElementID, "todo-checkbox-")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid todo ID: %v", idStr)
			}

			// Toggle completed status
			_, err = db.Exec("UPDATE todos SET completed = NOT completed WHERE id = ?", id)
			if err != nil {
				return fmt.Errorf("failed to update todo: %w", err)
			}

			return RenderAndUpdate(db)
		}
		return nil
	})

	// Handler for deleting a todo
	eb.On("click", "", func(event events.EventData) error {
		if strings.HasPrefix(event.ElementID, "delete-todo-") {
			idStr := strings.TrimPrefix(event.ElementID, "delete-todo-")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid todo ID: %v", idStr)
			}

			// Delete from database
			_, err = db.Exec("DELETE FROM todos WHERE id = ?", id)
			if err != nil {
				return fmt.Errorf("failed to delete todo: %w", err)
			}

			return RenderAndUpdate(db)
		}
		return nil
	})
}

// RenderAndUpdate renders the todo list and updates the tasks count.
func RenderAndUpdate(db *database.DB) error {
	// Retrieve current todos
	todos, err := GetAllTodos(db)
	if err != nil {
		return err
	}

	// Render the todo list
	html := Render(todos)
	if err := RenderToDOM(html, "#todo-list"); err != nil {
		return err
	}

	// Render the tasks count
	tasksCountHtml := RenderTasksCount(todos)
	if err := RenderToDOM(tasksCountHtml, "#tasks-count"); err != nil {
		return err
	}

	return nil
}

// GetAllTodos retrieves all todos from the database.
func GetAllTodos(db *database.DB) ([]Todo, error) {
	rows, err := db.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		return nil, fmt.Errorf("failed to query todos: %w", err)
	}

	var todos []Todo
	for _, row := range rows {
		id, err := strconv.ParseInt(fmt.Sprint(row["id"]), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse todo ID: %w", err)
		}
		completedInt, err := strconv.Atoi(fmt.Sprint(row["completed"]))
		if err != nil {
			return nil, fmt.Errorf("failed to parse completed field: %w", err)
		}
		todo := Todo{
			ID:        id,
			Title:     fmt.Sprint(row["title"]),
			Completed: completedInt != 0,
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

// Render returns the HTML representation of the todo list.
func Render(todos []Todo) string {
	html := ""
	for _, todo := range todos {
		// Determine if the todo is completed
		isChecked := ""
		labelClass := "ml-3 block text-gray-900 flex-grow"
		if todo.Completed {
			isChecked = "checked"
			labelClass += " line-through"
		}

		html += fmt.Sprintf(`
<li class="flex items-center bg-gray-100 p-3 rounded-lg shadow">
    <input id="todo-%d" type="checkbox" class="form-checkbox h-5 w-5 text-blue-600" %s data-trigger="change" data-element-id="todo-checkbox-%d">
    <label for="todo-%d" class="%s">
        %s
    </label>
    <button class="text-red-500 hover:text-red-700" data-trigger="click" data-element-id="delete-todo-%d">
        <i class="fas fa-trash"></i>
    </button>
</li>
`, todo.ID, isChecked, todo.ID, todo.ID, labelClass, todo.Title, todo.ID)
	}
	return html
}

// RenderTasksCount renders the tasks count.
func RenderTasksCount(todos []Todo) string {
	remaining := 0
	for _, todo := range todos {
		if !todo.Completed {
			remaining++
		}
	}
	tasksWord := "tasks"
	if remaining == 1 {
		tasksWord = "task"
	}
	return fmt.Sprintf("%d %s remaining", remaining, tasksWord)
}

// RenderToDOM renders the given HTML to the specified DOM element.
func RenderToDOM(html, selector string) error {
	return utils.Render(html, selector)
}
