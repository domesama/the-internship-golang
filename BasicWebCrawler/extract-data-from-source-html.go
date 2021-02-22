package BasicWebCrawler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Companies struct {
	Companies []CompanyLogo `json:"companies"`
}
type CompanyLogo struct {
	Logo string `json:"logo"`
}

func NewCompanies() *Companies {
	return &Companies{}
}

func (c *Companies) AppendLogos(logoStr string){
	c.Companies = append(c.Companies, CompanyLogo{
		logoStr,
	})
}

func (c *Companies) Redundant(str string) bool{
	for _,val := range c.Companies{
		if val.Logo == str{
			return true
		}
	}
	return false
}


var companies = NewCompanies()

// This will get called for each HTML element found
func processElement(index int, element *goquery.Selection){
	// See if the href attribute exists on the element
	src, exists := element.Attr("src")
	if exists {
		//fmt.Println(src)
		temp := strings.Split(src, "/")
		if temp[0] == "company"{
			formatted := "https://theinternship.io/"+src
			fmt.Println(src)
			if !companies.Redundant(formatted){
				companies.AppendLogos(formatted)
			}
		}
	}
}

func ExtractDataFromSourceHTML() string{
	// Make HTTP request
	response, err := http.Get("https://theinternship.io")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find all links and process them with the function
	// defined earlier
	document.Find("img").Each(processElement)

	c,err := json.Marshal(companies)
	if err != nil{
		fmt.Println(err)
	}
	return string(c)
}
