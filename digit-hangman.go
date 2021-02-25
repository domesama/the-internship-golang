package main

import (
	"fmt"
	"log"
	"strconv"
)

func DigitHangman(){
	slice := [12]int{}
	res := string("____________")
	var count int

	fmt.Println("Please input 12 numbers that ranges from 0-9")
	for i:=0; i < 12; i++{
		var temp int
		_,err := fmt.Scan(&temp)

		if err != nil{
			fmt.Println(err)
		}
		if temp >= 0 && temp <= 9{
			slice[i] = temp
		}else{
			log.Fatal("The given input range is beyond 0-9, please re-input the range")
		}
	}
	fmt.Println(slice)

	for i:= 0; i < 5; i++{
		var num int
		_,err := fmt.Scan(&num)

		if err != nil{
			fmt.Println(err)
		}
		pos, found := contains(slice, num)
		if found{
			existed,_ := strconv.ParseInt( string([]rune(res)[pos[0]]), 10, 4)
			if existed != int64(num) {
				count += len(pos)
				for _, val := range pos {
					temp := []rune(res)
					temp[val] = []rune(strconv.Itoa(num))[0]
					res = string(temp)
				}
			}
		} else{
			//res = append( res,[]rune(strconv.Itoa(num))[0])
			res += " " + strconv.Itoa(num)
		}
		fmt.Println(string(res))
	}
	fmt.Println("Your total score is:" ,count)
}

func contains(s [12]int, num int) ([]int, bool) {
	var res []int

	for i, v := range s {
		if v == num {
			res = append(res, i)
		}
	}
	if len(res) <= 0{
		return nil, false
	}
	return res, true
}