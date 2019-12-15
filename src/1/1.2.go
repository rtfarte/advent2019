package main

import (
	"advent2019/src/utils"
	"fmt"
	"log"
)

func main() {

	a, err := utils.Read()
	if err != nil {
		log.Fatal(err)
	}
	sum := 0
	for _, f := range a {
		sum += fuel(f)
	}

	fmt.Println(sum)
}

func fuel(v int) int {
	sum := 0
	x := v
	for x > 0 {
		x = (x / 3) - 2
		if x > 0 {
			sum += x
		}
	}
	return sum
}
