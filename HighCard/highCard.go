package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func showCard(card string) {
	fmt.Printf("|%s|\n", card)
}

func showInitialMessage() {
	fmt.Println()
	fmt.Println("First you will get one card. Please decide the latch")
	fmt.Println("If the next card is the same or larger than the previous card, you win.")
	fmt.Println("The latch woll return, so decide if you want to continue.")
	fmt.Println("The win will double if you win continuously.")
	fmt.Println("After that, the amount of money that returns 4 times and 8 times will increase.")
	fmt.Println("However, if you lose, you won't win.")
	fmt.Println("The game ends when you lose your money or go bankrupt.")
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("Game start. Your money is 100$.")
	fmt.Println()
}

func main() {

	var input string

	var previousCardValue int
	var nowCardValue int

	var money = 100
	var bet int
	var rate = 2

	var replayFlag int

	var cards []string
	cards = append(cards, "A")
	cards = append(cards, "2")
	cards = append(cards, "3")
	cards = append(cards, "4")
	cards = append(cards, "5")
	cards = append(cards, "6")
	cards = append(cards, "7")
	cards = append(cards, "8")
	cards = append(cards, "9")
	cards = append(cards, "10")
	cards = append(cards, "J")
	cards = append(cards, "Q")
	cards = append(cards, "K")

	showInitialMessage()

	for {

		rate = 2

		fmt.Println("First card.")
		rand.Seed(time.Now().UnixNano())
		cardIndex := rand.Intn(13)
		showCard(cards[cardIndex])

		previousCardValue = cardIndex

		fmt.Print("How money bets? (1$~100$) : ")
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		input = stdin.Text()
		bet, _ = strconv.Atoi(input)

		cardIndex = rand.Intn(13)
		showCard(cards[cardIndex])

		nowCardValue = cardIndex

		if nowCardValue > previousCardValue {
			fmt.Printf("You win. %d$ win\n", bet)
			fmt.Printf("Rate is %d. Continue? (1=Yes; 0=No) : ", rate)

			stdin.Scan()
			input = stdin.Text()
			replayFlag, _ = strconv.Atoi(input)

			for {
				if replayFlag == 0 {
					money += bet
					break
				}

				cardIndex = rand.Intn(13)
				showCard(cards[cardIndex])

				previousCardValue = nowCardValue
				nowCardValue = cardIndex

				if nowCardValue > previousCardValue {
					bet *= rate
					rate *= 2
					fmt.Printf("You win. %d$ win\n", bet)
					fmt.Printf("Rate is %d. Continue? (1=Yes; 0=No) : ", rate)

					stdin.Scan()
					input = stdin.Text()
					replayFlag, _ = strconv.Atoi(input)
				} else {
					money -= bet
					fmt.Printf("You lose. Your money is %d$.\n", money)
				}
			}
		} else {
			money -= bet
			fmt.Printf("You lose. Your money is %d$.\n", money)
		}

		// money > 1000 or bankrupt -> end
		if money > 1000 || money < 0 {
			break
		}

		fmt.Printf("Money is %d\n", money)
		fmt.Println()
	}

}
