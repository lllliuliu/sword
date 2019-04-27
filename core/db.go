package core

import (
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
	// gorm sqlite 驱动
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB 全局数据库处理
var DB *DBClient

// DBClient 结构体，包含额外方法
type DBClient struct {
	*gorm.DB
}

// initDB 数据库初始化
func initDB() error {
	var err error
	dbfile := Conf.Get("db.DB")
	paths, _ := filepath.Split(dbfile)
	err = os.MkdirAll(paths, 0755)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(Conf.Get("db.Driver"), dbfile)
	if err != nil {
		panic(err)
	}

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	DB = &DBClient{db}
	return nil
}
