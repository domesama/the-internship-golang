package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func DigitHangman(){
	slice := [12]int{}
	res := string("____________")
	var count int
	for i:=0; i < 12; i++{
		var temp int
		_,err := fmt.Scan(&temp)

		if err != nil{
			fmt.Println(err)
		}
		slice[i] = temp
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
			count += len(pos)
			for _,val := range pos{
				temp := []rune(res)
				temp[val] = []rune(strconv.Itoa(num))[0]
				res = string(temp)
			}
		} else{
			//res = append( res,[]rune(strconv.Itoa(num))[0])
			res += strconv.Itoa(num)
		}
		fmt.Println(addSpace(string(res)))
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

func addSpace(s string) string {
	buf := &bytes.Buffer{}
	for _, r := range s {
		buf.WriteRune(' ')
		buf.WriteRune(r)
	}
	return buf.String()
}