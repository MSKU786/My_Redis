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
	
		case '*': 
			lenLine, err := reader.ReadString('\n')
			if err != nil {
				return nil, err
			}

			arrLen, err := strconv.Atoi(strings.TrimSuffix(lenLine, "\r\n"))
			if err != nil {
				return nil, err
			}

			if arrLen == -1 {
				return nil, nil
			}

			arr := make([]interface{}, arrLen);
			for i:=0; i<arrLen; i++ {
				value, err := ParseRESP(reader)
				if err != nil {
					return nil, err
				}
				arr[i] = value
			}

			return arr, nil;
			
		default:
			return nil, fmt.Errorf("unknown RESP type: %c", prefix)
		}
}
	

func SerializeRESP(data interface{}) (string, error) {
	if data == nil {
		return "$-1\r\n", nil // Null Bulk String
	}

	switch v := data.(type) {
	case string:
		if v == "OK" {
			return "+OK\r\n", nil // Simple String
		}
		return fmt.Sprintf("$%d\r\n%s\r\n", len(v), v), nil // Bulk String
	case int:
		return fmt.Sprintf(":%d\r\n", v), nil // Integer
	case int64:
		return fmt.Sprintf(":%d\r\n", v), nil
	default:
		return "", fmt.Errorf("unsupported data type for RESP: %T", data)
	}
}
