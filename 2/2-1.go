package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strings"
    "strconv"
)


func main() {

    x   := 0
    y   := 0
    aim := 0

    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        fmt.Println(x, y)
		s := strings.Split(scanner.Text(), " ")
        fmt.Println(s)
        dir := s[0]
        val, _ := strconv.Atoi(s[1])
        switch {
        case dir == "up":       y -= val
        case dir == "down":     y += val
        case dir == "forward":  x += val
        } 
    }
    fmt.Println(x * y)

}