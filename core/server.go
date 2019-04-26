package core

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iohub/Ahocorasick"

)

// MapClasses 全局分类map，常驻内存
var MapClasses map[int]string

// ACMatcher AC算法匹配结构，常驻内存
var ACMatcher *cedar.Matcher

// ServerStart 开始服务
func ServerStart() {
	// 设置全局分类map
	MapClasses = getClassesToMap()

	// AC算法匹配结构
	ACMatcher = buildAC()

	// 设置模式
	gin.SetMode(Conf.Get("server.Mode"))

	// 设置访问日志
	f, _ := os.Create(Conf.Get("server.AccessLog"))
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	rv100 := router.Group("/v1.0.0")
	{
		// 获取原始词库，JSON格式
		rv100.GET("/words", wordsHandler)
		// 获取所有需要检查的词库，JSON格式
		rv100.GET("/check_words", checkWordsHandler)
		// 检查所传输的内容
		rv100.POST("/check", checkHandler)
	}

	port := fmt.Sprintf(":%d", Conf.GetInt("server.Port"))
	router.Run(port)
}

// buildAC 构建AC算法匹配结构
func buildAC() *cedar.Matcher{
	// 构建AC算法，加入匹配词
	m := cedar.NewMatcher()
	for _, w := range getCheckWords() {
		m.Insert([]byte(w.Name), w.ClassID)
	}
	m.Compile()

	return m
}

// wordsReq 词库处理参数
type wordsReq struct {
	Wtype int `form:"type" json:"type"`
}

// wordsHandler 获取词库，JSON格式
func wordsHandler(c *gin.Context) {
	var query wordsReq
	if err := c.ShouldBindQuery(&query); err != nil {
		respError(http.StatusBadRequest, err, c)
		return
	}
	// bodyByte, _ := json.Marshal(query)

	wds := getWords(query.Wtype)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": wds,
	})
}

// checkWordsHandler 获取所有需要检查的词库，JSON格式
func checkWordsHandler(c *gin.Context) {
	var rs []map[string]string
	for _, w := range getCheckWords() {
		rs = append(rs, map[string]string{
			"class": MapClasses[w.ClassID],
			"name": w.Name,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": rs,
	})
}

// checkReq 检查所传输的内容处理参数
type checkReq struct {
	Content string `form:"content" json:"content" binding:"required"`
}

// checkHandler 检查所传输的内容
func checkHandler(c *gin.Context) {
	var query checkReq
	if err := c.ShouldBindJSON(&query); err != nil {
		respError(http.StatusBadRequest, err, c)
		return
	}
	
	// 匹配
	seq := []byte(query.Content)
	req := ACMatcher.Match(seq)
	var rs []map[string]interface{}
	for _, item := range req {
		key := ACMatcher.TokenOf(seq, item)
		rs = append(rs, map[string]interface{}{
			"word": string(key),
			"class": MapClasses[item.Value.(int)],
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": rs,
	})
}

// respError 错误响应
func respError(code int, err error, c *gin.Context) {
	var message string

	allowCode := (code == http.StatusBadRequest || code == http.StatusUnauthorized)
	if err != nil && (Conf.GetBool("server.Debug") || allowCode) {
		message = err.Error()
	} else {
		message = errMess[code]
	}

	c.AbortWithStatusJSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
