package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// INI pulls current config from SBC
func (c Customer) INI() *string {
	client := &http.Client{Timeout: 3 * time.Second}
	req, err := http.NewRequest("GET", c.AC.URL, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(c.AC.User, c.AC.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Failed to connect to Audiocodes API")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	b := string(body)
	return &b
}
