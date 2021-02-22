package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func SortingAcronyms()  {
	var n int
	var res []string
	fmt.Println("Enter the size")
	_, err := fmt.Scanln(&n)

	if err != nil{
		fmt.Println(err)
	}

	for i := 0; i < n; i++ {
		var tempRes string
		consoleReader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the State name .... \n")

		input, _ := consoleReader.ReadString('\n')

		temp := strings.Split(input, " ")

		for _, str := range temp {
			r := []rune(string(str[0]))
			if unicode.IsUpper(r[0]) {
				tempRes += string(r)
			}
		}
		res = append(res, tempRes)
	}

	printSlice(res)
}
func printSlice(s []string) {
	if len(s) == 0 {
		return
	}
	fmt.Println(s[0])
	printSlice(s[1:])
}
