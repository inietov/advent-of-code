package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "unicode"
)

func main() {
    priority_sum := 0
    
    file, error := os.Open("real_input.txt")
    if error != nil {
        fmt.Println(error)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    for (scanner.Scan()){
        rucksack := scanner.Text()
        rs_size := len(rucksack)

        rucksack1 := rucksack[0: rs_size/2]
        rucksack2 := rucksack[rs_size/2: rs_size]

        for _, value := range rucksack1 {
            if (strings.IndexRune(rucksack2, value) != -1){
                if (unicode.IsLower(value)){
                    value_to_sum := value-96
                    priority_sum += int(value_to_sum)
                } else {
                    value_to_sum := value-38
                    priority_sum += int(value_to_sum)
                }
                break
            }
        }
    }
    fmt.Println(priority_sum)
}
