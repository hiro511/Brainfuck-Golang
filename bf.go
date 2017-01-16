package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

const maxMemory int = 1000

var memory = [maxMemory]byte{}
var mIndex = 0

func main() {
	if len(os.Args) != 2 {
		log.Fatal(errors.New("invalid arguents"))
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// EOF do not read
	program := make([]byte, stat.Size()-1)
	_, err = file.Read(program)
	if err != nil {
		log.Fatal(err)
	}
	parse(program)
}

func parse(program []byte) {
	for pIndex := 0; pIndex < len(program); pIndex++ {
		readOp(program, &pIndex)
	}
}

func readOp(program []byte, pIndex *int) {
	switch program[*pIndex] {
	case '+':
		memory[mIndex]++
	case '-':
		memory[mIndex]--
	case '<':
		mIndex--
	case '>':
		mIndex++
	case '[':
		for memory[mIndex] != 0 {
			pIndexTmp := *pIndex
			for (*pIndex)++; program[*pIndex] != ']'; (*pIndex)++ {
				readOp(program, pIndex)
			}
			*pIndex = pIndexTmp
		}
		moveEndOfLoop(program, pIndex)
	case ']':
	case ',':
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		memory[mIndex] = []byte(input)[0]
	case '.':
		fmt.Print(string(memory[mIndex]))
	default:
		log.Fatal(errors.New("invalid operation"))
	}
}

func moveEndOfLoop(program []byte, pIndex *int) {
	n := 1
	for n > 0 {
		*pIndex++
		switch program[*pIndex] {
		case '[':
			n++
		case ']':
			n--
		}
	}
}
