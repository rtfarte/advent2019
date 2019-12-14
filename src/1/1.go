package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0.0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f, _ := strconv.ParseFloat(scanner.Text(), 64)
		sum = sum + math.Floor((f/3)-2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
