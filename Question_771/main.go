package main

import (
	"fmt"
)

func numJewelsInStones(J string, S string) int {
	dict := make(map[string]int)
	for _, r := range S {
		fmt.Printf("%q\n", r)
		_, exists := dict[string(r)]
		if exists {
			dict[string(r)]++
		} else {
			dict[string(r)] = 1
		}
	}

	var count int
	for _, r := range J {
		fmt.Printf("%q", r)
		value, _ := dict[string(r)]
		count += value
	}

	return count
}

func main() {
	j := "aA"
	s := "aAAbbbbb"
	answer := numJewelsInStones(j, s)
	fmt.Printf("Output:%v", answer)
}
