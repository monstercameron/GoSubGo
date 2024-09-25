// database/database.go
package database

import (
	"fmt"
	"syscall/js"
)

// DB represents a connection to the SQL.js database
type DB struct {
	sqljs js.Value
	db    js.Value
}

// NewDatabase initializes a new SQL.js database
func NewDatabase() (*DB, error) {
	sqljs := js.Global().Get("SQL")
	if !sqljs.Truthy() {
		return nil, fmt.Errorf("SQL.js is not loaded or not available")
	}

	// Check if SQL.js is using the new API (v1.x.x and above)
	if sqljs.Get("Database").Truthy() {
		db := sqljs.Get("Database").New()
		if !db.Truthy() {
			return nil, fmt.Errorf("failed to create new SQL.js database")
		}
		return &DB{sqljs: sqljs, db: db}, nil
	}

	// Fallback to older API
	db := sqljs.New()
	if !db.Truthy() {
		return nil, fmt.Errorf("failed to create new SQL.js database")
	}
	return &DB{sqljs: sqljs, db: db}, nil
}

// Exec executes a query without returning any rows
func (db *DB) Exec(query string, args ...interface{}) (js.Value, error) {
	jsArgs := make([]interface{}, len(args)+1)
	jsArgs[0] = query
	for i, arg := range args {
		jsArgs[i+1] = db.convertArg(arg)
	}

	result := db.db.Call("run", jsArgs...)
	if !result.Truthy() {
		return js.Undefined(), fmt.Errorf("failed to execute query")
	}

	return result, nil
}

// Query executes a query that returns rows
func (db *DB) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	jsArgs := make([]interface{}, len(args)+1)
	jsArgs[0] = query
	for i, arg := range args {
		jsArgs[i+1] = db.convertArg(arg)
	}

	stmt := db.db.Call("prepare", query)
	if !stmt.Truthy() {
		return nil, fmt.Errorf("failed to prepare statement: %s", query)
	}
	defer stmt.Call("free")

	rows := stmt.Call("get", jsArgs[1:]...)
	if !rows.Truthy() {
		return nil, fmt.Errorf("failed to execute query")
	}

	return db.parseRows(rows), nil
}

// Close closes the database connection
func (db *DB) Close() error {
	if db.db.Get("close").Truthy() {
		db.db.Call("close")
	}
	return nil
}

// convertArg converts a single Go argument to a JavaScript value
func (db *DB) convertArg(arg interface{}) interface{} {
	switch v := arg.(type) {
	case int, int64, float64, bool, string:
		return v
	case []byte:
		return js.Global().Get("Uint8Array").New(len(v))
	default:
		return js.Null()
	}
}

// parseRows converts SQL.js result rows to Go slice of maps
func (db *DB) parseRows(rows js.Value) []map[string]interface{} {
	var result []map[string]interface{}
	for i := 0; i < rows.Length(); i++ {
		row := make(map[string]interface{})
		rowObj := rows.Index(i)
		keys := js.Global().Get("Object").Call("keys", rowObj)
		for j := 0; j < keys.Length(); j++ {
			key := keys.Index(j).String()
			row[key] = db.convertJSToGo(rowObj.Get(key))
		}
		result = append(result, row)
	}
	return result
}

// convertJSToGo converts a JavaScript value to a Go value
func (db *DB) convertJSToGo(jsVal js.Value) interface{} {
	switch jsVal.Type() {
	case js.TypeNumber:
		return jsVal.Float()
	case js.TypeBoolean:
		return jsVal.Bool()
	case js.TypeString:
		return jsVal.String()
	case js.TypeNull, js.TypeUndefined:
		return nil
	default:
		return jsVal
	}
}