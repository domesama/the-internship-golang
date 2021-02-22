package BasicWebCrawler

import (
	"log"
	"net/http"
)

//Created a handler for the Logo GET requests
func CompanyLogoGET(w http.ResponseWriter, r *http.Request)  {
	res := ExtractDataFromSourceHTML()
	_,err := w.Write([]byte(res))

	if err != nil{
		log.Println(err)
	}
}