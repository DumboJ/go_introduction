package main

import (
	"fmt"
	"strings"
)

func main() {
	//create a tic_tac_toe board 创建一个“井”字格
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	//模拟玩家
	board[0][0] = "X"
	board[2][2] = "O"
	board[2][0] = "X"
	board[0][2] = "0"
	board[1][1] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s \n", strings.Join(board[i], ""))
	}
}
