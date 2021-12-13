package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
)


func main() {
    var lines []string

    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    


