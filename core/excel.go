package core

import (
	"strings"
	"time"

	"sword/core/models"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// excel excel处理结构，包含文件地址和格式化词库
type excel struct {
	file string
}

// resolve 从文件解析相关数据
func (e *excel) resolve() (models.Words, map[int]models.Class, error) {
	xlsx, err := excelize.OpenFile(e.file)
	if err != nil {
		return nil, nil, ERR_EXCEL_FILE_FAILED
	}

	sheet := xlsx.GetSheetName(1)
	if sheet == "" {
		return nil, nil, ERR_EXCEL_SHEET_FAILED
	}

	var wds models.Words
	wcs := make(map[int]models.Class)

	var wclass, wtype int
	var name string

	for n, row := range xlsx.GetRows(sheet) {
		for m, cell := range row {
			// 第0行每6列的值，作为分类
			if n == 0 && m%6 == 0 {
				wcs[m/6+1] = models.Class{Class: strings.TrimSpace(cell)}
			}

			// 从第3行开始记录敏感词
			if n > 2 {
				if m%2 == 0 {
					wclass = m/6 + 1
					switch m%6 + 1 {
					case models.Verb:
						wtype = models.Verb
					case models.Noun:
						wtype = models.Noun
					case models.Exclusive:
						wtype = models.Exclusive
					}
					name = strings.TrimSpace(cell)
				} else {
					t, _ := time.Parse("2006-01-02 15:04:05", strings.TrimSpace(cell))
					if name != "" {
						wds = append(wds, models.Word{
							ClassID: wclass,
							TypeID:  wtype,
							Name:    name,
							Time:    t,
						})
					}
				}
			}
		}
	}

	return wds, wcs, nil
}
