package models

import "time"

// 词类型，固定常量
const (
	_ = iota
	// Verb 动词
	Verb int = iota*2 - 1
	// Noun 名词
	Noun
	// Exclusive 专属词语
	Exclusive
)

// WTypes 敏感词类型说明
var WTypes = map[int]string{
	Verb:      "动词",
	Noun:      "名词",
	Exclusive: "专属词语",
}

// Word 原始敏感词
type Word struct {
	ID      int
	ClassID int `json:"-"`
	TypeID  int
	Name    string
	Time    time.Time
	Class   string `gorm:"-"`
}

// Words 敏感词列表
type Words []Word

// TableName 指定表名
func (Word) TableName() string {
	return "words"
}
