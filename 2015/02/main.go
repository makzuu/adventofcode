package main

import (
	"os"
	"log"
	"fmt"
	"errors"
	"bufio"
	"io"
	"strings"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer f.Close()

	const (
		l = iota
		w
		h
	)

	r := bufio.NewReader(f)

	sf := 0
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if (errors.Is(err, io.EOF)) {
				break
			} else {
				log.Fatalf("error while reading file: %v", err)
			}
		}
		stringDimensions := strings.Split(line, "x")
		dimentions := make([]int, 3)
		for i, d := range stringDimensions {
			dimentions[i], err = strconv.Atoi(strings.TrimSpace(d))
			if err != nil {
				log.Fatalf("error while converting string to number: %v", err)
			}
		}
		// 2*l*w + 2*w*h + 2*h*l
		smallestSide := dimentions[l] * dimentions[w]

		sf += smallestSide * 2

		bSide := dimentions[w] * dimentions[h]
		if bSide < smallestSide {
			smallestSide = bSide
		}

		sf += bSide * 2
		
		cSide := dimentions[h] * dimentions[l]
		if cSide < smallestSide {
			smallestSide = cSide
		}

		sf += cSide * 2

		sf += smallestSide
	}

	fmt.Printf("square feet of wrapping paper: %v\n", sf)
}
