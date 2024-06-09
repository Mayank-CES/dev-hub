package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	payload := []byte("hello I am Mayank")
	hashAndBroadcast(bytes.NewReader(payload))
}

func hashAndBroadcast(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	hash := sha1.Sum(b)
	fmt.Println(hex.EncodeToString(hash[:]))

	return broadcast(r)
}

func broadcast(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Println("string of the bytes: ", string(b))

	return broadcast(r)
}
