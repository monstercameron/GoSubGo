// main.go
//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"log"

	"github.com/monstercameron/GoSubGo/database"
	"github.com/monstercameron/GoSubGo/events"
	"github.com/monstercameron/GoSubGo/todolist"
	"github.com/monstercameron/GoSubGo/utils"
)

func main() {
	c := make(chan struct{}, 0)

	// Initialize the in-memory SQL.js database
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize the todos table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		completed INTEGER NOT NULL DEFAULT 0
	)`)
	
	// close the database connection when the function returns
	defer db.Close()

	// Initialize the todos table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatalf("Failed to create todos table: %v", err)
	}

	// Initialize the EventBus
	eventBus := events.NewEventBus()

	// Subscribe todo list handlers to the event bus
	todolist.SubscribeAll(eventBus, db)

	// Perform the initial page render
	// if err := InitialPageRender(db, "#root"); err != nil {
	// 	log.Fatalf("Failed to render page: %v", err)
	// }

	// Start listening for JavaScript events
	eventBus.Listen()

	<-c // Prevent the function from returning
}

// InitialPageRender performs the initial render using data from the database.
func InitialPageRender(db *database.DB, selector string) error {
	todos, err := todolist.GetAllTodos(db)
	if err != nil {
		return fmt.Errorf("failed to get todos: %w", err)
	}

	html := todolist.Render(todos)

	return utils.Render(html, selector)
}