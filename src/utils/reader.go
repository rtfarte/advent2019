package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Read() ([]int, error) {
	file, err := os.Open("../inputs/modules.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	a := make([]int, 100)
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		f, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		a[i] = int(f)
		i++
	}

	if err := scanner.Err(); err != nil {
		return a, err
	}

	return a, nil
}
