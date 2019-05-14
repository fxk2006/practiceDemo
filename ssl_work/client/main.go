package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post("https://uatsky.yesbank.in:444/app/uat/fundtransfer2R/getbalance","application/json", nil)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
