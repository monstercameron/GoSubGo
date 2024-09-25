// utils/utils.go
package utils

import (
	"fmt"
	"syscall/js"
)

// Render injects the given HTML string into the specified DOM element identified by the selector.
func Render(html, selector string) error {
	document := js.Global().Get("document")
	if !document.Truthy() {
		return fmt.Errorf("document is not available")
	}
	element := document.Call("querySelector", selector)
	if !element.Truthy() {
		return fmt.Errorf("element with selector '%s' not found", selector)
	}
	element.Set("innerHTML", html)
	return nil
}