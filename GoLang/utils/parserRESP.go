
package utils

import (
	"fmt"
)

func parserRESP(data interface{})  interface{} {

}


if data == nil {
	return "$-1\r\n"
}

switch v := data.(type) {
case string:
	if v == "OK" {
		return "+OK\r\n"
	}
	return fmt.Sprintf("$%d\r\n%s\r\n", len(v), v)
case int:
	return fmt.Sprintf(":%d\r\n", v)
case int64:
	return fmt.Sprintf(":%d\r\n", v)
default:
	return fmt.Sprintf("Unsupported data type for RESP: %T", data)
}