package models

// Class 词分类
type Class struct {
	ID    int
	Class string
}

// Classes 敏感词列表
type Classes []Class

// TableName 指定表名
func (Class) TableName() string {
	return "classes"
}
