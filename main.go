//http://localhost:3000

package main

import (
	st "./structs"
	js "encoding/json"
	f "fmt"
	io "io/ioutil"
	"log"
	"net/http"
	"os"
)

type PostReq struct {
	ID int `json:"id"`
}

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

func RestHandler(w http.ResponseWriter, r *http.Request) {

	// 処理の最後にBodyを閉じる
	defer r.Body.Close()

	if r.Method == "POST" {

		// リクエストボディをJSONに変換
		var postreq PostReq
		decoder := js.NewDecoder(r.Body)
		err := decoder.Decode(&postreq)
		if err != nil { // エラー処理
			log.Fatal(err)
		}

		// jsonfile読み取り
		read_file, err := io.ReadFile("api/data.json")
		if err != nil {
			log.Fatal(err)
		}
		var poke []st.Pokemon
		// JSONから構造体へ変換
		if err := js.Unmarshal(read_file, &poke); err != nil {
			log.Fatal(err)
		}

		// リクエストボディのIDに該当するナンバーのデータをテキストファイルに書き出し
		// ファイル名：ナンバー.txt
		for _, p := range poke {
			if p.No == postreq.ID {

				filename := f.Sprintf("%d.txt", p.No)

				// ファイル作成
				file, err := os.Create(filename)
				if err != nil {
					log.Fatal(err)
				}
				// 処理終了後ファイルクローズ
				defer file.Close()

				// ファイルにナンバーに該当する名前を書き込む
				_, err = file.WriteString(p.Name)
				if err != nil {
					log.Fatal(err)
				}
				break
			}
		}
		// レスポンスとしてステータスコード201を送信
		w.WriteHeader(http.StatusCreated)

	}
}

func main() {
	// ルーティング、呼び出すハンドラ指定？？
	// ブラウザ出力
	http.HandleFunc("/", IndexHandler)

	http.HandleFunc("/create", RestHandler)

	// ポート指定？？
	http.ListenAndServe(":3000", nil)

}
