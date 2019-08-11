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

// 出題する4桁の数字にダブりがないか確認する関数
func isAnswerNumberDuplicate(answer int) bool {
	answerString := strconv.Itoa(answer)
	var checkList []int
	for i := 0; i < 10; i++ {
		checkList = append(checkList, 0)
	}

	for _, answerNumberString := range answerString {
		answerNumber, _ := strconv.Atoi(string(answerNumberString))
		checkList[answerNumber]++
	}

	for i := 0; i < 10; i++ {
		if checkList[i] > 1 {
			return true
		}
	}

	return false
}

// ユーザの入力が有効な値か確認する関数
func isInputValid(input string) bool {
	var validInput = regexp.MustCompile(`^[0-9]{4}$`)
	if validInput.MatchString(input) {
		return true
	}
	return false
}

// ユーザの入力と正解を比較しEAT数とBITE数を返す関数
func getEatAndBite(input, answer string) (nEat, nBite int) {

	// 1番目の数字がEat -> eatCheckList[0] = true
	var eatCheckList []bool
	for i := 0; i < 4; i++ {
		eatCheckList = append(eatCheckList, false)
	}

	// Eat判定
	for i := 0; i < 4; i++ {
		checkNumber, _ := strconv.Atoi(string(input[i]))
		for j := 0; j < 4; j++ {
			number, _ := strconv.Atoi(string(answer[j]))
			if checkNumber == number && i == j {
				nEat++
				eatCheckList[j] = true
			}
		}
	}

	// Bite判定
	for i := 0; i < 4; i++ {
		if eatCheckList[i] == true {
			continue
		}

		isNumberExist := false // input内の数字がanswer内にあるかどうか（あればtrue）
		checkNumber, _ := strconv.Atoi(string(input[i]))
		for j := 0; j < 4; j++ {
			if eatCheckList[j] == true {
				continue
			}
			number, _ := strconv.Atoi(string(answer[j]))
			if checkNumber == number {
				isNumberExist = true
				break
			}
		}

		if isNumberExist {
			nBite++
		}
	}

	return nEat, nBite

}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Numer0n start.")

	var answer int
	for {
		answer = rand.Intn(9000) + 1000
		if isAnswerNumberDuplicate(answer) != true {
			break
		}
	}

	stdin := bufio.NewScanner(os.Stdin)
	for {
		var input string

		for {
			// ユーザの入力を受け取る
			fmt.Print("Input Number : ")
			stdin.Scan()
			input = stdin.Text()

			// 入力値が正常な値かチェックする
			if isInputValid(input) {
				break
			} else {
				fmt.Println("Input is not valid")
				fmt.Println("Input 4-digits number (ex. 1234)")
				fmt.Println()
			}
		}

		// Eat数とBite数を取得する
		nEat, nBite := getEatAndBite(input, strconv.Itoa(answer))

		// 結果を表示する（Eat, Bite）
		fmt.Printf("%dEAT-%dBITE\n\n", nEat, nBite)

		// 正解したら抜ける
		if input == strconv.Itoa(answer) {
			break
		}
	}

}
