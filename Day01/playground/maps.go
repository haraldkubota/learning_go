package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words:=strings.Fields(s)
	res:=make(map[string]int)
	for _,v := range words {
		res[v]=res[v]+1
	}
		
	return res
}

func main() {
	wc.Test(WordCount)
}
