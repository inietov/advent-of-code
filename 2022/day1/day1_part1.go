package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "sort"
)

func main() {
    file, error := os.Open("real_input.txt")

    if error != nil {
        fmt.Println(error)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    calories := make([]int, 0)
    addition := 0
    for (scanner.Scan()){
        number, error := strconv.Atoi(scanner.Text())
        addition = addition + number

        if error != nil {
            calories = append(calories, addition)
            addition = 0
        }
    }

    // fmt.Println(calories)
    sort.Ints(calories)
    // fmt.Println(calories)
    fmt.Println("Max Calories for the input: ", calories[len(calories)-1])
}
