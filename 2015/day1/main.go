package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	floor := 0

	b, err := r.ReadByte()
	for err == nil {
		if b == '(' {
			floor++
		} else if b == ')' {
			floor--
		} else {
			log.Printf("bad character: %q", b)
		}

		b, err = r.ReadByte()
	}
	fmt.Printf("floor: %v\n", floor)

	log.Fatalf("error while reading bytes: %v", err)
}
