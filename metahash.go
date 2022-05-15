package main

import (
	"crypto/sha256"
)

//create a function for hashing input data
func hash_input(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	return string(h.Sum(nil))
}

//create a function for divide the input data into blocks
func divide_input(input string) []string {
	var blocks []string
	var block string
	for i := 0; i < len(input); i++ {
		block = block + string(input[i])
		if len(block) == 8 {
			blocks = append(blocks, block)
			block = ""
		}
	}
	return blocks
}

//create a struct that holds the block
type HashList struct {
	Hash1  []byte
	Hash2  []byte
	Hash3  []byte
	Hash4  []byte
	Hash5  []byte
	Hash6  []byte
	Hash7  []byte
	Hash8  []byte
	Hash9  []byte
	Hash10 []byte
	Hash11 []byte
}

//create a function that take
