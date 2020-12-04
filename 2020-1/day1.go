package main

import (
    "bufio"
    "fmt"
    "log"
	"os"
	"strconv"
)

func main() {
	// Save input from https://adventofcode.com/2020/day/1/input, data content might be different from users
	file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	arr := make([]int, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		if scanner != nil || scanner.Text() != "" {
			number, _ := strconv.Atoi(scanner.Text())
			arr = append(arr, number)
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}

	// 2020-1 advent code is here "https://adventofcode.com/2020/day/1"
	fmt.Println("Anser for problem 1: ", problem1(arr))
	fmt.Println("Anser for problem 2: ", problem2(arr))
}

func problem1 (arr []int) int {
	for i:=0; i<len(arr); i++ {
		for j:=i+1; j<len(arr); j++ {
			if (arr[i] + arr[j]) == 2020 {
				fmt.Println("[DEBUG] ", arr[i], arr[j])
				return arr[i] * arr[j]
			}

		}
	}
	return -1
}

func problem2 (arr []int) int {
	for i:=0; i<len(arr); i++ {
		for j:=i+1; j<len(arr); j++ {
			for z:=j+1; z<len(arr); z++ {
				if (arr[i] + arr[j] + arr[z]) == 2020 {
					fmt.Println("[DEBUG] ", arr[i], arr[j], arr[z])
					return arr[i] * arr[j] * arr[z]
				}
			}
		}
	}
	return -1
}