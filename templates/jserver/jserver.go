package jserver

import (
	//"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


/*range .Properties*/
type /*spaceToUpperCamelCase .Title*/ struct {
	/*range .Properties*//*spaceToUpperCamelCase .Title*/ /*if eq .Type "object"*//*spaceToUpperCamelCase .Schema.Title*//*else*//*typeConvertToGo .Schema.Type*//*end*/ `json:"/*.Key*/"`
	/*end*/
}
/*end*/


var d JsonStruct
//
//func (h Hoge) MarshalJSON() ([]byte, error) {
//	if h.Valid {
//		return json.Marshal(h.Int)
//	}
//
//	return json.Marshal([]byte(nil))
//}
//
//func (h *Hoge) UnmarshalJSON(b []byte) error {
//	var i interface{}
//	err := json.Unmarshal(b, &i)
//	if err != nil {
//		return err
//	}
//
//	h.Valid = true
//	h.Int = int(int64(i.(float64)))
//
//	return nil
//}
//
func get(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func post(w http.ResponseWriter, req *http.Request) {
	d = JsonStruct{}

	content, _ := ioutil.ReadAll(req.Body)
	str := string(content)
	err := json.Unmarshal([]byte(str), &d)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}

func Start() {
	http.HandleFunc("/post", post)
	http.HandleFunc("/get", get)

	log.Printf("Start Go HTTP Server")
	fmt.Println("POST:http://localhost:4000/post")
	fmt.Println(" GET:http://localhost:4000/get")
	err := http.ListenAndServe(":4000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

