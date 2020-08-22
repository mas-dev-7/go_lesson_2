# go_lesson_2

```
go run main.go
```

別のターミナル開いて

## POST
```
curl -X POST http://localhost:3000/create -d '{"id":『ポケモンの図鑑ナンバー』}'
```
### 例
ポケモンの図鑑ナンバーのところを1にして実行すると、main.goと同じ階層に1.jsonが作成され、中身はフシギダネのデータ
<br>
<br>
## GET
```
curl -X GET http://localhost:3000/get/『POSTしたid』
```
### 例
『POSTしたid』に1を入れて実行すると、フシギダネのデータがjson形式で返ってくる
<br>予めPOSTしたデータのみしか返ってこない
