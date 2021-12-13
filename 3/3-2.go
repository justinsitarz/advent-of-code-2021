package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "math"
    "strconv"
)


func main() {
    var o2_rating, co2_rating []string
    var o2_string, co2_string string

    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        o2_rating = append(o2_rating, scanner.Text())
        co2_rating = append(co2_rating, scanner.Text())
    }

    for i := 0; len(o2_rating) > 1; i ++ {
        o2_rating = getMostCommon(o2_rating, i)
    }

    for i := 0; len(co2_rating) > 1; i ++ {
        co2_rating = getLeastCommon(co2_rating, i)
    }

    for _, elem := range(o2_rating) {
        o2_string = o2_string + elem
    }
    for _, elem := range(co2_rating) {
        co2_string = co2_string + elem
    }

    o2, _   := strconv.ParseInt(o2_string, 2, 64)
    co2, _ := strconv.ParseInt(co2_string, 2, 64)
    fmt.Println(o2 * co2)
}



func getMostCommon(nums []string, index int) []string {
    count := len(nums)
    var total int
    var new_nums []string
    for _, num := range(nums) {
        total += int(num[index])
    }
    val := math.Round(float64(total) / float64(count))
    for _, num := range(nums) {
        if int(num[index]) == int(val) {
            new_nums = append(new_nums, num)
        }
    }
    return new_nums
}

func getLeastCommon(nums []string, index int) []string {
    count := len(nums)
    var total int
    var new_nums []string
    for _, num := range(nums) {
        total += int(num[index])
    }
    val := math.Round(float64(total) / float64(count))
    val = math.Abs(val - 97)
    for _, num := range(nums) {
        if int(num[index]) == int(val) {
            new_nums = append(new_nums, num)
        }
    }
    return new_nums

}



