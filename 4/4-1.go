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


    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var lines []string
    var boards [][][]int
    var board [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
    	lines = append(lines, scanner.Text())       
        }


    for i, l := range lines {
        if i < 2 {
            continue
        }
        if l == "" {
            boards = append(boards, board)
            board = make([][]int, 0)
            continue
        }

        nums := strings.Fields(l)
        n := make([]int, 0)
        for _, number := range nums {
            i, err := strconv.Atoi(number)
            if err != nil {
                log.Fatal(err)
            }
            n = append(n, i)
        }
        board = append(board, n)
    }

    for _, n := range nums {
        markBoards(n, boards)
        if winner, ok := hasWinner(boards); ok {
            fmt.Println("winner score: ", calculateScore(n, winner))
            fmt.Println("winner board: ", winner)
            os.Exit(0)
        }
    }
}

func hasWinner(boards [][][]int) ([][]int, bool) {
    for _, board := range boards {
        for i := 0; i < len(board); i++ {
            // check vertically
            rowWon := true
            for j := 0; j < len(board[i]); j++ {
                if board[i][j] != -1 {
                    rowWon = false
                    break
                }
            }
            if rowWon {
                return board, true
            }
        }

        // these are 5x5 and not varring in size.
        l := len(board[0])
        for col := 0; col < l; col++ {
            colWon := true
            for i := 0; i < len(board); i++ {
                if board[i][col] != -1 {
                    colWon = false
                    break
                }
            }
            if colWon {
                return board, true
            }
        }
    }

    return nil, false
}

//     bingo_calls = strings.Split(lines[0], ",")

//     // Build the boards into bingo_boards

//     var bingo_board [][]string
//     for i := 2; i < len(lines); i ++ {
//         if lines[i] == "" {
//             bingo_boards = append(bingo_boards, bingo_board)
//             bingo_board = nil
            
//         } else {
//             bingo_board = append(bingo_board, strings.Fields(lines[i]))
//             if i == len(lines) - 1 {
//                 bingo_boards = append(bingo_boards, bingo_board)
//             }
//         }

//     }

//     // Loop through calls - for each call, loop through all bingo boards, check and update hits with an 'x'

//     for _, call := range(bingo_calls) {
//         for i, _ := range(bingo_boards) {
//             bingo_boards[i] = checkBoard(bingo_boards[i], call)
//             fmt.Println(bingo_boards)
//         }  
//     }
// }



// func checkBoard(board [][]string, call string)  [][]string {
//     for i, _ := range(board)  {
//         for j, _ := range(board[i]) {
//             if board[i][j] == call {
//                 board[i][j] = "x"
//             }
//         }
//     }
//     if checkForWin(board)

//     return
// }

// func checkForWin(board [][]string)
//     for i, _ := range(board) {
//         for j, _ := range(board) {

//         }

//     }



// // words := strings.Fields(someString)

