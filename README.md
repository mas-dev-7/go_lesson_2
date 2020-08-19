# go_lesson_1
pokemon

```
go run main.go
```

別のターミナル開いて

```
curl -X POST http://localhost:3000/create -d '{"id":『ポケモンの図鑑ナンバー』}'
```
ポケモンの図鑑ナンバーのところを1にして実行すると、main.goと同じ階層に1.txtが作成され、中身はフシギダネ
