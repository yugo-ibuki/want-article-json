package services

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Item struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func SearchArticle(args []string) {
	resp, err := http.Get("https://qiita.com/api/v2/items?query=" + args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data []Item

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	// dataというディレクトリがなければ作成
	if _, err := os.Stat("data"); err != nil {
		os.Mkdir("data", 0777)
	}

	// 日時のファイル名をdataディレクトリ配下に作成
	time := time.Now()
	fp, err := os.Create("data/" + time.Format("2000-01-01-15-04-05") + ".json")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	// dataをjson化
	j, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// jsonをファイルに書き込み
	fp.Write(j)
}
