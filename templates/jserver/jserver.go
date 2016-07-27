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

// 関数には適切な名前をつける。特に振る舞いが特別なものは気をつける
func GetJSONHandler(w http.ResponseWriter, r *http.Request) {

	// package名と被る命名はバグを生み出す原因なのでやめる
	j, err := json.Marshal(d)
	if err != nil {
		// Postに書いた通り
		w.WriteHeader(400)
		log.Println(err)
		return
	}
	// status codeを返す
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// 同上
func PostJSONHandler(w http.ResponseWriter, req *http.Request) {

	// すぐ死ぬ変数は変数名を短くしておく
	// あとちゃんとエラーを取る
	c, err := ioutil.ReadAll(req.Body)
	defer req.Close()
	if err != nil {
		// とりあえず雑にエラーメッセージを返しておく
		// エラーが発生して、その後の実行に影響が出ると思われるものは
		// その場で終わらせておく
		// responseに適切に起きたエラーをstatus codeと一緒に返す
		log.Println(err)
		w.WriteHeader(400)
		return
	}
	// []byte -> string -> []byteの型変換を行っていたのでstrを削除

	// 宣言は使う前にすると読みやすい
	d = JSONStruct{}
	err = json.Unmarshal(c, &d)
	if err != nil {
		// log packageを入れているのでとりあえずfmtではなくlogを使う
		// print debugで挿入したfmtはちゃんと消しておく
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// errorを返すようにする
func Start() error {
	http.HandleFunc("/post", PostJSONHandler)
	http.HandleFunc("/get", GetJSONHandler)

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

