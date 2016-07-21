package main

import (
	"./jserver"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseResponse(res *http.Response) (string, int) {
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return string(contents), res.StatusCode
}

func assert(data map[string]interface{}) bool {
	b, _ := json.Marshal(data)
	_, err := http.Post("http://localhost:4000/post", "application/json", bytes.NewReader(b))
	if err != nil {
		fmt.Println("Error: POST FAIL")
		return false
	}

	res, _ := http.Get("http://localhost:4000/get")
	c, _ := ParseResponse(res)
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
		"bar":  "bar",
		"buz":  "foo",
		"hoge": 123,
	}
	fmt.Println("assert 1")
	result1 := assert(data1)

	data2 := map[string]interface{}{
		"bar":  "bar",
		"baz":  "hoo",
		"hoge": 0,
	}
	fmt.Println("assert 2")
	result2 := assert(data2)

	data3 := map[string]interface{}{
		"bar": "bar",
		"buz": "foo",
	}

	fmt.Println("assert 3")
	result3 := assert(data3)

	fmt.Println("result:", result1, result2, result3)
}
