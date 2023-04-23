package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func test() {
	var err error
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://ipinfo.io/AS13335/json", nil)
	if err != nil {
		log.Panic(err)
	}
	token := "fa204583ea0443"
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	var buff bytes.Buffer
	buff.ReadFrom(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(buff.String())
}
