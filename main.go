//http://localhost:3000

package main

import (
	st "./structs"
	js "encoding/json"
	f "fmt"
	io "io/ioutil"
	"log"
	"net/http"
	_ "os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// jsonfile読み取り
	read_file, err := io.ReadFile("api/data.json")
	if err != nil {
		log.Fatal(err)
	}

	var poke []st.Pokemon

	if err := js.Unmarshal(read_file, &poke); err != nil {
		log.Fatal(err)
	}

	// No:Nameの形式で出力
	for _, p := range poke {
		f.Fprintf(w, "%d : %s\n", p.No, p.Name)
	}

}

func main() {
	// ルーティング、呼び出すハンドラ指定？？
	http.HandleFunc("/", IndexHandler)

	// ポート指定？？
	http.ListenAndServe(":3000", nil)

}
