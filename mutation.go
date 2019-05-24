package main

import (
	"fmt"
	"strings"
)

var dnkFirst = "GAGCCTACTAACGGGAT"
var dnkSecond = "CATCGTAATGACGGCCT"

func checkstr() bool{

	if len(dnkFirst) != len (dnkSecond) {
		return false
	}
	return true
}

func main() {

	count:=0

	if checkstr() {
		for i, v := range dnkFirst {
			if strings.EqualFold(string(v), string(dnkSecond[i])) {
				count++
			}
		}
		fmt.Println("Both DNK have", count, "common runes")
	} else {
		fmt.Println("error in the input strings")
	}


}
