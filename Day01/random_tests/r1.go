package main

import (
	"fmt"
)

var testCases = []string{
	"0000",
	"0189",
	"ABEF",
	"abcf",
	"000Z",
	"defg",
	"9i@#",
}

var myArray = [5]int{5, 6, 7, 8, 9}

func main() {
	for k, v := range testCases {
		fmt.Println(k, v)
		fmt.Println(checkHex(v))
	}
	for k, v := range myArray {
		fmt.Println(k, v)
	}

}

func checkHex(s string) bool {
	res := true
	for _, c := range s {
		// fmt.Printf("%c ", c)
		switch {
		case !(c >= '0' && c <= '9' || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F'):
			res = false
		}
	}
	return res
}

func init() {
	fmt.Println("Calling init!")
}
