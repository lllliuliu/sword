package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"sword/core"
)

var dataExcel string

// dataCmd 命令直接构建数据源
var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "基于Excel文件生成数据库，默认使用sqlite3",
	Long:  "默认在当前目录查找Excel文件，如果文件在其他目录请使用-e参数指定文件路径；文件需要特定格式。",
	Run: func(cmd *cobra.Command, args []string) {
		if dataExcel == "" {
			fmt.Println("请使用-e参数指定文件路径", N)
			return
		}
		fmt.Println("数据库生成中...")

		if d, err := core.NewData(dataExcel); err != nil {
			fmt.Println("数据库生成失败", N)
			fmt.Println("错误信息为：", err)
		} else {
			if err := d.ExcelToDB(); err == nil {
				fmt.Println("数据库生成成功", Y)
			} else {
				fmt.Println("数据库生成失败", N)
				fmt.Println("错误信息为：", err)
			}
		}

		return
	},
}

func init() {
	rootCmd.AddCommand(dataCmd)

	dataCmd.Flags().StringVarP(&dataExcel, "excel", "e", "", "指定excel文件路径")
}
