package utils

// Adjusted to proper Go module path

import (
	"strconv"
	"sync"
	"time"
)

var (
	Store       = sync.Map{}
	ExpireStore = sync.Map{}
) // Assuming NewStore initializes the store

// Define a function type of commands
type CommandFunc func(args []string) interface{}

var commands = map[string]CommandFunc{
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
					return "Error: GET command requires 1 argument"
			}

			key := args[0]
			value, exists := Store.Load(key)
			ttl, ttlExists := ExpireStore.Load(key)

			if ttlExists {
					expirationTime, ok := ttl.(int64)
					if !ok {
							return "Error: Invalid TTL format"
					}
					if expirationTime < time.Now().Unix() {
							Store.Delete(key)
							ExpireStore.Delete(key)
							return nil
					}
			}

			if exists {
					return value
			}
			return nil
	},
	"EXPIRE": func(args []string) interface{} {
			if len(args) < 2 {
					return "Error: EXPIRE command requires 2 arguments"
			}

			key, seconds := args[0], args[1]
			ttl, err := strconv.Atoi(seconds)
			if err != nil {
					return "Error: Invalid TTL value"
			}

			_, ok := Store.Load(key)
			if ok {
					currentTime := time.Now().Unix()
					expirationTime := currentTime + int64(ttl)
					ExpireStore.Store(key, expirationTime)
					return 1
			}
			return 0
	},
	"DEL": func(args []string) interface{} {
			var count = 0
			for _, key := range args {
					_, ok := Store.Load(key)
					if ok {
							Store.Delete(key)
							ExpireStore.Delete(key)
							count++
					}
			}
			return count
	},
}