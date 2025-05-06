package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ParserRESP(data []byte) ([]string, error)  {
	fmt.Println("debug 0", data);
	str := string(data)
	fmt.Println("debug 1", str);
	lines := strings.Split(str, "\r\n");
	fmt.Println("debug 1", lines);
	if len(lines) < 1 {
		return nil, errors.New("invalid input")
	}

	switch lines[0][0] {
		case '*' : 
			count, err := strconv.Atoi(lines[0][1:])
			if err != nil {
				return nil, fmt.Errorf("invalid array count: %v", err)
			}

			var elements []string;

			for i:=0; i<count; i++ {
				// RESP Array follows format: *<count>\r\n$<len>\r\n<value>\r\n
				// Starting from line 1: lines[1] = $len, lines[2] = value, etc.
				valIndex := 2*i + 2
				if valIndex >= len(lines) {
					return nil, errors.New("malformed RESP array")
				}
				elements = append(elements, lines[valIndex])
			}
			return elements, nil
		default:
			return nil, fmt.Errorf("unknown RESP type: %c", lines[0][0])
	}

}


func SerializeRESP(data interface{}) interface{} {
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
}