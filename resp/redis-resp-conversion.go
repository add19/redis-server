package resp

import (
	"fmt"
	"strconv"
	"strings"
)

type ByteArray []byte
type Integer int
type String string
type BulkString struct {
	Value *String
}
type StringArray []string
type IntegerArray []Integer
type Byte byte
type NilResponse struct {
}

type RedisError struct {
	msg string
}

type Serializable interface {
	Serialize() string
}

func Serialization[T Serializable](input T) string {
	return input.Serialize()
}

func (b Byte) Serialize() string {
	output := fmt.Sprintf(":%d\r\n", b)
	return output
}

func (n NilResponse) Serialize() string {
	output := "*-1\r\n"
	return output
}

func (s BulkString) Serialize() string {
	if s.Value == nil {
		return "$-1\r\n"
	}
	output := fmt.Sprintf("$%d\r\n%s\r\n", len(*s.Value), *s.Value)
	return output
}

func (b IntegerArray) Serialize() string {
	output := fmt.Sprintf("*%d\r\n", len(b))
	for _, val := range b {
		output += val.Serialize()
	}
	return output
}

func (b StringArray) Serialize() string {
	output := fmt.Sprintf("*%d\r\n", len(b))
	for _, val := range b {
		output += fmt.Sprintf("$%d\r\n%s\r\n", len(val), val)
	}
	return output
}

func (i Integer) Serialize() string {
	output := fmt.Sprintf(":%v\r\n", i)
	return output
}

func (i String) Serialize() string {
	output := fmt.Sprintf("+%v\r\n", i)
	return output
}

func (r RedisError) Serialize() string {
	output := fmt.Sprintf("-%v\r\n", r.msg)
	return output
}

// func SerializeByteArray(input interface{}) string {
// 	output := fmt.Sprintf("*%d\r\n", len(input))
// 	for _, val := range input {
// 		output += fmt.Sprintf("$%d\r\n%s\r\n", len(val), val)
// 	}
// 	return output
// }

func DeserializeInt(input string) int {
	m, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error : " + err.Error())
	}
	return m
}

func ParseIntValue(input string) []string {
	return strings.Split(input[1:], "\r\n")
}

func ParseErrorValue(input string) []string {
	return strings.Split(input[1:], "\r\n")
}

func ParseBulkArray(input string) []string {
	vals := strings.Split(input[1:], "\r\n")

	var res []string
	if vals[0] == "-1" {
		return res
	}
	for i := range vals {
		if i%2 == 1 {
			res = append(res, vals[i])
		}
	}
	return res
}

func ParseArray(input string) []string {
	vals := strings.Split(input[1:], "\r\n")
	var res []string

	for i := range vals {
		if i == 0 {
			continue
		}

		if vals[i] != "" && vals[i][0] == 36 {
			res = append(res, vals[i+1])
		} else if vals[i] != "" && vals[i][0] == ':' {
			//to check if any other logic is needed for this type and not just add to res
			res = append(res, vals[i][1:])
		}
	}

	return res
}

func Deserialization(input string) []string {
	dataType := input[0:1]

	switch dataType {
	case "*":
		return ParseArray(input)
	case "$":
		parsedStr := ParseBulkArray(input)
		return parsedStr
	case "+":
		return strings.Split(input[1:], "\r\n")
	case "-":
		errMsg := ParseErrorValue(input)
		return errMsg
	case ":":
		val := ParseIntValue(input)
		intVal, _ := strconv.Atoi(val[0])
		return []string{strconv.Itoa(intVal)}
	default:
		fmt.Println("Unknown type")
	}
	return []string{input}
}
