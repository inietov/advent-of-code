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
    var rucksacks []string
    file, error := os.Open("real_input.txt")
    if error != nil {
        fmt.Println(error)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    for (scanner.Scan()){
        rucksacks = append(rucksacks, scanner.Text())
    }

    aux := 0
    for range rucksacks {       
        if (aux < len(rucksacks)){
            team := rucksacks[aux:aux+3]
            
            for _, value := range team[0] {
                if (strings.IndexRune(team[1], value) != -1){
                    if (strings.IndexRune(team[2], value) != -1){
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
            aux += 3
        }
    }

    fmt.Println(priority_sum)
}
