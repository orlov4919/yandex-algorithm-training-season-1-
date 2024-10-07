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

func isMarkGood(sum, n, numberNeededMark int) bool {
	return 2*(sum+5*numberNeededMark) >= 7*(n+numberNeededMark)
}

func binSearch(l, r, n, sum int) int {
	for r-l > 1 {
		mid := (r + l) / 2

		if isMarkGood(sum, n, mid) {
			r = mid
		} else {
			l = mid + 1
		}

	}
	if isMarkGood(sum, n, l) {
		return l
	}
	return r
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	a, _ := strconv.Atoi(cleanInput(input))

	input, _ = reader.ReadString('\n')
	b, _ := strconv.Atoi(cleanInput(input))

	input, _ = reader.ReadString('\n')
	c, _ := strconv.Atoi(cleanInput(input))

	n := a + b + c

	fmt.Println(binSearch(0, n, n, 2*a+3*b+4*c))
}
