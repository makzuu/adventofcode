package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
    "unicode"
    "strconv"
)

func getLines(filename string) []string {
    var lines []string

    f, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    sc := bufio.NewScanner(f)
    for sc.Scan() {
        lines = append(lines, sc.Text())
    }
    f.Close()

    return lines
}

func getSumPartNumbers(lines []string) int{
    var sum int
    var buffer string

    for i, line := range lines {
        partNumber := false
        for j, char := range line {
            if unicode.IsDigit(char) {
                buffer += string(char)

                for ii := i - 1; ii <= i + 1; ii++ {
                    for jj := j - 1; jj <= j + 1; jj++ {
                        if ii < 0 || ii >= len(lines) || jj < 0 || jj >= len(line) {
                            continue
                        }
                        if !unicode.IsDigit(rune(lines[ii][jj])) && rune(lines[ii][jj]) != '.' {
                            partNumber = true
                        }
                    }
                }
            } else {
                if partNumber {
                    n, _ := strconv.Atoi(buffer)
                    sum += n
                }
                buffer = ""
                partNumber = false
            }
        }
        if buffer != "" && partNumber {
            n, _ := strconv.Atoi(buffer)
            sum += n
        }
    }

    return sum
}

func main() {
    lines := getLines("input")
    sum := getSumPartNumbers(lines)
    fmt.Println("=", sum)
}
