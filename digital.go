package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputString = "4539 1488 0343 6467"

func getArray(inputString string) bool {
	var numSlice []int
	var numTestValue int
	count := 0
	readyString := strings.Replace(inputString, " ", "", -1)

	if len(string(readyString))%2 != 0 {
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
				count += numTestValue
			} else {
				fmt.Println("error: Unexpected char in the string")
				os.Exit(1)
			}
		}
	}

	fmt.Println(count)
	if count%10 == 0 {
		return true
	}
	return false

}
func main() {
	fmt.Println(getArray(inputString))
}
