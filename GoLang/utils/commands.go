package utils 

// Define a function type of commands
type CommandFunc func(args []string) interface{}

var commands = map[string]CommandFunc {
	"SET": func(args []string) interface{} {
			if (len(args) < 2) {
					return "ERR wrong number of arguments for 'set'"
			}

			key, value = args[0],args[1];
			store.Store.store(key, value);
			
	}
}
