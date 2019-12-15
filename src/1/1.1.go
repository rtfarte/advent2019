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
		sum = sum + (f / 3) - 2
	}

	fmt.Println(sum)
}
