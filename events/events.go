// events/events.go
package events

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

// EventData represents the data structure passed when an event is triggered.
type EventData struct {
	EventType string                 `json:"eventType"`
	ElementID string                 `json:"elementID"`
	Params    map[string]interface{} `json:"params"`
}

// EventBus manages event subscriptions and dispatching.
type EventBus struct {
	subscribers map[string]map[string]func(EventData) error
}

// NewEventBus creates a new EventBus instance.
func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string]map[string]func(EventData) error),
	}
}

// On registers an event handler for a specific event and element ID.
func (eb *EventBus) On(eventType, elementID string, handler func(EventData) error) {
	if _, exists := eb.subscribers[eventType]; !exists {
		eb.subscribers[eventType] = make(map[string]func(EventData) error)
	}
	eb.subscribers[eventType][elementID] = handler
}

// Publish triggers the event handler for a given event.
func (eb *EventBus) Publish(event EventData) error {
	if handlers, exists := eb.subscribers[event.EventType]; exists {
		if handler, exists := handlers[event.ElementID]; exists {
			return handler(event)
		}
	}
	return fmt.Errorf("no handler found for event: %s, element: %s", event.EventType, event.ElementID)
}

// Listen sets up the event listener for JavaScript events.
func (eb *EventBus) Listen() {
	js.Global().Set("handleEvent", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var eventData EventData
		if err := json.Unmarshal([]byte(args[0].String()), &eventData); err != nil {
			fmt.Printf("Error unmarshaling event: %v\n", err)
			return nil
		}
		if err := eb.Publish(eventData); err != nil {
			fmt.Printf("Error handling event: %v\n", err)
		}
		return nil
	}))
}