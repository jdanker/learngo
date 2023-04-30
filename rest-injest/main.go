package main

import (
	"net/http"
	"time"
	"encoding/json"
)

var client *http.Client

type CatFact struct {
	Fact string 
	length int
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)


}

func main(){
	// timeout if no response
	client = &http.Client{Timeout: 10 * time.Second}
}

