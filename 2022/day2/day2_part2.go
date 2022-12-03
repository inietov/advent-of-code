package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    score := 0
    var election byte

    res_points := map[byte]int{
        'Z': 6,
        'Y': 3,
        'X': 0,
    }

    election_points := map[byte]int{
        'A': 1,
        'B': 2,
        'C': 3,
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
        expected_result := scanner.Text()[2]
        
        if (expected_result == 'X'){
            if (elf == 'A'){
                election = 'C'
            }
            if (elf == 'B'){
                election = 'A'
            }
            if (elf == 'C'){
                election = 'B'
            }

            score += res_points['X'] + election_points[election]
        }

        if (expected_result == 'Y'){
            score += res_points['Y'] + election_points[elf]
        }

        if (expected_result == 'Z'){
            if (elf == 'A'){
                election = 'B'
            }
            if (elf == 'B'){
                election = 'C'
            }
            if (elf == 'C'){
                election = 'A'
            }

            score += res_points['Z'] + election_points[election]
        }

    }

    fmt.Println(score)

}
