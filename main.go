package main

import (
	"errors"
	"fmt"
	"strconv"

	"tudai2021.com/model"
)

func parser(inp string) (model.Result, error) {
	if inp[0:2] == "NN" || inp[0:2] == "TX" {
		if stringToInt(inp[2:4]) != 0 {
			if len([]rune(inp[4:])) == stringToInt(inp[2:4]) {
				r1 := model.NewResult(inp[0:2], stringToInt(inp[2:4]), inp[4:])
				return r1, nil

			} else {
				return model.Result{}, errors.New("cantidad de caracteres no coinciden")
			}
		} else {
			return model.Result{}, errors.New("debe indicar cantidad en tipo numérico")
		}
	} else {
		return model.Result{}, errors.New("tipo iválido")
	}

}

func stringToInt(s string) int {
	i1, _ := strconv.Atoi(s)
	return i1
}

func main() {
	p, err := parser("TX13ABCDEFJKLMSS2dS")
	fmt.Println(p, err)
}
