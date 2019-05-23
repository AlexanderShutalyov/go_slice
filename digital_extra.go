package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputStr = "4539 1488 0343 6467"

func checkToSlice(inputStr string) []int {
	var numTestValue int
	var numSlice []int
	readyString := strings.Replace(inputStr, " ", "", -1)
	if len(readyString) < 2 && len(readyString)%2 != 0 {
		fmt.Println("error: Unexpected length of the string")
		os.Exit(1)
	}

	for i, v := range readyString {
		if i%2 != 0 {
			if serv, err := strconv.Atoi(string(v)); err == nil {
				if numTestValue = serv * 2; numTestValue > 9 {
					numTestValue = numTestValue - 9
				}
				numSlice = append(numSlice, numTestValue)
			} else {
				fmt.Println("error: Unexpected char in the string")
				os.Exit(1)
			}
		}
	}

	return numSlice
}

func checkExtSlice(inputStr string) bool {
	var sum = 0
	for _, v := range checkToSlice(inputStr) {
		sum += v
	}

	if sum%10 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(checkExtSlice(inputStr))
}
