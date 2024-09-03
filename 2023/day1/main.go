package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "unicode"
    "strconv"
)

func readLines(sc *bufio.Scanner) (lines []string) {
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return
}

func getNumbers(lines []string) (numbers [][]rune) {
	for i, line := range lines {
        numbers = append(numbers, []rune{})
		for _, char := range line {
			if unicode.IsNumber(char) {
				numbers[i] = append(numbers[i], char)
			}
		}
	}
	return
}

func getSum(numbers [][]rune) int {
    sum := 0
    for _, numsInLine := range numbers {
        num, _ := strconv.Atoi(string(numsInLine[0]) + string(numsInLine[len(numsInLine) - 1]))
        sum += num
    }
    return sum
}

func main() {
	fileName := "input"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(f)
	lines := readLines(sc)
	f.Close()

    nums := getNumbers(lines)
    sum := getSum(nums)
    fmt.Println(sum)
}
