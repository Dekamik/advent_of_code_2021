package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
)

func ReadInts(r io.Reader) ([]int, error) {
    scanner := bufio.NewScanner(r)
    scanner.Split(bufio.ScanWords)
    var result []int
    for scanner.Scan() {
        x, err := strconv.Atoi(scanner.Text())
        if err != nil {
            return result, err
        }
        result = append(result, x)
    }
    return result, scanner.Err()
}

func main() {
    if os.Args[1] == "1" {
        file := os.Args[2]
        reader, err := os.Open(file)
        if err != nil {
            os.Exit(1)
        }
        ints, err := ReadInts(reader)

        var increases = 0
        for i, integer := range ints {
            if i == 0 {
                continue
            }
            if integer > ints[i - 1] {
                increases++
            }
        }

        fmt.Println(increases)
        os.Exit(0)
    } else if os.Args[1] == "2" {
        
    }
}
