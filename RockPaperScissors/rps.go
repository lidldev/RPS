package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var player string
	var computer string
	for {
		player = getUserChoice(player)
		fmt.Print("Your Choice: ")
		fmt.Println(showChoice(player))

		computer = getComputerChoice(computer)
		fmt.Print("Computers Choice: ")
		fmt.Println(showChoice(computer))

		chooseWinner(player, computer)
		fmt.Println("If You Want Exit After The Result Press Q If You Still Want To Play Just Press Enter")
		var char string
		scanner.Scan()
		char = scanner.Text()
		if char == "q" || char == "Q" {
			break
		}
	}
}
func getUserChoice(player string) string {
	fmt.Println("Rock Paper Scissors Game!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Choose one of the following")
		fmt.Println("'r' for rock")
		fmt.Println("'p' for paper")
		fmt.Println("'s' for scissors")
		scanner.Scan()
		player = scanner.Text()
		if player != "r" && player != "p" && player != "s" {
			fmt.Println("Invalid choice. Please choose 'r', 'p', 's', or 'q'")
			break
		} else {
			return player
		}
	}
	return player
}
func getComputerChoice(computer string) string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(3) + 1
	switch num {
	case 1:
		return "r"
	case 2:
		return "p"
	case 3:
		return "s"
	default:
		break
	}
	return computer
}
func showChoice(choice string) string {
	switch choice {
	case "r":
		fmt.Println("Rock")
		break
	case "p":
		fmt.Println("Paper")
		break
	case "s":
		fmt.Println("Scissors")
		break
	}
	return choice
}
func chooseWinner(player, computer string) {
	switch player {
	case "r":
		if computer == "r" { //rock
			fmt.Println("*******************Tie!!!!*******************")
		} else if computer == "p" { //paper
			fmt.Println("*******************Computer Wins!!!!*******************")
		} else {
			fmt.Println("*******************You Win!!!!*******************")
		}
		break
	case "p":
		if computer == "p" { //rock
			fmt.Println("*******************Tie!!!!*******************")
		} else if computer == "s" { //paper
			fmt.Println("*******************Computer Wins!!!!*******************")
		} else {
			fmt.Println("*******************You Win!!!!*******************")
		}
		break
	case "s":
		if computer == "s" { //rock
			fmt.Println("*******************Tie!!!!*******************")
		} else if computer == "r" { //paper
			fmt.Println("*******************Computer Wins!!!!*******************")
		} else {
			fmt.Println("*******************You Win!!!!*******************")
		}
		break
	}
}
