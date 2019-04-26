package core

// InitCore 核心初始化
func InitCore(cfile string) {
	initConf(cfile)
	initDB()
}
