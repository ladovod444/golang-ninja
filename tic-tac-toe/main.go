package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

const (
    empty = " "
    playerX = "X"
    playerO = "O"
)

func main() {
    board := [3][3]string{
        {empty, empty, empty},
        {empty, empty, empty},
        {empty, empty, empty},
    }

    currentPlayer := playerX

    scanner := bufio.NewScanner(os.Stdin)
    for {
        printBoard(board)
        fmt.Printf("Player %s, enter your move (row[1-3] col[1-3]): ", currentPlayer)
        scanner.Scan()
        input := scanner.Text()
        move := strings.Split(input, " ")

        if len(move) != 2 {
            fmt.Println("Invalid input! Please enter row and column as R C.")
            continue
        }

        row, err1 := strconv.Atoi(move[0])
        col, err2 := strconv.Atoi(move[1])

        if err1 != nil || err2 != nil || !isValidMove(row, col, board) {
            fmt.Println("Invalid move! Make sure your row and column are from 1 to 3 and the cell is empty.")
            continue
        }

        board[row-1][col-1] = currentPlayer

        if isWinner(currentPlayer, board) {
            printBoard(board)
            fmt.Printf("Player %s wins!\n", currentPlayer)
            break
        }

        if isBoardFull(board) {
            printBoard(board)
            fmt.Println("It's a tie!")
            break
        }

        currentPlayer = switchPlayer(currentPlayer)
    }
}

func printBoard(board [3][3]string) {
    fmt.Println("  1 2 3")
    for r, row := range board {
        fmt.Print(r+1, " ")
        for c, col := range row {
            if c == 0 {
                fmt.Print(col)
            } else {
                fmt.Printf("|%s", col)
            }
        }
        fmt.Println()
        if r < 2 {
            fmt.Println("  -----")
        }
    }
    fmt.Println()
}

func isValidMove(row int, col int, board [3][3]string) bool {
    return row > 0 && row <= 3 && col > 0 && col <= 3 && board[row-1][col-1] == empty
}

func switchPlayer(current string) string {
    if current == playerX {
        return playerO
    }
    return playerX
}

func isWinner(player string, board [3][3]string) bool {
    // Check rows, columns, and diagonals
    return checkLines(player, board) || checkDiagonals(player, board)
}

func checkLines(player string, board [3][3]string) bool {
    for i := 0; i < 3; i++ {
        if (board[i][0] == player && board[i][1] == player && board[i][2] == player) ||
           (board[0][i] == player && board[1][i] == player && board[2][i] == player) {
               return true
           }
    }
    return false
}

func checkDiagonals(player string, board [3][3]string) bool {
    return (board[0][0] == player && board[1][1] == player && board[2][2] == player) ||
           (board[0][2] == player && board[1][1] == player && board[2][0] == player)
}

func isBoardFull(board [3][3]string) bool {
    for _, row := range board {
        for _, col := range row {
            if col == empty {
                return false
            }
        }
    }
    return true
}