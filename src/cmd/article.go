package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	SearchService "github.com/yugo-ibuki/want-article-json/src/services"
)

// versionCmd represents the version command
var articleCmd = &cobra.Command{
	Use:   "article",
	Short: "探したい記事のキーワードを入れてください",
	Long:  "キーワードを元にQiitaから記事を取得して一覧を表示します。",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			fmt.Println("キーワードは一つだけ入力してください")
			os.Exit(1)
		}
		SearchService.SearchArticle(args)
	},
}

func init() {
	rootCmd.AddCommand(articleCmd)
}
