package main

import (
    "fmt"
    "log"
    "os"
    "bufio"
    "strings"
    "strconv"
)

const ( 
    redCubesMax = iota + 12
    greenCubesMax
    blueCubesMax
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

func getSets(game string) []string {
    return strings.Split(game, ";")
}

func validateSet(set []string) bool {
    var redCount, greenCount, blueCount int

    for _, pair := range set {
        p := strings.Fields(strings.TrimSpace(pair))
        n, _ := strconv.Atoi(p[0])
        color := p[1]

        switch color {
        case "red":
            redCount += n
        case "green":
            greenCount += n
        case "blue":
            blueCount += n
        }
    }

    if redCount > redCubesMax || greenCount > greenCubesMax || blueCount > blueCubesMax {
        return false
    }

    return true
}

func isValidGame(game string) bool {
    sets := getSets(game)
    for _, set := range sets {
        s := strings.Split(set, ",")
        if !validateSet(s) {
            return false
        }
    }

    return true
}

func main() {
    var sumPossibleGameIds int

    lines := getLines("input")
    for i, game := range lines {
        gameId := i + 1
        if !isValidGame(strings.Split(game, ":")[1]) {
            continue
        }
        sumPossibleGameIds += gameId
    }

    fmt.Println("=", sumPossibleGameIds)
}
