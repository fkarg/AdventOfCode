package main

import (
    "io/ioutil"
    "strconv"
    "strings"
    "fmt"
    "os"
)

func getMinPower(input string) []int {
    lines := strings.Split(input, "\n")
    listRequired := []int{}

    for _, line := range lines {
        if line == "" {
            continue
        }
        parts := strings.Split(line, ": ")
        subsets := strings.Split(parts[1], "; ")

        red, green, blue := 0, 0, 0

        for _, subset := range subsets {
            cubes := strings.Split(subset, ", ")
            for _, cube := range cubes {
                countAndColor := strings.Split(cube, " ")
                count, _ := strconv.Atoi(countAndColor[0])
                color := countAndColor[1]

                switch color {
                case "red":
                    if count > red {
                        red = count
                    }
                case "green":
                    if count > green {
                        green = count
                    }
                case "blue":
                    if count > blue {
                        blue = count
                    }
                }

            }
        }
        listRequired = append(listRequired, red * green * blue)
    }

    return listRequired
}


func main() {
    var filename string
    fmt.Printf("Len: [%d]\n", len(os.Args))
    if len(os.Args) <= 1 {
        filename = "input.txt"
        fmt.Println("Using default [input.txt]")
    } else {
        filename = os.Args[1]
        fmt.Printf("Using custom [%s]\n", filename)
    }
    bytesRead, _ := ioutil.ReadFile(filename)
    input := string(bytesRead)
    // fmt.Println(input)

    possibleGames := getMinPower(input)

    fmt.Println("Power of games:", possibleGames)
    sum := 0
    for _, id := range possibleGames {
        sum += id
    }
    fmt.Println("Sum of minimum power:", sum)
}
