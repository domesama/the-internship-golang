package main

import (
	"log"
	"net/http"
	"the-internship-assignments/BasicWebCrawler"
)

func main(){

	//Greetings, please uncomment each following codes to run

	//First Question
	//SortingAcronyms()

	//SecondQuestion
	//fmt.Print(FloatingPrime())

	//ThirdQuestion
	//DigitHangman()

	//Fourth Question

	//4.1 Only
	//BasicWebCrawler.ExtractDataFromSourceHTML()

	//4.2 with the integration of standard net/http router from go with the integration of previous functions from 4.1
	http.HandleFunc("/logo", BasicWebCrawler.CompanyLogoGET)
	log.Fatal(http.ListenAndServe(":8080", nil))

}