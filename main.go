package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func main() {
	//msg := " Hello there my friend"
	ranGen := uuid.New().String()
	ranByt := []byte(ranGen)
	ranByt = ranByt[:16]

	msg := "Hello this is my message for encrytion and decryption"

	rslt, err := encDec(ranByt, msg)
	if err != nil {
		log.Panic(" Error while printing the result value!!")
	}

	fmt.Println(string(rslt))

	rslt2, err := encDec(ranByt, string(rslt))
	if err != nil {
		log.Panic(" Error while printing the result value!!")
	}
	fmt.Println(string(rslt2))
}

func encDec(key []byte, input string) ([]byte, error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("couldn't new cipher %w", err)
	}

	iv := make([]byte, aes.BlockSize)

	s := cipher.NewCTR(b, iv)

	buff := &bytes.Buffer{}
	sw := cipher.StreamWriter{
		S: s,
		W: buff,
	}

	_, err = sw.Write([]byte(input))
	if err != nil {
		return nil, fmt.Errorf(" error while doing sw.write %w", err)
	}

	return buff.Bytes(), nil

}
