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


    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    m       := make(map[int]int)
    count   := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "")
        count += 1
        for ind, val := range(s) {
            int_val, _ := strconv.Atoi(val)
            m[ind] = m[ind] + int_val
        
        }
    } 

    var gamma, epsilon []string
    var gamma_string, epsilon_string string

    for i := 0; i < len(m); i ++ {
        fmt.Println(m[i])
        if m[i] > count / 2 {
            gamma   = append(gamma, "1")
            epsilon = append(epsilon, "0")

        } else {
            gamma   = append(gamma, "0")
            epsilon = append(epsilon, "1")
        }
    }
    for _, elem := range(gamma) {
        gamma_string = gamma_string + elem
    }
    for _, elem := range(epsilon) {
        epsilon_string = epsilon_string + elem
    }
    gamma_rate, _   := strconv.ParseInt(gamma_string, 2, 64)
    epsilon_rate, _ := strconv.ParseInt(epsilon_string, 2, 64)
    fmt.Println(gamma_rate * epsilon_rate)
}
