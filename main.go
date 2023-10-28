package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	msg := " Hello there my friend"
	encoded := encode(msg)
	fmt.Println("Encoded Message", encoded)

	s, err := decode(encoded)
	if err != nil {
		fmt.Println(" error outputing decoded message")
	}

	fmt.Println("our Decoded message ", s)
}

func encode(msg string) string {
	return base64.StdEncoding.EncodeToString([]byte(msg))
}

func decode(encoded string) (string, error) {
	s, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return "", fmt.Errorf("decoding failed: %w", err)
	}

	return string(s), nil
}
