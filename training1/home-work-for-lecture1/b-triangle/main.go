package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	res := []byte(isTriangle(getArgs(data)))
	err = os.WriteFile("output.txt", res, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func getArgs(data []byte) []int {
	sidesStrings := strings.Split(string(data), "\n")
	sidesNumbers := make([]int, 3)
	var err error
	for i, v := range sidesStrings {
		sidesNumbers[i], err = strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
	}
	return sidesNumbers
}

func isTriangle(sides []int) string {
	// 1. none is zero
	for _, v := range sides {
		if v == 0 {
			return "NO"
		}
	}
	// 2. sum of any two sides is bigger than the third
	sides = append(sides, sides...)
	for i := 0; i < 3; i++ {
		if sides[i] >= sides[i+1]+sides[i+2] {
			return "NO"
		}
	}
	return "YES"
}
