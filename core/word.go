package core

import (
	"fmt"
	"io/ioutil"
	"sword/core/models"
)

// CheckForContent 从文件检查内容
func CheckForContent(c string) ([]map[string]interface{}, error) {
	b := []byte(c)
	acmat := buildAC()
	mapclass := getClassesToMap()
	req := acmat.Match(b)
	var rs []map[string]interface{}
	for _, item := range req {
		key := acmat.TokenOf(b, item)
		rs = append(rs, map[string]interface{}{
			"word":  string(key),
			"class": mapclass[item.Value.(int)],
		})
	}

	return rs, nil
}

// CheckForPath 从文件检查内容
func CheckForPath(path string) ([]map[string]interface{}, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	acmat := buildAC()
	mapclass := getClassesToMap()
	req := acmat.Match(b)
	var rs []map[string]interface{}
	for _, item := range req {
		key := acmat.TokenOf(b, item)
		rs = append(rs, map[string]interface{}{
			"word":  string(key),
			"class": mapclass[item.Value.(int)],
		})
	}

	return rs, nil
}

// FetchWords 从数据源获取词库
func FetchWords() ([]map[string]string, error) {
	var rs []map[string]string
	mapclass := getClassesToMap()
	for _, w := range getCheckWords() {
		rs = append(rs, map[string]string{
			"class": mapclass[w.ClassID],
			"word":  w.Name,
		})
	}
	return rs, nil
}

// getClassesToMap 从分类表获取分类map，id为key
func getClassesToMap() map[int]string {
	var cs models.Classes
	DB.Table("classes").Select("id,class").Find(&cs)

	rs := make(map[int]string)
	for _, class := range cs {
		rs[class.ID] = class.Class
	}

	return rs
}

// getWords 根据传入的类型获取敏感词
func getWords(t int) models.Words {
	var wds models.Words
	d := DB.Table("words").Select("words.id,name,class_id,type_id,time,class").Joins("INNER JOIN classes ON words.class_id = classes.id")
	if t > 0 {
		d = d.Where("words.type_id=?", t)
	}
	d.Find(&wds)

	return wds
}

// getCheckWords 获取所有需要检查的词
// 从 check_words 中直接获取
func getCheckWords() []models.CheckWord {
	var cws []models.CheckWord
	DB.Table("check_words").Select("class_id,name").Find(&cws)

	return cws
}

// buildCheckWords 构建所有需要检查的词
// - 专属名词都需要检查
// - 同类下动词和名词的笛卡尔积都需要检查
func buildCheckWords() map[int][]string {
	// 总词库
	cw := make(map[int][]string)

	// 分类获取敏感词
	verbs := getWords(models.Verb)
	nouns := getWords(models.Noun)
	exclusives := getWords(models.Exclusive)

	// 动名词笛卡尔积
	sort := make(map[int]map[int][]string)
	for _, v := range verbs {
		if _, exist := sort[v.ClassID]; exist {
			sort[v.ClassID][models.Verb] = append(sort[v.ClassID][models.Verb], v.Name)
		} else {
			temp := make(map[int][]string)
			temp[models.Verb] = append(temp[models.Verb], v.Name)
			sort[v.ClassID] = temp
		}
	}
	for _, n := range nouns {
		if _, exist := sort[n.ClassID]; exist {
			sort[n.ClassID][models.Noun] = append(sort[n.ClassID][models.Noun], n.Name)
		} else {
			temp := make(map[int][]string)
			temp[models.Noun] = append(temp[models.Noun], n.Name)
			sort[n.ClassID] = temp
		}
	}
	for cid, t := range sort {
		for _, verb := range t[models.Verb] {
			for _, noun := range t[models.Noun] {
				cw[cid] = append(cw[cid], fmt.Sprintf("%s%s", verb, noun))
			}
		}
	}

	// 添加专属词语
	for _, e := range exclusives {
		cw[e.ClassID] = append(cw[e.ClassID], e.Name)
	}

	return cw

}
