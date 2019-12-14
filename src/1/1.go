package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := uint32(0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f, _ := strconv.ParseUint(scanner.Text(), 10, 64)
		sum = sum + (uint32(f) / 3) - 2
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
