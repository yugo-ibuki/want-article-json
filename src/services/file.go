package services

import (
	"encoding/json"
	"fmt"
	"github.com/yugo-ibuki/want-article-json/src/structures"
	"log"
	"os"
	"time"
)

func Make(i []structures.Item) {
	// dataというディレクトリがなければ作成
	if _, err := os.Stat("data"); err != nil {
		os.Mkdir("data", 0777)
	}

	// 日時のファイル名をdataディレクトリ配下に作成
	now := time.Now()
	fp, err := os.Create("data/" + now.Format("2000-01-01-15-04-05") + ".json")
	if err != nil {
		log.Fatal(err)
	}

	defer fp.Close()
	// ファイルに書き込み
	write(fp, i)
}

func write(fp *os.File, i []structures.Item) {
	// dataをjson化
	j, err := json.Marshal(i)
	if err != nil {
		fmt.Println("json化に失敗しました")
		log.Fatal(err)
	}

	// jsonをファイルに書き込み
	_, err = fp.Write(j)
	if err != nil {
		fmt.Println("ファイルへの書き込みに失敗しました")
		log.Fatal(err)
	}
}
