package BasicWebCrawler

import (
	"log"
	"net/http"
)

//CompanyLogoGET is a handler for the Logo GET requests
func CompanyLogoGET(w http.ResponseWriter, r *http.Request)  {
	res := ExtractDataFromSourceHTML()
	_,err := w.Write([]byte(res))

	if err != nil{
		log.Println(err)
	}
}