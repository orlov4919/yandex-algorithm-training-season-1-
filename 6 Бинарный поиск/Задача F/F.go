package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cleanInput(input string) string {
	return strings.TrimSpace(input)
}

func strToIntSlice(input, sep string) []int {
	splitedInput := strings.Split(cleanInput(input), sep)
	intSlice := make([]int, len(splitedInput))

	for ind, el := range splitedInput {
		intEl, _ := strconv.Atoi(el)
		intSlice[ind] = intEl
	}

	return intSlice
}

func isGoodtime(currentTime, n, timeX, timeY int) bool {
	return currentTime/timeX+currentTime/timeY >= n
}

func binSearch(minimumTime, maximumTime, n, timeX, timeY int) int {

	for maximumTime-minimumTime > 1 {
		midTime := (maximumTime + minimumTime) / 2

		if isGoodtime(midTime, n, timeX, timeY) {
			maximumTime = midTime
		} else {
			minimumTime = midTime + 1
		}
	}

	if isGoodtime(minimumTime, n, timeX, timeY) {
		return minimumTime
	}

	return maximumTime
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputStr, _ := reader.ReadString('\n')
	inputIntSlice := strToIntSlice(inputStr, " ")
	n, timeX, timeY := inputIntSlice[0], inputIntSlice[1], inputIntSlice[2]

	fmt.Println(min(timeX, timeY) + binSearch(0, (n-1)*(max(timeX, timeY)), n-1, timeX, timeY))
}
