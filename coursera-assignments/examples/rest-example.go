package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	APIURL := "https://httpbin.org/get"
	// This example uses the NewRequest function, which is cool.
	req, err := http.NewRequest(http.MethodGet, APIURL, nil)
	if err != nil {
		panic(err)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", string(body))

	// Here is a second method to do basically the same thing
	// This method actually looks a lot more elegant than the one above
	// NOTE the blank identifier _ This basically means ignore the error if one occures during the data assignment to ioutil
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=USD")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	// Lets do a POST!
	// First things first, create the JSON map for the data to post
	jsonData := map[string]string{"firstname": "Brian", "lastname": "Bartwood"}

	// Next, Marshal that app into actual JSON, meaning you need to transform the memory representation of an object into json.
	jsonValue, _ := json.Marshal(jsonData)

	// Next, you need to send the post command. Notice how you include the header information "application/json" as well.
	// Note: you don't need to initialize the variable resonse and err because they are already initialized above.
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	// LETS TRY A POST ANOTHER WAY. DEDCIDE WHICH ONE TO USE

	request, _ := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client = &http.Client{}
	response, err = client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

}
