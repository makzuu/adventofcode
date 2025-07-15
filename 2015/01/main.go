package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"io"
	"errors"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	floor := 0
	characterPosition := 0
	curPosition := 1

	b, err := r.ReadByte()
	for err == nil {
		if b == '(' {
			floor++
		} else if b == ')' {
			floor--
		} else {
			log.Printf("bad character: %q", b)
		}

		if characterPosition == 0 && floor == -1 {
			characterPosition = curPosition
		}

		b, err = r.ReadByte()
		curPosition++
	}
	fmt.Printf("floor: %v\n", floor)
	fmt.Printf("position: %v\n", characterPosition)

	if !errors.Is(err, io.EOF) {
		log.Printf("error while reading bytes: %v", err)
	}
}
