package main

import (
    "fmt"
    "io/ioutil"
    "strings"
//    "regexp"
)

// return number based on presented string
func getNum(s string) int {
    switch s {
    case "0":
        return 0
    case "1", "one":
        return 1
    case "2", "two":
        return 2
    case "3", "three":
        return 3
    case "4", "four":
        return 4
    case "5", "five":
        return 5
    case "6", "six":
        return 6
    case "7", "seven":
        return 7
    case "8", "eight":
        return 8
    case "9", "nine":
        return 9
    default:
        return 0
    }
    return 0
}


func FindOverlappingMatches(pattern, text string) []string {
    var matches []string
    textLength := len(text)
    patterns := strings.Split(pattern, "|")

    for i := 0; i < textLength; i++ {
        for _, p := range patterns {
            patternLength := len(p)
            if i+patternLength <= textLength && text[i:i+patternLength] == p {
                matches = append(matches, text[i:i+patternLength])
            }
        }
    }

    return matches
}


func main() {
    // bytesRead, _ := ioutil.ReadFile("input_test.txt")
    bytesRead, _ := ioutil.ReadFile("input.txt")

    fileContent := string(bytesRead)

    lines := strings.Split(fileContent, "\n")

    // numbers := make([]int, len(lines))
    sum := 0

    for _, l := range lines {
        foundNum := false
        firstDigit := 0
        lastDigit := 0
        // for _, s := range l {
        //     f, n := isNum(s)
        //     if f {
        //         if !foundNum {
        //             firstDigit = n
        //             foundNum = true
        //         }
        //         lastDigit = n
        //     }
        // }
        // sum += firstDigit * 10 + lastDigit
        // numbers[i] = firstDigit * 10 + lastDigit
        // fmt.Printf("%d%d\n", firstDigit, lastDigit)

        // first, get all matching string parts
        // re := regexp.MustCompile("[0-9]|one|two|three|four|five|six|seven|eight|nine")
        matches := FindOverlappingMatches("1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine", l)
        fmt.Println()
        fmt.Println(l)
        fmt.Println(matches)
        // fmt.Println(re.FindAllString(l, -1))
        for _, num := range matches {
            n := getNum(num) // guarenteed to be proper number
            if n <= 0 {
                continue
            }
            // fmt.Printf("%d\n", getNum(n))
            if !foundNum {
                firstDigit = n
                foundNum = true
            }
            lastDigit = n
        }
        fmt.Println("Final: ", firstDigit * 10 + lastDigit)
        sum += firstDigit * 10 + lastDigit

    }
    fmt.Printf("\nTotal: %d", sum)
}
