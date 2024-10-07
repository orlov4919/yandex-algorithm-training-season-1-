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

func isGoodRectangle(w, h, a, b, n int) bool {
	return (w/a)*(h/b) >= n || (w/b)*(h/a) >= n
}

func binSearch(l, r, n, a, b, w, h int) int {

	for r-l > 1 {
		mid := (r + l) / 2
		if isGoodRectangle(w, h, a+2*mid, b+2*mid, n) {
			l = mid
		} else {
			r = mid - 1
		}
	}

	if isGoodRectangle(w, h, a+2*r, b+2*r, n) {
		return r
	}

	return l
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	inputData := strToIntSlice(input, " ")
	n, a, b, w, h := inputData[0], inputData[1], inputData[2], inputData[3], inputData[4]

	fmt.Println(binSearch(0, max(w, h)/2, n, a, b, w, h))
}
