# アプリ概要

読みたい記事の記号を入れることで、

関連記事を引っ張ってきて、JSON形式で保存してくれるアプリ

# 参考記事

https://qiita.com/tag1216/items/b0b90e30c7e581aa2b00

# 使い方
```go
go mod tidy
```

```go
go run main.go {keyword}
```

ルートにdataディレクトリに
キーワードで検索したQiita記事を最新20件取得した結果が
JSON形式で保存されます。

データ構造は以下です。
```go
[
    {
        title: string,
        created_at: Date,
    }
]
```
