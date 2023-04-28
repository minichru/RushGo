package examples

import (
	rushgo "RushGo/RushGo"
	"fmt"
	"io/ioutil"
)

func Get() {
	resp, err := rushgo.Get("https://httpbin.org/get")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(string(body))
}

func Post() {
	data := "RushGo is the best"
	resp, err := rushgo.Post("https://httpbin.org/post", []byte(data))

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(body))
}

func JsonPARSE() {
	resp, err := rushgo.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(body))

	data, err := rushgo.ParseJSON(body)
	if err != nil {
		fmt.Println("Error", err)
	}

	headers := data["headers"].(map[string]interface{})

	fmt.Println("Host: ", headers["Host"])
}

func BodyContains() {
	resp, err := rushgo.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(body))

	contains := rushgo.ResponseBodyContains(body, "User-Agent")

	if contains {
		fmt.Print("Found yay")
		//do something
	} else {
		fmt.Println("Damn no find..")
		//do something else
	}
}
