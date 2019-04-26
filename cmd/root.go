// Initialize cobra and reload configure

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"sword/core"
)

const (
	Y = "\u2714"
	N = "\u2718"
)

var cfile string

const defaultConf = ".sword"

// rootCmd 表示基础命令，也就是根命令
var rootCmd = &cobra.Command{
	Use:   "sword",
	Short: "敏感词的查阅和内容敏感词检查",
	Long: `直接获取敏感词列表.	
或启动一个基于HTTP协议的接口服务，提供敏感词列表获取和文字内容检查的功能.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&cfile, "config", "c", "", "配置文件(默认使用当前目录的 "+defaultConf+" 文件)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "全局帮助信息")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfile != "" {
		// Use config file from the flag.
		core.InitCore(cfile)
	} else {
		// 默认使用当前目录的 .sword.toml 文件
		core.InitCore(defaultConf)
	}
}
