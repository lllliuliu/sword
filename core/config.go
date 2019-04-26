package core

import (
	"fmt"

	"github.com/spf13/viper"
	// "github.com/fsnotify/fsnotify"
)

const (
	filePath = ""
	fileName = "./conf"
)

// Conf 全局变量配置
var Conf *Config

// Config 配置结构体
type Config struct {
	*viper.Viper
}

// InitConf 配置初始化
func initConf(fileName string) {
	Conf = newConfig(fileName)
}

// newConfig 创建一个配置结构
// 读取指定配置文件，暂不读取配置文件夹
func newConfig(fileName string) *Config {
	viper := viper.New()
	viper.SetConfigName(fileName)

	// if filePath != "" {
	// 	viper.AddConfigPath(filePath)
	// }
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Not found config file:%s", fileName)
	}

	// 运行时文件重载，暂时不开启
	// viper.WatchConfig()
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	fmt.Printf("Config file changed:%s", e.Name)
	// })

	return &Config{viper}
}

// Get 封装最常用的字符串配置
func (c *Config) Get(key string) string {
	return c.GetString(key)
}
