// Package of Redis Serialization Protocol (RESP)
package resp

import (
	"bufio"
	"fmt"
)

type Type byte

const (
	// Ex: PING
	SimpleString Type = '+'
	// Number of characters in a word
	BulkString Type = '$'
	// Number of sets of word
	ArrayString Type = '*'
	Integer     Type = ':'
	Error       Type = '-'
)

type Value struct {
	Type  Type
	Bytes []byte
	array []Value
}

// String converts resp value into a string.
// If Value cannot be converted, return empty string
func (v Value) String() string {
	if v.Type == BulkString || v.Type == SimpleString {
		return string(v.Bytes)
	}

	return ""
}

// Array() converts Value to an array.
// If Value cannot be converted, returns empty array.
func (v Value) Array() []Value {
	if v.Type == ArrayString {
		return v.array
	}

	return []Value{}
}

func Decode(input *bufio.Reader) (Value, error) {
	prefixByte, err := input.ReadByte()
	if err != nil {
		return Value{}, err
	}

	switch string(prefixByte) {
	case string(SimpleString):
		return decodeSimpleString(input)
		// case string(BulkString):
		// 	return Value{}, nil
		// case string(ArrayString):
		// 	return Value{}, nil
	}

	return Value{}, nil
}

func decodeSimpleString(input *bufio.Reader) (Value, error) {

	fmt.Println("decodeSimpleString")

	bytes, err := readUntilCRLF(input)
	if err != nil {
		return Value{}, nil
	}

	fmt.Println("Bytes", bytes, string(bytes))

	return Value{}, nil
	// return Value{
	// 	Type:  SimpleString,
	// 	Bytes: bytes,
	// }, nil

}

// func decodeBulkString(input *bufio.Reader) (Value, error) {

// }

func readUntilCRLF(input *bufio.Reader) ([]byte, error) {
	fmt.Println("readUntilCRLF")

	content := []byte{}

	// for {
	// fmt.Println("Starting loop.....")
	b, err := input.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	fmt.Println("READING", string(b))

	content = append(content, b...)
	fmt.Println("EOF", string(content[len(content)-2]))
	fmt.Println("isBreaking?", rune(content[len(content)-2]) == '\r')
	// 	if len(content) >= 2 && rune(content[len(content)-2]) == '\r' {
	// 		break
	// 	}
	// }

	fmt.Println("Content", content, string(content))

	// Return content minus CRLF
	return content[:len(content)-2], nil
}
