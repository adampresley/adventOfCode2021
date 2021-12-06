package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type BoardSpot struct {
	Value  int
	Marked bool
}

type Board [][]*BoardSpot
type BoardCollection []Board

func main() {
	/*****************************************************************************
	 * Part 1
	 ****************************************************************************/
	numbers, boards := readInput()
	boardWinner := -1

	for index, board := range boards {
		printBoard(index, board)
	}

	fmt.Printf("\nPart 1:\n\n")

	for _, number := range numbers {
		boardWinner = markBoards(boards, number, []int{})

		if boardWinner > -1 {
			fmt.Printf("We have a winner!!\n")
			fmt.Printf("------------------------\n")

			winningSumPart1 := calculateScore(sumUnmarked(boards[boardWinner]), number)
			fmt.Printf("The winning number is %d!\n", winningSumPart1)
			break
		}
	}

	/*****************************************************************************
	 * Part 2
	 ****************************************************************************/
	boardWinner = -1
	winningSumPart2 := 0

	resetBoards(boards)

	fmt.Printf("\nPart 2:\n\n")

	for len(boards) > 0 {
		for _, number := range numbers {
			boardWinner = markBoards(boards, number, []int{})

			if boardWinner > -1 {
				winningSumPart2 = calculateScore(sumUnmarked(boards[boardWinner]), number)

				boards = append(boards[:boardWinner], boards[boardWinner+1:]...)
				boardWinner = -1
				break
			}
		}
	}

	fmt.Printf("We have a winner!!\n")
	fmt.Printf("------------------------\n")

	fmt.Printf("The winning number is %d!\n", winningSumPart2)
}

func readInput() (numbers []int, boards BoardCollection) {
	rawBytes, _ := ioutil.ReadFile("../../inputs/day4.txt")
	rawString := string(rawBytes)

	// Read the top line for numbers
	firstSplit := strings.SplitN(rawString, "\n", 2)
	numbers = stringArrayToInts(strings.Split(firstSplit[0], ","))

	// Read the boards
	boardsSplit := strings.Split(strings.TrimSpace(firstSplit[1]), "\n\n")
	boards = make(BoardCollection, len(boardsSplit))

	for boardIndex, board := range boardsSplit {
		board = strings.TrimSpace(board)

		newBoard := make(Board, 5)
		boardLines := strings.Split(board, "\n")

		for lineIndex, line := range boardLines {
			nums := stringArrayToInts(strings.Split(line, " "))
			newLine := make([]*BoardSpot, 5)

			for numIndex, num := range nums {
				newLine[numIndex] = &BoardSpot{
					Value:  num,
					Marked: false,
				}
			}

			newBoard[lineIndex] = newLine
		}

		boards[boardIndex] = newBoard
	}

	return
}

func markBoards(boards BoardCollection, number int, exclude []int) (boardWinner int) {
	boardWinner = -1

	for boardIndex, board := range boards {
		excludeBoard := false

		for _, b := range exclude {
			if boardIndex == b {
				excludeBoard = true
				break
			}
		}

		if excludeBoard {
			fmt.Printf("Skipping board %d\n", boardIndex)
			continue
		}

		// fmt.Printf("marking board with %d, (%d)\n", number, boardIndex)
		for _, boardLine := range board {
			for _, spot := range boardLine {
				if spot.Value == number {
					spot.Marked = true

					if boardWinner == -1 && checkBoardForWin(board) {
						boardWinner = boardIndex
						return
					}
				}
			}
		}
	}

	return
}

func checkBoardForWin(board Board) bool {
	for pos := 0; pos < 5; pos++ {
		count := 0

		// Vertical
		for y := 0; y < 5; y++ {
			if board[y][pos].Marked {
				count++
			}
		}

		if count == 5 {
			return true
		}

		count = 0

		// Horizontal
		for x := 0; x < 5; x++ {
			if board[pos][x].Marked {
				count++
			}
		}

		if count == 5 {
			return true
		}
	}

	return false
}

func sumUnmarked(board Board) (unmarkedSum int) {
	for _, boardLine := range board {
		for _, spot := range boardLine {
			if !spot.Marked {
				unmarkedSum += spot.Value
			}
		}
	}

	return
}

func resetBoards(boards BoardCollection) {
	for _, board := range boards {
		for _, line := range board {
			for _, spot := range line {
				spot.Marked = false
			}
		}
	}
}

func calculateScore(unmarkedSum, winningNumber int) (sum int) {
	return unmarkedSum * winningNumber
}

func stringArrayToInts(input []string) (result []int) {
	for _, s := range input {
		if s != "" {
			i, _ := strconv.Atoi(s)
			result = append(result, i)
		}
	}

	return
}

func printBoard(index int, board Board) {
	fmt.Printf("Board #%d\n-----------------------------------\n", index)

	for _, line := range board {
		for _, bs := range line {
			if bs.Marked {
				fmt.Printf("*%d ", bs.Value)
			} else {
				fmt.Printf("%d ", bs.Value)
			}
		}

		fmt.Printf("\n")
	}

	fmt.Printf("\n")
}
