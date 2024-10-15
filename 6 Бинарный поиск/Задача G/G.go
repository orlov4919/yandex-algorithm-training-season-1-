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

func isGoodFrame(frame, n, m, s int) bool {
	return n*m-(n-2*frame)*(m-2*frame) <= s
}

func binSearchFrame(minFrame, maxFrame, n, m, s int) int {
	for maxFrame-minFrame > 1 {
		midFrame := (maxFrame + minFrame) / 2

		if isGoodFrame(midFrame, n, m, s) {
			minFrame = midFrame
		} else {
			maxFrame = midFrame - 1
		}
	}

	if isGoodFrame(maxFrame, n, m, s) {
		return maxFrame
	}

	return minFrame
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(cleanInput(input))

	input, _ = reader.ReadString('\n')
	m, _ := strconv.Atoi(cleanInput(input))

	input, _ = reader.ReadString('\n')
	s, _ := strconv.Atoi(cleanInput(input))

	fmt.Println(binSearchFrame(0, min(m, n)/2, n, m, s))
}
