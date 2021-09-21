package main

import (
	"fmt"
	"strconv"

	"tudai2021.com/model"
)

func parser(inp string) model.Result {
	r1 := model.NewResult(inp[0:2], stringToInt(inp[2:4]), inp[4:])
	return r1
}

func stringToInt(s string) int {
	i1, _ := strconv.Atoi(s)
	return i1
}

func main() {
	fmt.Println(parser("TX13ABCDEFJKLMSSSSS"))
}
