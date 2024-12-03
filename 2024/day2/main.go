package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var fileName = "input"

func main() {
	safeReports := countSafeReports()
	fmt.Printf("safe reports: %d\n", safeReports)
}

func countSafeReports() int {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	safeCount := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		strReport := strings.Fields(sc.Text())
		var report []int
		for _, v := range strReport {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			report = append(report, n)
		}

		if checkReport(report) {
			safeCount++
		}
	}
	if err = sc.Err(); err != nil {
		log.Fatal(err)
	}

	return safeCount
}

const (
	undefined = iota
	increasing
	decreasing
)

func checkReport(report []int) bool {
	state := undefined
	for i := 1; i < len(report); i++ {
		prev := report[i-1]
		cur := report[i]

		if state == undefined {
			if prev < cur {
				state = increasing
			} else if prev > cur {
				state = decreasing
			} else {
				return false
			}
		}

		if state == increasing && prev > cur {
			return false
		}

		if state == decreasing && prev < cur {
			return false
		}

		diff := math.Abs(float64(prev - cur))
		if diff == 0 || diff > 3 {
			return false
		}
	}
	return true
}
