package main

import (
	st "./structs"
	js "encoding/json"
	f "fmt"
	io "io/ioutil"
	"log"
	_ "os"
)

func main() {
	// jsonfile読み取り
	read_file, err := io.ReadFile("api/data.json")
	if err != nil {
		log.Fatal(err)
	}

	var poke []st.Pokemon

	if err := js.Unmarshal(read_file, &poke); err != nil {
		log.Fatal(err)
	}

	// テストjson file	全権取得
	// f.Println(json)

	// テスト一部検索
	for _, p := range poke {
		f.Printf("%d : %s\n", p.No, p.Name)
	}

}
