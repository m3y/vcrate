package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	currency = flag.String("c", "xrp", "currency")
)

func vcrate(currency string) (body string, retErr error) {
	url := "https://coincheck.com/api/rate/" + currency + "_jpy"
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(resBody), nil
}

func main() {
	flag.Parse()

	body, err := vcrate(*currency)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)
}
