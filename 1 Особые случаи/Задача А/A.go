package main

import (
	"fmt"
)

func main() {
	var tRoom, tWant int
	var state string

	fmt.Scan(&tRoom, &tWant)
	fmt.Scan(&state)

	if state == "fan" || tRoom >= tWant && state == "heat" || tRoom <= tWant && state == "freeze" {
		fmt.Println(tRoom)
	} else {
		fmt.Println(tWant)
	}

}
