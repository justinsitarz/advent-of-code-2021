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
 	count 	:= 0
   	tmp 	:= 199
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		num, _ := strconv.Atoi(lineStr)
		if num > tmp {
			count += 1
		}
		tmp = num		
	}
	fmt.Println(count)
}