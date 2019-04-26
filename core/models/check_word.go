package models

// CheckWord 检查词
type CheckWord struct {
	ID      int
	ClassID int
	Name    string
}

// TableName 指定表名
func (CheckWord) TableName() string {
	return "check_words"
}
