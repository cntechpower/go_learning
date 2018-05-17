package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type coinsData struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price_usd"`
}

func main() {
	resp, err := http.Get("https://api.coinmarketcap.com/v1/ticker/?limit=0")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	body1 := string(body)
	fmt.Printf(body1)
	if err != nil {
		log.Fatal(err)
	}

	var c []coinsData
	err = json.Unmarshal(body, &c)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%v\n", c)
}
