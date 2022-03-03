package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type coinbaseResponse struct {
	Data data
}
type data struct {
	Base     string
	Currency string
	Amount   string
}

var p string

func main() {
	go getBitcoinPrice()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, p)
}

func getBitcoinPrice() string {
	sum := 0
	for {
		sum++
		var result coinbaseResponse
		response, _ := http.Get("https://api.coinbase.com/v2/prices/spot?currency=GBP")
		body, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(body, &result)
		p = result.Data.Amount
		fmt.Printf("Updated price to %s\n", p)
		time.Sleep(60 * time.Second)
	}
}
