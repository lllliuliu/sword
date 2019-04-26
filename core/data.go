package core

import (
	"strings"

	"sword/core/models"
)

// Data 数据处理结构
type Data struct {
	db      *DBClient
	excel   *excel
	words   models.Words
	classes map[int]models.Class
}

// NewData 构造data
func NewData(file string) (*Data, error) {
	return &Data{
		db:      DB,
		excel:   &excel{file},
		words:   nil,
		classes: make(map[int]models.Class),
	}, nil
}

// ExcelToDB 从excel文件获取数据到数据库
func (d *Data) ExcelToDB() error {
	var err error

	d.words, d.classes, err = d.excel.resolve()
	if err != nil {
		return err
	}

	if err = d.init(); err != nil {
		return err
	}

	return d.buildForExcel()
}

// init 初始化
func (d *Data) init() error {
	sql := `
	DROP TABLE IF EXISTS words;
	CREATE TABLE IF NOT EXISTS words(
		id INTEGER NOT NULL PRIMARY KEY,
		name VARCHAR NOT NULL,
		class_id INTEGER,
		type_id INTEGER,
		time DATETIME
	);
	DROP TABLE IF EXISTS classes;
	CREATE TABLE IF NOT EXISTS classes(
		id INTEGER NOT NULL PRIMARY KEY,
		class VARCHAR NOT NULL
	);
	DROP TABLE IF EXISTS check_words;
	CREATE TABLE IF NOT EXISTS check_words(
		id INTEGER NOT NULL PRIMARY KEY,
		class_id INTEGER,
		name VARCHAR NOT NULL
	);
	`
	err := d.db.Exec(sql).Error
	if err != nil {
		return err
	}
	return nil
}

// buildForExcel 构建数据
func (d *Data) buildForExcel() error {

	// 构建 classes 表
	sqlStr := "INSERT INTO classes(id, class) VALUES "
	vals := []interface{}{}

	for id, class := range d.classes {
		sqlStr += "(?, ?),"
		vals = append(vals, id, class.Class)
	}
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	err := d.db.Exec(sqlStr, vals...).Error
	if err != nil {
		return err
	}

	// 构建 words 表，数据较多每100条插入一次
	max := 100
	var sqlSlice []string
	vals = []interface{}{}
	sqlPre := "INSERT INTO words(name, class_id, type_id, time) VALUES "
	for _, word := range d.words {
		if len(sqlSlice) == max {
			err := d.db.Exec(sqlPre+strings.Join(sqlSlice, ","), vals...).Error
			if err != nil {
				return err
			}

			sqlSlice = []string{}
			vals = []interface{}{}
		}

		sqlSlice = append(sqlSlice, "(?, ?, ?, ?)")
		vals = append(vals, word.Name, word.ClassID, word.TypeID, word.Time)
	}

	// 构建check_words表，先用buildCheckWords构建所有需要检查的词，然后存入
	// 数据库，数据较多每100条插入一次
	sqlSlice = []string{}
	vals = []interface{}{}
	sqlPre = "INSERT INTO check_words(class_id, name) VALUES "
	for cid, ws := range buildCheckWords() {
		for _, w := range ws {
			if len(sqlSlice) == max {
				err := d.db.Exec(sqlPre+strings.Join(sqlSlice, ","), vals...).Error
				if err != nil {
					return err
				}

				sqlSlice = []string{}
				vals = []interface{}{}
			}

			sqlSlice = append(sqlSlice, "(?, ?)")
			vals = append(vals, cid, w)
		}
	}

	return nil
}
