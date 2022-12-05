package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)
func Intersection(a, b []int) (c []int) {
    m := make(map[int]bool)

    for _, item := range a {
        m[item] = true
    }

    for _, item := range b {
        if _, ok := m[item]; ok {
            c = append(c, item)
        }
    }
    return
}

func main() {
    count_contains := 0
    file, error := os.Open("real_input.txt")

    if error != nil {
        fmt.Println(error)
    }

    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    
    var elves []string
    for scanner.Scan() {
        elves = append(elves, scanner.Text())
    }
    
    for _, pair_elves := range elves {
        ranges := strings.Split(pair_elves, ",")
        
        ranges_elf1 := strings.Split(ranges[0], "-")
        ranges_elf2 := strings.Split(ranges[1], "-")
        
        start_elf1 := 0
        end_elf1 := 0
        start_elf2 := 0
        end_elf2 := 0
        for i, value := range ranges_elf1{
            if (i == 0){
               start_elf1, _ = strconv.Atoi(value)
            } else {
                end_elf1, _ = strconv.Atoi(value)
            }
        }

        for i, value := range ranges_elf2{
            if (i == 0){
               start_elf2, _ = strconv.Atoi(value)
            } else {
                end_elf2, _ = strconv.Atoi(value)
            }
        }
        
        var elf1 []int
        var elf2 []int
        for start := start_elf1; start <= end_elf1; start++ {
            elf1 = append(elf1, start)
        }
        for start := start_elf2; start <= end_elf2; start++ {
            elf2 = append(elf2, start)
        }
        
        intersection := Intersection(elf1, elf2)
        if(len(intersection) > 0){
            count_contains++
        }
    }
    fmt.Println(count_contains)
}
