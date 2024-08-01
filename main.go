package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/gigga1998/consoleReader/pkg/reader"
)


var GAME_FIELD = [3][3]string  {
	{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "},
}
var PLAYER_FIGURE string


func menu() {
	fmt.Println("Выберите игровую фигуру.")
	fmt.Println("* Введите X, если хотите играть за X.")
	fmt.Println("* Введите O, если хотите играть за O.")
	fmt.Println("* Введите EXIT, если хотите выйти.")
}


func showGameField() {
	fmt.Printf("---|---|---\n")
	fmt.Printf(" %s | %s | %s \n", GAME_FIELD[0][0], GAME_FIELD[0][1], GAME_FIELD[0][2])
	fmt.Printf("---|---|---\n")
	fmt.Printf(" %s | %s | %s \n", GAME_FIELD[1][0], GAME_FIELD[1][1], GAME_FIELD[1][2])
	fmt.Printf("---|---|---\n")
	fmt.Printf(" %s | %s | %s \n", GAME_FIELD[2][0], GAME_FIELD[2][1], GAME_FIELD[2][2])
	fmt.Printf("---|---|---\n")
}


func showGameCoord() {
	fmt.Println("---|---|---")
	fmt.Println(" 1 | 2 | 3 ")
	fmt.Println("---|---|---")
	fmt.Println(" 4 | 5 | 6 ")
	fmt.Println("---|---|---")
	fmt.Println(" 7 | 8 | 9 ")
	fmt.Println("---|---|---")
}


func peekFigure() int8{
	x := consoleReader.ReadString("Ваш ввод: ")
	if strings.TrimRight(x, "\n") == "O"{
		PLAYER_FIGURE = "O"
		return 0
	}else if strings.TrimRight(x, "\n") == "X"{
		PLAYER_FIGURE = "X"
		return 0
	}else if strings.TrimRight(x, "\n") == "EXIT"{
		return 1
	}
	fmt.Println("!!!Команда не найдена!!!")
	return -1
}


func enter_coordinate() (int8, int8) {
	choice := consoleReader.ReadString("Ваш ход. Введите координату(1-9): ")
	choice = strings.TrimSpace(choice)
	if choice == "1" {
		return 0, 0
	} else if choice == "2" {
		return 0, 1
	} else if choice == "3" {
		return 0, 2
	} else if choice == "4" {
		return 1, 0
	} else if choice == "5" {
		return 1, 1
	} else if choice == "6" {
		return 1, 2
	} else if choice == "7" {
		return 2, 0
	} else if choice == "8" {
		return 2, 1
	} else if choice == "9" {
		return 2, 2
	} else if choice == "EXIT" {
		return -1, 0
	}
	return -1, -1
}

func check_the_winner() string{
	var i int8
	for i = 0; i < 3; i++ {
		if GAME_FIELD[i][0] == GAME_FIELD[i][1] && GAME_FIELD[i][1] == GAME_FIELD[i][2] {
			if GAME_FIELD[i][0] == "X" {
				return "X"
			} else if GAME_FIELD[i][0] == "O" {
				return "O"
			}

		} else if GAME_FIELD[0][i] == GAME_FIELD[1][i] && GAME_FIELD[1][i] == GAME_FIELD[2][i] {
			if GAME_FIELD[0][i] == "X" {
				return "X"
			} else if GAME_FIELD[0][i] == "O" {
				return "O"
			}
		}
	}
	if GAME_FIELD[0][0] == GAME_FIELD[1][1] && GAME_FIELD[1][1] == GAME_FIELD[2][2] {
		if GAME_FIELD[0][0] == "X" {
			return "X"
		} else if GAME_FIELD[0][0] == "O" {
			return "O"
		}
	} else if GAME_FIELD[0][2] == GAME_FIELD[1][1] && GAME_FIELD[1][1] == GAME_FIELD[2][0] {
		if GAME_FIELD[2][0] == "X" {
			return "X"
		} else if GAME_FIELD[2][0] == "O" {
			return "O"
		}
	}
	return ""
}


func put_in_table(i int8, j int8) int8{
	if i == -1 && j == -1 {
		return -1
	} else if i == -1 && j == 0 {
		return 1
	}
	GAME_FIELD[i][j] = PLAYER_FIGURE
	return 0
}


func main() {
	menu()
	var err_state int8
	var state = true
	for state{
		err_state = peekFigure()
		if err_state == 0{
			state = false
		}
		if err_state == 1{
			fmt.Println("Пока!")
			os.Exit(0)
		}
	}


	fmt.Println("Выберите клетку из набора координат 1-9:")
	showGameCoord()
	var winner string
	for true {
		var i, j = enter_coordinate()
		var status = put_in_table(i, j)
		if status == 1 {
			os.Exit(0)
		} else if status == -1{
			fmt.Println("Такой команды нет!")
			continue
		} else if status == 0 {
			showGameField()
			winner = check_the_winner()
		}
		if winner == "X" || winner == "O"{
			break
		}
	}
	fmt.Printf("Победил игорк %s\n", winner)
	os.Exit(0)
}
