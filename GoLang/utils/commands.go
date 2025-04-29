package utils

// Adjusted to proper Go module path

import (
	"sync"
	"time"
)

var (
	Store       = sync.Map{}
	ExpireStore = sync.Map{}
) // Assuming NewStore initializes the store

// Define a function type of commands
type CommandFunc func(args []string) interface{}

var commands = map[string]CommandFunc {
	"SET": func(args []string) interface{} {
			if len(args) < 2 {
				return "Error: SET command requires at least 2 arguments"
			}

			key, value := args[0], args[1]
			Store.Store(key, value)
			return "OK"
	},
	"GET": func(args []string) interface{} {
			if len(args) < 1 {
				return "Error: GET command requires at least 1 argument"
			}
	
			key := args[0]
			value, exists := Store.Load(key) // Using Store.Load to retrieve the value and check existence
			ttl, ttlExists := ExpireStore.Load(key) // Using ExpireStore.Load to retrieve TTL and check existence
	
			if ttlExists {
				if ttl.(int64) < time.Now().Unix() { // Assuming TTL is stored as int64 (Unix timestamp)
					Store.Delete(key) // Assuming Store.Delete removes the key
					ExpireStore.Delete(key) // Assuming ExpireStore.Delete removes the TTL
					return nil
				}
			}
	
			if exists {
				return value
			}
			return nil
		},
}
