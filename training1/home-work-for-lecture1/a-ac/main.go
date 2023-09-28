package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}

func getArgs(data []byte) (int, int, string) {
	lines := strings.Split(string(data), "\n")
	tRoom, err := strconv.Atoi(strings.Fields(lines[0])[0])
	if err != nil {
		log.Fatal(err)
	}
	tCond, err := strconv.Atoi(strings.Fields(lines[0])[1])
	if err != nil {
		log.Fatal(err)
	}
	mode := lines[1]
	return tRoom, tCond, mode
}

func resTemp(tRoom, tCond int, mode string) int {
	"freeze" // can only decrease temp, if needed
	"heat"   // can only increase temp, if needed
	"auto"   // does both
	"fan"    // doesn't change temp
	return 0
}
