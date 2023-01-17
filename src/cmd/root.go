package cmd

import (
	"fmt"
	"github.com/yugo-ibuki/want-article-json/src/utils"
)

func Root() {
	apiKey, err := config.GetApiKey()
	if err != nil {
		fmt.Printf("Can't read env file: %v", err)
	}

	fmt.Println(apiKey)
}
