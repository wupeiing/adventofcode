package main

import (
	"bufio"
    "fmt"
    "log"
	"os"
	"strings"
	"strconv"
)

func main() {
	// Save input from https://adventofcode.com/2020/day/1/input, data content might be different from users
	file, err := os.Open("input2.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	displayCountRules := make([]string, 0)
	displayChar := make([]string, 0)
	targetPasswords := make([]string, 0)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		displayCountRules = append(displayCountRules, s[0])
		displayChar = append(displayChar, s[1])
		targetPasswords = append(targetPasswords, s[2])
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}

	// 2020-1 advent code is here "https://adventofcode.com/2020/day/2"
	fmt.Println("Anser for problem 1: ", problem1(displayCountRules, displayChar, targetPasswords))
	fmt.Println("Anser for problem 2: ", problem2(displayCountRules, displayChar, targetPasswords))
}

func problem1 (displayCountRules, displayChar, targetPasswords []string) int {
	validCount := 0
	for i:=0; i<len(displayCountRules); i++ {
		sizeSpl := strings.Split(displayCountRules[i], "-")
		min, _ := strconv.Atoi(sizeSpl[0])
		max, _ := strconv.Atoi(sizeSpl[1])
		charCount := strings.Count(targetPasswords[i], string(displayChar[i][0]))
		if (charCount >= min) && (charCount <= max) {
			validCount++
		}
	}
	return validCount
}

func problem2 (displayCountRules, displayChar, targetPasswords []string) int {
	validCount := 0
	for i:=0; i<len(displayCountRules); i++ {
		sizeSpl := strings.Split(displayCountRules[i], "-")
		min, _ := strconv.Atoi(sizeSpl[0])
		max, _ := strconv.Atoi(sizeSpl[1])
		if string(targetPasswords[i][min-1]) == string(displayChar[i][0]) && string(targetPasswords[i][max-1]) != string(displayChar[i][0]) {
			if i == len(displayCountRules)-1 {
				fmt.Println(targetPasswords[i], min, string(targetPasswords[i][min-1]), max, string(targetPasswords[i][max-1]), string(displayChar[i][0]))
			}
			validCount++
		}
	}

	return validCount
}


