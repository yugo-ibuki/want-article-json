package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Item struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func SearchArticle(args []string) {
	resp, err := http.Get("https://qiita.com/api/v2/items?query=" + args[0])
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var data []Item

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	for _, item := range data {
		fmt.Printf("%s %s\n", item.CreatedAt, item.Title)
	}
}
