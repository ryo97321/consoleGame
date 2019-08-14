package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

func checkReplayFlag(input string) bool {
	validInput := regexp.MustCompile(`^[01]$`)
	if validInput.MatchString(input) {
		return true
	}
	return false
}

func calcSumPips(nDice int) int {
	var sum = 0
	var pip = 0

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < nDice; i++ {
		pip = rand.Intn(3) + 1
		sum += pip
	}
	return sum
}

func showStatus(presentPlace int) {
	for i := 0; i < presentPlace-1; i++ {
		fmt.Print("_")
	}
	fmt.Print("▲ ")
	for i := 0; i < 30-presentPlace; i++ {
		fmt.Print("_")
	}
	fmt.Print("|end")
	fmt.Println()
}

func main() {

	var input string
	var nDice = 0
	var sumPips = 0
	var presentPlace int
	var goalPlace = 30
	var left int
	var gameCount int

	var replayFlag = 1
	for replayFlag == 1 {

		// reset Param
		presentPlace = 1
		left = 30
		gameCount = 0

		fmt.Println("\n▲ Sugoroku ▼")
		fmt.Println("Goal is at 30")
		fmt.Println("You can select number of dice (pattern : 1,2,3)")
		fmt.Println("If you at Goal just -> end!")
		fmt.Println("If dice's sum over, you go over (ex. 31 -> 29)")
		fmt.Println("-----------------------------------------------")
		fmt.Println()
		showStatus(presentPlace)

		stdin := bufio.NewScanner(os.Stdin)

		// first Game
		gameCount++

		fmt.Printf("\nLeft : %d. Number of dice : ", goalPlace)
		stdin.Scan()
		input = stdin.Text()
		nDice, _ = strconv.Atoi(input)

		sumPips = calcSumPips(nDice)
		fmt.Printf("Sum pips : %d", sumPips)

		presentPlace = sumPips
		fmt.Println()
		showStatus(presentPlace)

		left -= presentPlace

		// After first Game
		for left != 0 {
			gameCount++
			fmt.Printf("\nLeft : %d. Number of dice : ", left)
			stdin.Scan()
			input = stdin.Text()
			nDice, _ = strconv.Atoi(input)

			sumPips = calcSumPips(nDice)
			fmt.Printf("Sum pips : %d", sumPips)

			presentPlace += sumPips

			if presentPlace > 30 {
				presentPlace = 30 - (presentPlace - 30)
			}
			fmt.Println()
			showStatus(presentPlace)

			left = 30 - presentPlace
		}

		fmt.Printf("\nCongratulation! GameCount is %d\n", gameCount)

		for {
			fmt.Print("Replay? (1=Yes; 0=No) : ")
			stdin.Scan()
			input = stdin.Text()
			if checkReplayFlag(input) {
				replayFlag, _ = strconv.Atoi(input)
				break
			} else {
				fmt.Println("Invalid input (0 or 1)")
			}
		}
	}
}
