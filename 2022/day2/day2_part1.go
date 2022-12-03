package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    score := 0
    result := ""

    res_points := map[string]int{
        "win":  6,
        "tie":  3,
        "lose": 0,
    }

    election_points := map[byte]int{
        'X': 1,
        'Y': 2,
        'Z': 3,
    }
    
    file, error := os.Open("real_input.txt")
    if error != nil {
        fmt.Println(error)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    for (scanner.Scan()){
        elf := scanner.Text()[0]
        strategy := scanner.Text()[2]
        
        if ((elf == 'A' && strategy == 'Y') || (elf == 'B' && strategy == 'Z') || (elf == 'C' && strategy == 'X') ){
            result = "win"
        }

        if ((elf == 'A' && strategy == 'X') || (elf == 'B' && strategy == 'Y') || (elf == 'C' && strategy == 'Z') ){
            result = "tie"
        }
        
        if ((elf == 'A' && strategy == 'Z') || (elf == 'B' && strategy == 'X') || (elf == 'C' && strategy == 'Y') ){
            result = "lose"
        }

        score += res_points[result] + election_points[strategy]
    }

    fmt.Println(score)

}
