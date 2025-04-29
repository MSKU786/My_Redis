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
		"EXPIRE": func(args []string) interface{} { 
				key, seconds := args[0], args[1]

				ttl , err := strconv.Atoi(seconds)
				if (err != nil) {
					return nil
				}
				ttl = ttl * 1000;

				value, ok := Store[key];
				if (ok) {
					var currentTime = time.Now().Unix();
					ExpireStore[key] =  currentTime + ttl;
					return 1;
				}
				return 0;
		},
		"DEL": (args) => {
			var count = 0
			for i : range(args) {
				_, ok := Store[i];
				if (ok) {
					Store.Delete(i);
					ExpireStore.Delete(i);
					count += 1
				}
			}
			return count;
		}
}
