package cmd

import (
	"fmt"
	"sword/core"

	"github.com/spf13/cobra"
)

var serveAction string

// serveCmd命令提供web服务
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "启动网络服务，提供敏感词获取接口和文字内容检查接口",
	Long:  `默认启动80端口，提供一个基于HTTP服务的接口服务，具体文档请查阅xxx.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.ServerStart()
		fmt.Println("启动服务")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(&serveAction, "action", "a", "", "需要执行的操作,包含start、stop、reload")
}
