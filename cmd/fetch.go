package cmd

import (
	"encoding/json"
	"fmt"
	"sword/core"

	"github.com/spf13/cobra"
)

// fetchCmd 命令直接提供敏感词表
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "直接获取所有需要检查的敏感词表",
	Long:  `使用 JSON 格式，获取敏感词表.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("数据获取中...")
		if data, err := core.FetchWords(); err == nil {
			fmt.Println("数据获取成功", Y)
			s, _ := json.MarshalIndent(data, "", " ")
			fmt.Println(string(s))
		} else {
			fmt.Println("数据获取失败", N)
			fmt.Println("错误信息为：", err)
			fmt.Println("请检查是否使用 data 命令生成数据库")
		}

	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
