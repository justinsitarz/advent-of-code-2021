package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strconv"
)

func main() {

    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var lines []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		lines = append(lines, num)	
	}

    length  := len(lines) - 2
    counter := 0

    tmp := lines[0] + lines[1] + lines[2]

	for i := 1; i < length; i++ {
        val := lines[i] + lines[i + 1] + lines[i + 2]
        fmt.Println(tmp, val)
        if val > tmp {
            counter += 1
        }
        tmp = val
    }
    fmt.Println(counter)
}