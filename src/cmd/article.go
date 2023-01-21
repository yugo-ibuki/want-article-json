package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var articleCmd = &cobra.Command{
	Use:   "article",
	Short: "探したい記事のキーワードを入れてください",
	Long:  "キーワードを元にQiitaから記事を取得して一覧を表示します。",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("article called")
	},
}

func init() {
	rootCmd.AddCommand(articleCmd)
}
