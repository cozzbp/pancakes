package main

import (
	"fmt"
	"strconv"

	"github.com/cozzbp/pancakes/pancakes"
)

func main() {
	fmt.Println("Enter the number of test cases T, followed by T test cases")
	fmt.Println("ex:")
	fmt.Println("2")
	fmt.Println("-+")
	fmt.Println("+-")

	var input string
	fmt.Scanln(&input)
	numCases, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	if numCases < 1 || numCases > 100 {
		fmt.Println("Enter a number between 1-100")
		return
	}
	cases := make([]string, numCases)
	for i := range cases {
		fmt.Scanln(&cases[i])
	}
	for i, v := range cases {
		p := pancakes.New(v)
		num := p.Flip()
		fmt.Printf("Case #%d: %d\n", i+1, num)
	}
}
