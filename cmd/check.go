package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"sword/core"
)

var content, path string

// checkCmd 命令检查内容
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "检查内容是否包含敏感词",
	Long:  `使用 JSON 格式，获取敏感词表.`,
	Run: func(cmd *cobra.Command, args []string) {
		if content == "" && path == "" {
			fmt.Println("请使用-c参数指定检查内容或使用-p参数指定检查文档路径", N)
			return
		}

		var err error
		var data []map[string]interface{}

		if content != "" {
			data, err = core.CheckForContent(content)
			if err != nil {
				fmt.Println("指定内容检查失败", N)
				fmt.Println("错误信息为：", err)
			}
		}

		if path != "" {
			data, err = core.CheckForPath(path)
			if err != nil {
				fmt.Println("指定文档路径检查失败", N)
				fmt.Println("错误信息为：", err)
				return
			}
		}

		s, _ := json.MarshalIndent(data, "", " ")
		fmt.Println(string(s))
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().StringVarP(&content, "content", "t", "", "直接指定需要检查的内容")
	checkCmd.Flags().StringVarP(&path, "path", "p", "", "指定需要检查的文档路径")
}
