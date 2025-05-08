package utils

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ParseRESP(reader *bufio.Reader) (interface{}, error) {
		prefix, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
	
		switch prefix {
		case '+': // Simple string
			line, err := reader.ReadString('\n')
			if err != nil {
				return nil, err
			}
			return strings.TrimSuffix(line, "\r\n"), nil
	
		case '-': // Error
			line, err := reader.ReadString('\n')
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("RESP error: %s", strings.TrimSuffix(line, "\r\n"))
	
		case ':': // Integer
			line, err := reader.ReadString('\n')
			if err != nil {
				return nil, err
			}
			return strconv.Atoi(strings.TrimSuffix(line, "\r\n"))
	
		case '$': // Bulk string
			lenLine, err := reader.ReadString('\n')
			if err != nil {
				return nil, err
			}
	
			strLen, err := strconv.Atoi(strings.TrimSuffix(lenLine, "\r\n"))
			if err != nil {
				return nil, err
			}
			if strLen == -1 {
				return nil, nil // null bulk string
			}
	
			buf := make([]byte, strLen+2) // +2 for \r\n
			_, err = io.ReadFull(reader, buf)
			if err != nil {
				return nil, err
			}
			return string(buf[:strLen]), nil
	
		case '*': // Array
			lenLine, err := reader.ReadString('\n')
			if err != nil {
				return nil, err
			}
			arrLen, err := strconv.Atoi(strings.TrimSuffix(lenLine, "\r\n"))
			if err != nil {
				return nil, err
			}
	
			var items []interface{}
			for i := 0; i < arrLen; i++ {
				item, err := parseRESP(reader)
				if err != nil {
					return nil, err
				}
				items = append(items, item)
			}
			return items, nil
	
		default:
			return nil, fmt.Errorf("unknown RESP type: %c", prefix)
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