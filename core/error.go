package core

import "errors"

var (
	ERR_DB_CONNECT_FAILED = errors.New("数据库连接失败")
	ERR_DB_INIT_FAILED    = errors.New("数据库初始化失败")
	ERR_DB_PREPARE_FAILED = errors.New("数据库预处理失败")
	ERR_DB_EXE_FAILED     = errors.New("数据库语句执行失败")

	ERR_EXCEL_FILE_FAILED  = errors.New("Excel文件打开错误")
	ERR_EXCEL_SHEET_FAILED = errors.New("Excel打开工作表错误")
)

var (
	// errMess 错误代码和信息
	errMess = map[int]string{
		400: "参数错误",
		500: "服务器错误",
	}
)
