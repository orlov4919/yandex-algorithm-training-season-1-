package main

import (
	"bufio"
	"fmt"
	"math"
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

func isGoodSquare(squareSide, w, h, n int) bool {
	return (squareSide/w)*(squareSide/h) >= n
}

func binSearch(l, r, w, h, n int) int {
	for r-l > 1 {
		mid := (r + l) / 2

		if isGoodSquare(mid, w, h, n) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if isGoodSquare(l, w, h, n) {
		return l
	}
	return r
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	taskData := strToIntSlice(input, " ")
	w, h, n := taskData[0], taskData[1], taskData[2]
	fmt.Println(binSearch(1, max(w, h)*int(math.Sqrt(float64(n))+1), w, h, n))
}
