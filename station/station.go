package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func showInitialMessage() {
	fmt.Println("The train runs as 100km/h. Distance to the station is 500m.")
	fmt.Println("Please select brake strength and stop the train at station.")
	fmt.Println("You can select 0 to 2 brake strength every time.")
	fmt.Println("-----------------------------------------------------------")
}

func showTrainPlace(leftDistance int) {
	trainPlace := leftDistance / 10
	for i := 0; i < trainPlace; i++ {
		fmt.Print("_")
	}
	fmt.Println("□ □ □")
}

func calcLeftDistance(newSpeed, leftDistance int) int {
	newSpeedMeterPerSecondFloat64 := float64(newSpeed) * 1000.0 / 3600.0
	newSpeedMeterPerSecond := int(math.Ceil(newSpeedMeterPerSecondFloat64))
	distanceFloat64 := ((float64(newSpeedMeterPerSecond) * 10.0) + 35.0) / 36.0
	distance := int(math.Ceil(distanceFloat64))
	leftDistance -= distance
	return leftDistance
}

func main() {
	var nowSpeed int // km/h
	var newSpeed int // km/h
	var brakeStrength int
	var input string
	var inputInteger int
	var leftDistance int
	var countSecond int
	var replayFlag int
	stdin := bufio.NewScanner(os.Stdin)

	for {
		// reset Param
		nowSpeed = 100
		brakeStrength = 0
		leftDistance = 500
		countSecond = 0

		showInitialMessage()

		for {
			countSecond++

			showTrainPlace(leftDistance)

			fmt.Printf("Left distance %dm | Speed %dkm/h | brake strength %d\n", leftDistance, nowSpeed, brakeStrength)
			fmt.Print("Brake strength? (2=strong, 1=weak, 0=keep) : ")

			stdin = bufio.NewScanner(os.Stdin)
			stdin.Scan()
			input = stdin.Text()
			inputInteger, _ = strconv.Atoi(input)

			if inputInteger == 2 {
				brakeStrength++
			} else if inputInteger == 1 {
				brakeStrength--
			}

			newSpeed = nowSpeed - brakeStrength
			leftDistance = calcLeftDistance(newSpeed, leftDistance)
			nowSpeed = newSpeed

			if leftDistance <= 0 {
				break
			}
		}

		fmt.Printf("%ds Distance to station is %dm.\n", countSecond, leftDistance)

		fmt.Print("Replay? (1=Yes; 0=No) : ")
		stdin.Scan()
		input = stdin.Text()
		replayFlag, _ = strconv.Atoi(input)

		if replayFlag == 0 {
			break
		} else {
			fmt.Println()
		}
	}

}
