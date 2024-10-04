package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func cleanInput(input string) string {
	return strings.TrimSpace(input)
}

func strToIntSlice(input, sep string, intArr []int) []int {

	splitedInput := strings.Split(cleanInput(input), sep)

	for ind, el := range splitedInput {
		intEl, _ := strconv.Atoi(el)
		intArr[ind] = intEl
	}

	return intArr
}

func binSearch(arr *[]int, n, el int) int {
	var l, r, mid int

	l, r = 0, n-1

	for r-l > 1 {
		mid = (r + l) / 2

		if el <= (*arr)[mid] {
			r = mid
		} else {
			l = mid
		}

		mid = (r + l) / 2
	}

	if math.Abs(float64((*arr)[l]-el)) <= math.Abs(float64((*arr)[r]-el)) {
		mid = l
	} else {
		mid = r
	}

	return (*arr)[mid]
}

func main() {
	var input string

	var splitedInput []string

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	input, _ = reader.ReadString('\n')
	splitedInput = strings.Split(cleanInput(input), " ")

	n, _ := strconv.Atoi(splitedInput[0])
	k, _ := strconv.Atoi(splitedInput[1])
	sortedArr := make([]int, n)
	inputNum := make([]int, k)

	input, _ = reader.ReadString('\n')
	sortedArr = strToIntSlice(input, " ", sortedArr)
	input, _ = reader.ReadString('\n')
	inputNum = strToIntSlice(input, " ", inputNum)

	pointerToSortedArr := &sortedArr

	for _, el := range inputNum {
		writer.WriteString(strconv.Itoa(binSearch(pointerToSortedArr, n, el)) + "\n")
	}

	writer.Flush()
}
