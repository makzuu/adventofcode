package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const fileName = "input"

func main() {
	leftList, rightList := getLists()
	slices.Sort(leftList)
	slices.Sort(rightList)

	totalDist := getTotalDist(leftList, rightList)
	fmt.Println(totalDist)
}

func getLists() ([]int, []int) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var leftList, rightList []int
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		elements := strings.Fields(line)
		leftElement, err := strconv.Atoi(elements[0])
		if err != nil {
			log.Fatal(err)
		}
		rightElement, err := strconv.Atoi(elements[1])
		if err != nil {
			log.Fatal(err)
		}

		leftList = append(leftList, leftElement)
		rightList = append(rightList, rightElement)
	}
	if err = sc.Err(); err != nil {
		log.Fatal(err)
	}

	return leftList, rightList
}

func getTotalDist(ll, rl []int) int {
	totalDist := 0
	for i := 0; i < len(ll); i++ {
		totalDist += int(math.Abs(float64(ll[i] - rl[i])))
	}
	return totalDist
}