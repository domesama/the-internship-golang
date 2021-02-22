package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"math/big"
)

//FloatingPrime is a function that receives a float and tries to calculate each individual base to check if it's a Prime number or not
func FloatingPrime() bool{
	var input float64

	//Receives input
	fmt.Println("Please enter any number that ranges from 1.0 - 10.0")
	_,err := fmt.Scan(&input)

	if err != nil{
		fmt.Println(err)
	}

	//Check the range of the inputted
	if input < 1 || input > 10{
		log.Fatal("The given input is invalid")
	}

	//Converts to string and splits according to the decimals
	str := strconv.FormatFloat(input, 'f', 3, 64)

	splited := strings.Split(str, ".")
	temp := splited[0]

	//For loops the converted strings has already been splited
	for _,val := range string([]rune(splited[1])){

		//Concat the splited strings together
		temp += string(val)
		n, _ := strconv.ParseInt(temp,10,32 )

		//Check if that number is prime
		if big.NewInt(n).ProbablyPrime(0){
			fmt.Println("Is a prime number at:",n)
			return true
		}
	}
	return false
}
