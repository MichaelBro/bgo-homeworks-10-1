package main

import (
	"bgo-homeworks-10/pkg/currencies"
	"log"
	"net/http"
)

func main() {
	url := "https://raw.githubusercontent.com/netology-code/bgo-homeworks/master/10_client/assets/daily.xml"
	res, err := http.Get(url)

	if err != nil {
		log.Println(err)
	}

	err = currencies.XmlToJson(res.Body)
}
