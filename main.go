//http://localhost:3000

package main

import (
	st "./structs"
	"bytes"
	js "encoding/json"
	f "fmt"
	"github.com/gorilla/mux"
	io "io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// POSTのBody変換用
type PostReq struct {
	ID int `json:"id"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// jsonfile読み取り
	read_file, err := io.ReadFile("api/data.json")
	if err != nil {
		f.Println(err)
		return
	}

	var stpoke []st.Pokemon

	if err := js.Unmarshal(read_file, &stpoke); err != nil {
		f.Println(err)
		return
	}

	// No:Nameの形式で出力
	for _, p := range stpoke {
		f.Fprintf(w, "%d : %s\n", p.No, p.Name)
	}

}

func CreateTextHandler(w http.ResponseWriter, r *http.Request) {

	// 処理の最後にBodyを閉じる
	defer r.Body.Close()

	// リクエストボディをJSONに変換
	var postreq PostReq
	decoder := js.NewDecoder(r.Body)
	err := decoder.Decode(&postreq)
	if err != nil {
		f.Println(err)
		return
	}

	// jsonfile読み取り
	read_file, err := io.ReadFile("api/data.json")
	if err != nil {
		f.Println(err)
		return
	}
	var stpoke []st.Pokemon
	// JSONから構造体へ変換
	if err := js.Unmarshal(read_file, &stpoke); err != nil {
		f.Println(err)
		return
	}

	// リクエストボディのIDに該当するナンバーのデータをテキストファイルに書き出し
	// ファイル名：ナンバー.json
	for _, p := range stpoke {
		if p.No == postreq.ID {

			filename := f.Sprintf("%d.json", p.No)

			// ファイル作成
			file, err := os.Create(filename)
			if err != nil {
				f.Println(err)
				return
			}
			// 処理終了後ファイルクローズ
			defer file.Close()

			// 構造体→json変換
			jsp, err := js.Marshal(p)
			if err != nil {
				f.Println(err)
				return
			}

			// ファイル出力用整形処理
			jspout := new(bytes.Buffer)
			// スペース4つでインデント付ける
			js.Indent(jspout, jsp, "", "    ")

			// ファイルにナンバーに該当するデータを書き込む
			_, err = file.WriteString(jspout.String())
			if err != nil {
				f.Println(err)
				return
			}
			break
		}
	}
	// レスポンスとしてステータスコード201を送信
	w.WriteHeader(http.StatusCreated)
}

/*
func AllGetTextHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	filename := f.Sprintf("%d.txt", id)
	pokename, err := io.ReadFile(filename)
	if err != nil {
		f.Println(err)
	}
	getpoke := GetRes{
		ID:   id,
		Name: string(pokename),
	}
	js.NewEncoder(w).Encode(getpoke)
}
*/

func SingleGetTextHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// パスパラメータ取得
	getid := mux.Vars(r)

	// パスパラメータのidをint型に変換
	id, err := strconv.Atoi(getid["id"])
	if err != nil {
		f.Println(err)
		return
	}

	// POSTされたファイル探索
	filename := f.Sprintf("%d.json", id)
	jspoke, err := io.ReadFile(filename)
	if err != nil {
		f.Fprintf(w, "Not such post-data")
		return
	}
	var stpoke st.Pokemon
	if err := js.Unmarshal(jspoke, &stpoke); err != nil {
		f.Println(err)
		return
	}
	js.NewEncoder(w).Encode(stpoke)
}

func main() {
	// ルーティング、ハンドラ指定
	// ブラウザ出力
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)

	r.HandleFunc("/create", CreateTextHandler).Methods("POST")
	//	r.HandleFunc("/get", AllGetTextHandler).Methods("GET")
	r.HandleFunc("/get/{id}", SingleGetTextHandler).Methods("GET")

	// ポート指定(ポート番号は適当)
	log.Fatal(http.ListenAndServe(":3000", r))

}
