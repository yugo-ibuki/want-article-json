package services

import (
	"encoding/json"
	"github.com/yugo-ibuki/want-article-json/src/structures"
	"io"
	"log"
	"net/http"
)

func SearchArticle(args []string) ([]structures.Item, error) {
	resp, err := http.Get("https://qiita.com/api/v2/items?query=" + args[0])
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var data []structures.Item

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return data, nil
}
