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

func isInputValid(input string) bool {
	validInput := regexp.MustCompile(`^[0-9]+$`)
	if validInput.MatchString(input) != true {
		return false
	}

	number, _ := strconv.Atoi(input)
	if number >= 0 && number <= 100 {
		return true
	}
	return false
}

func main() {

	var count int
	var radarPoint int
	var radarRange int
	var radarRangeMin int
	var radarRangeMax int
	var input string

	fmt.Println()
	fmt.Println("------------------")
	fmt.Println("--- radar Game ---")
	fmt.Println("------------------")
	fmt.Println()

	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(101)

	stdin := bufio.NewScanner(os.Stdin)

	for {
		count++
		fmt.Printf("[%d]\n", count)

		fmt.Print("Input radar point : ")
		stdin.Scan()
		input = stdin.Text()

		// check input
		if isInputValid(input) == false {
			fmt.Println("invalid input.")
			fmt.Println()
			count--
			continue
		}
		radarPoint, _ = strconv.Atoi(input)

		fmt.Print("Input radar range : ")
		stdin.Scan()
		input = stdin.Text()

		// check input
		if isInputValid(input) == false {
			fmt.Println("invalid input.")
			fmt.Println()
			count--
			continue
		}
		radarRange, _ = strconv.Atoi(input)

		// set radarRangeMin
		if radarPoint-radarRange < 0 {
			radarRangeMin = 0
		} else {
			radarRangeMin = radarPoint - radarRange
		}

		// set radarRangeMax
		if radarPoint+radarRange > 100 {
			radarRangeMax = 100
		} else {
			radarRangeMax = radarPoint + radarRange
		}

		if target == radarPoint {
			fmt.Println()
			fmt.Printf("target [%d] found.\n", target)
			fmt.Printf("Try count : %d\n", count)
			break
		}

		if target >= radarRangeMin && target <= radarRangeMax {
			fmt.Println("Hit.")
		} else {
			fmt.Println("Miss.")
		}
		fmt.Println()
	}
}
