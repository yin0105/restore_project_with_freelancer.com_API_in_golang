package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	// postData := make([]byte, 100)
	req, err := http.NewRequest("GET", "https://www.freelancer.com/api/projects/0.1/projects/active/?compact=&project_types%5B%5D=fixed&max_avg_price=500&min_avg_price=250&query=django", nil)
	if err != nil {
		os.Exit(1)
	}
	req.Header.Add("freelancer-oauth-v1", "1Dik9bnPVKncY80lae7OeE7mg1JR5r")
	resp, err := client.Do(req)
	if err != nil {
		print("OK")
	}
	fmt.Println(resp)
	fmt.Println("#########################")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	fmt.Println(string(body[:]))

	// postData := make([]byte, 100)
	// req, err := http.NewRequest("POST", "http://example.com", bytes.NewReader(postData))
	// if err != nil {
	// 	os.Exit(1)
	// }
	// req.Header.Add("User-Agent", "myClient")
	// resp, err := client.Do(req)
	// defer resp.Body.Close()
	// fmt.Println(resp)
}
