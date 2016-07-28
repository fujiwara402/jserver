package main

import (
	"./jserver"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func parseResponse(res *http.Response) string {
	c, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		log.Println(err)
		return ""
	}

	return string(c)
}

func assert(data map[string]interface{}) bool {
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return false
	}

	_, err = http.Post("http://localhost:4000/post", "application/json", bytes.NewReader(b))
	if err != nil {
		log.Println(err)
		return false
	}

	r, err := http.Get("http://localhost:4000/get")
	if err != nil {
		log.Println(err)
		return false
	}

	c := parseResponse(r)

	fmt.Println("post:", string(b))
	fmt.Println(" get:", c)

	if c == string(b) {
		return true
	}

	return false
}

func main() {
	go jserver.Start()

	data1 := map[string]interface{}{
		"message":        "Hello, World.",
		"created_at":     "2016-05-09T19:45:32Z",
		"null_admit_int": 0,
	}

	fmt.Println("assert 1")
	result1 := assert(data1)

	fmt.Println("result:", result1)
}
