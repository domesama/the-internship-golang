package main

import (
	"fmt"
	"strconv"
	"strings"
	"math/big"
)

func FloatingPrime() bool{
	var input float64

	_,err := fmt.Scan(&input)

	if err != nil{
		fmt.Println(err)
	}

	if input < 1 || input > 10{
		fmt.Println("The given input is invalid")
	}

	str := strconv.FormatFloat(input, 'f', 5, 64)

	splited := strings.Split(str, ".")
	temp := splited[0]
	//fmt.Print(splited)

	for _,val := range string([]rune(splited[1])){
		temp += string(val)
		n, _ := strconv.ParseInt(temp,10,32 )
		if big.NewInt(n).ProbablyPrime(0){
			fmt.Println(n)
			return true
		}
	}
	return false
}
