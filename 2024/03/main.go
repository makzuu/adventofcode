package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

const filename = "input"

const (
	OPERATION = iota
	N1
	N2
)

func main() {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(f)
	result := 0
	state := OPERATION

	for sc.Scan() {
		line := sc.Text()

		buf := ""
		n1Buf := ""
		n2Buf := ""
		target := "mul"
		for _, r := range line {
			switch state {
			case OPERATION:
				if r == '(' {
					start := len(buf) - len(target)
					if start >= 0 && start < len(buf) && buf[start:] == target {
						state = N1
						buf = ""
					} else {
						buf = ""
					}
				} else {
					buf += string(r)
				}
			case N1:
				if r == ',' {
					if len(n1Buf) == 0 || len(n1Buf) > 3 {
						n1Buf = ""
						state = OPERATION
						continue
					}
					state = N2
				} else if unicode.IsNumber(r) {
					n1Buf += string(r)
				} else {
					n1Buf = ""
					state = OPERATION
				}
			case N2:
				if r == ')' {
					if len(n2Buf) == 0 || len(n2Buf) > 3 {
						n1Buf = ""
						n2Buf = ""
						state = OPERATION
						continue
					}
					n1, err := strconv.Atoi(n1Buf)
					if err != nil {
						log.Fatal(err)
					}
					n2, err := strconv.Atoi(n2Buf)
					if err != nil {
						log.Fatal(err)
					}
					result += n1 * n2
					n1Buf = ""
					n2Buf = ""
					state = OPERATION
				} else if unicode.IsNumber(r) {
					n2Buf += string(r)
				} else {
					n1Buf = ""
					n2Buf = ""
					state = OPERATION
				}
			}
		}
	}
	if err = sc.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
