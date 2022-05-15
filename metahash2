package main

import ("fmt" 
		"math"
		"hash"
		"crypto")



//create a function for hashing the input data
func hash_input(input string) string {
	h := crypto.Sha256.New()
	h.Write([]byte(input))
	return string(h.Sum(nil))
}

//crearte a function and this function is used to divide the hashed output data into blocks
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

//create a function that take the input data and divide it into blocks
func hash_list(input string) HashList {
	var hashlist HashList
	var blocks []string
	blocks = divide_input(input)
	for i := 0; i < len(blocks); i++ {
		switch i {
		case 0:
			hashlist.Hash1 = []byte(hash_input(blocks[i]))
		case 1:
			hashlist.Hash2 = []byte(hash_input(blocks[i]))
		case 2:
			hashlist.Hash3 = []byte(hash_input(blocks[i]))
		case 3:
			hashlist.Hash4 = []byte(hash_input(blocks[i]))
		case 4:
			hashlist.Hash5 = []byte(hash_input(blocks[i]))
		case 5:
			hashlist.Hash6 = []byte(hash_input(blocks[i]))
		case 6:
			hashlist.Hash7 = []byte(hash_input(blocks[i]))
		case 7:
			hashlist.Hash8 = []byte(hash_input(blocks[i]))
		case 8:
			hashlist.Hash9 = []byte(hash_input(blocks[i]))
					case 9:
						hashlist.Hash10 = []byte(hash_input(blocks[i]))
					case 10:
						hashlist.Hash11 = []byte(hash_input(blocks[i]))
		}

	}
	return hashlist
}


