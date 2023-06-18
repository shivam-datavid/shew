package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: ./shew <method> <url> optional<data>")
		return
	}

	method := os.Args[1]
	url := os.Args[2]

	if method == "GET" {
		makeGetRequest(url)
	} else if method == "POST" {
		data := os.Args[3]
		makePostRequest(url, data)
	} else {
		fmt.Println("Invalid method. Please use GET or POST.")
	}

}

func makeGetRequest(url string) {

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request")
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("GET Response:", string(body))

}

func makePostRequest(url, data string) {
	payload := strings.NewReader(data)

	response, err := http.Post(url, "application/json", payload)

	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("POST Response:", string(body))
}
