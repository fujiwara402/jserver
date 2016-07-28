package jserver

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/pkg/errors"
)

/*range .Properties*/
type /*spaceToUpperCamelCase .Title*/ struct {
	/*range .Properties*//*spaceToUpperCamelCase .Title*/ /*if eq .Type "object"*//*spaceToUpperCamelCase .Schema.Title*//*else*//*typeConvertToGo .Schema.Type*//*end*/ `json:"/*.Key*/"`
	/*end*/
}
/*end*/

var ( 
	Status = Sample{}
)

func NewSample() *Sample{
	return &Sample{}
}

func (h NullAdmitInt) MarshalJSON() ([]byte, error) {
	if h.Valid {
		return json.Marshal(h.Int)
	}

	return json.Marshal([]byte(nil))
}

func (h *NullAdmitInt) UnmarshalJSON(b []byte) error {
	var i interface{}
	err := json.Unmarshal(b, &i)
	if err != nil {
		return err
	}

	h.Valid = true
	h.Int = int(int64(i.(float64)))

	return nil
}

func (s *Sample) FromJSON(b []byte) error {
	err := json.Unmarshal(b, s)
	if err != nil {
		return err
	}

	err = s.ValidateFromJSON()
	if err != nil {
		return err
	}

	return nil
}

func (s *Sample) ToJSON() ([]byte, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	err = s.ValidateToJSON()
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (s *Sample) SaveStatus() error {
	Status = *s
	return nil
}

func (s *Sample) ValidateFromJSON() error {
	return nil
}

func (s *Sample) ValidateToJSON() error {
	return nil
}

func GetSampleHandler(w http.ResponseWriter, r *http.Request) {
	j, err := Status.ToJSON() 
	if err != nil {
		Failed(&w,errors.Wrap(err,"Decoding Error"))
		return
	}

	Success(&w, j)
	return
}

func PostSampleHandler(w http.ResponseWriter, req *http.Request) {
	c, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		Failed(&w,errors.Wrap(err, "Request Error"))
		return
	}

	s := NewSample()
	err = s.FromJSON(c) 
	if err != nil {
		Failed(&w,errors.Wrap(err, "Decoding Error"))
		return
	}

	err = s.SaveStatus()
	if err != nil {
		Failed(&w,errors.Wrap(err, "Save Sample Error"))
	}
	
	Success(&w, nil)
	return
}

func Success(w *http.ResponseWriter, b []byte) {
	(*w).WriteHeader(200)
	(*w).Header().Set("Content-Type", "application/json")
	_, err := (*w).Write(b)
	if err != nil {
		return
	}
	return
}

func Failed(w *http.ResponseWriter, e error) {
	(*w).WriteHeader(400)
	_, err := (*w).Write([]byte(e.Error()))
	if err != nil {
		return
	}
	return
}

func Start() error {
	http.HandleFunc("/post", PostSampleHandler)
	http.HandleFunc("/get", GetSampleHandler)

	log.Printf("Start Go HTTP Server")
	fmt.Println("POST:http://localhost:4000/post")
	fmt.Println(" GET:http://localhost:4000/get")

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return err
	}
	return nil
}

